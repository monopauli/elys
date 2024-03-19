package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	atypes "github.com/elys-network/elys/x/assetprofile/types"
	"github.com/elys-network/elys/x/launchpad/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k Keeper) isEnabledToken(ctx sdk.Context, spendingToken string) bool {
	params := k.GetParams(ctx)

	for _, token := range params.SpendingTokens {
		if token == spendingToken {
			return true
		}
	}
	return false
}

func (k Keeper) CreateOrder(ctx sdk.Context, orderMaker string, spendingToken string, elysAmount math.Int, bonusRate sdk.Dec) {
	params := k.GetParams(ctx)
	order := types.Purchase{}
	order.OrderId = k.LastOrderId(ctx) + 1
	order.OrderMaker = orderMaker
	order.SpendingToken = spendingToken
	order.ElysAmount = elysAmount
	order.BonusAmount = bonusRate.MulInt(elysAmount).TruncateInt()
	order.TokenAmount = sdk.NewDecFromInt(elysAmount).Mul(sdk.NewDec(1000_000)).Quo(params.InitialPrice).RoundInt()
	order.ReturnedElysAmount = sdk.ZeroInt()

	k.SetOrder(ctx, order)
}

func (k msgServer) BuyElys(goCtx context.Context, msg *types.MsgBuyElys) (*types.MsgBuyElysResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if params.LaunchpadStarttime > uint64(ctx.BlockTime().Unix()) {
		return nil, types.ErrLaunchpadNotStarted
	}

	if params.LaunchpadStarttime+params.LaunchpadDuration < uint64(ctx.BlockTime().Unix()) {
		return nil, types.ErrLaunchpadAlreadyFinished
	}

	if !k.isEnabledToken(ctx, msg.SpendingToken) {
		return nil, types.ErrNotEnabledSpendingToken
	}

	asset, found := k.assetProfileKeeper.GetEntry(ctx, msg.SpendingToken)
	if !found {
		return nil, errorsmod.Wrapf(atypes.ErrAssetProfileNotFound, "denom: %s", msg.SpendingToken)
	}

	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	spendingCoins := sdk.Coins{sdk.NewCoin(asset.Denom, msg.TokenAmount)}
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, spendingCoins)
	if err != nil {
		return nil, err
	}

	price := k.oracleKeeper.GetAssetPriceFromDenom(ctx, asset.Denom)

	elysAmount := price.MulInt(msg.TokenAmount).Mul(sdk.NewDec(1000_000)).Quo(params.InitialPrice).RoundInt()

	soldAmount := params.SoldAmount.Add(elysAmount)
	if soldAmount.GT(params.TotalReserve) {
		return nil, types.ErrOverflowTotalReserve
	}

	// 0-20% raise 100% bonus
	// 20-30% raise 90% bonus
	// 30-40% raise bonus 80%
	// 40-50% raise bonus 70%
	// 50-60% raise bonus 60%
	// 60-70% raise bonus 50%
	// 70-80% raise bonus 40%
	// 80-100% raise bonus 30%
	soldSoFar := params.SoldAmount
	for index, raisePercent := range params.BonusInfo.RaisePercents {
		roundMaxRaise := sdk.NewDecWithPrec(int64(raisePercent), 2).Mul(sdk.NewDecFromInt(params.TotalReserve)).RoundInt()
		if soldSoFar.LT(roundMaxRaise) {
			bonusPercent := params.BonusInfo.BonusPercents[index]
			bonusRate := sdk.NewDecWithPrec(int64(bonusPercent), 2)
			if roundMaxRaise.GTE(soldAmount) {
				roundSellAmount := soldAmount.Sub(soldSoFar)
				k.CreateOrder(ctx, msg.Sender, msg.SpendingToken, roundSellAmount, bonusRate)
				break
			} else {
				roundSellAmount := roundMaxRaise.Sub(soldSoFar)
				k.CreateOrder(ctx, msg.Sender, msg.SpendingToken, roundSellAmount, bonusRate)
				soldSoFar = roundMaxRaise
			}
		}
	}

	params.SoldAmount = soldAmount
	k.SetParams(ctx, params)

	elysCoins := sdk.Coins{sdk.NewCoin(ptypes.Elys, elysAmount)}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, elysCoins)
	if err != nil {
		return nil, err
	}

	return &types.MsgBuyElysResponse{}, nil
}

func (k msgServer) ReturnElys(goCtx context.Context, msg *types.MsgReturnElys) (*types.MsgReturnElysResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	launchpadEndTime := params.LaunchpadStarttime + params.LaunchpadDuration
	if launchpadEndTime > uint64(ctx.BlockTime().Unix()) {
		return nil, types.ErrLaunchpadNotFinished
	}

	if params.LaunchpadStarttime+params.ReturnDuration < uint64(ctx.BlockTime().Unix()) {
		return nil, types.ErrLaunchpadReturnPeriodFinished
	}

	order := k.GetOrder(ctx, msg.OrderId)
	if order.OrderId == 0 {
		return nil, types.ErrPurchaseOrderNotFound
	}

	if order.OrderMaker != msg.Sender {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "expected %s, got %s", order.OrderMaker, msg.Sender)
	}

	maxReturnAmount := sdk.NewDecWithPrec(int64(params.MaxReturnPercent), 2).MulInt(order.ElysAmount).RoundInt()
	if order.ReturnedElysAmount.Add(msg.ReturnElysAmount).GT(maxReturnAmount) {
		return nil, types.ErrExceedMaxReturnAmount
	}

	returnTokenAmount := msg.ReturnElysAmount.Mul(order.TokenAmount).Quo(order.ElysAmount)

	asset, found := k.assetProfileKeeper.GetEntry(ctx, order.SpendingToken)
	if !found {
		return nil, errorsmod.Wrapf(atypes.ErrAssetProfileNotFound, "denom: %s", order.SpendingToken)
	}

	coins := sdk.Coins{sdk.NewCoin(asset.Denom, returnTokenAmount)}
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgReturnElysResponse{}, nil
}

func (k msgServer) WithdrawRaised(goCtx context.Context, msg *types.MsgWithdrawRaised) (*types.MsgWithdrawRaisedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	if params.WithdrawAddress != msg.Sender {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "expected %s, got %s", params.WithdrawAddress, msg.Sender)
	}

	// Once return period is over, can withdraw all
	returnEndTime := params.LaunchpadStarttime + params.LaunchpadDuration + params.ReturnDuration
	if returnEndTime < uint64(ctx.BlockTime().Unix()) {
		addr := sdk.MustAccAddressFromBech32(msg.Sender)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.Coins(msg.Coins))
		if err != nil {
			return nil, err
		}
	}

	totalWithdrawAmount := sdk.Coins(params.WithdrawnAmount).Add(msg.Coins...)
	maxWithdrawAmount := sdk.Coins{}
	orders := k.GetAllOrders(ctx)
	for _, order := range orders {
		orderWithdrawable := sdk.NewDecWithPrec(int64(100-params.MaxReturnPercent), 2).MulInt(order.TokenAmount).RoundInt()
		asset, found := k.assetProfileKeeper.GetEntry(ctx, order.SpendingToken)
		if !found {
			return nil, errorsmod.Wrapf(atypes.ErrAssetProfileNotFound, "denom: %s", order.SpendingToken)
		}
		coin := sdk.NewCoin(asset.Denom, orderWithdrawable)
		maxWithdrawAmount = maxWithdrawAmount.Add(coin)
	}

	if totalWithdrawAmount.IsAnyGT(maxWithdrawAmount) {
		return nil, types.ErrExceedMaxWithdrawableAmount
	}

	coins := sdk.Coins(msg.Coins)
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawRaisedResponse{}, nil
}

// Update params through gov proposal
func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	// store params
	k.SetParams(ctx, msg.Params)

	return &types.MsgUpdateParamsResponse{}, nil
}
