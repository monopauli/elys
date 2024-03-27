package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	simapp "github.com/elys-network/elys/app"
	keepertest "github.com/elys-network/elys/testutil/keeper"
	"github.com/elys-network/elys/x/launchpad/keeper"
	"github.com/elys-network/elys/x/launchpad/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LaunchpadKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestMsgServerBuyElys(t *testing.T) {
	app := simapp.InitElysTestApp(true)
	ctx := app.BaseApp.NewContext(true, tmproto.Header{})
	k := app.LaunchpadKeeper
	SetupStableCoinPrices(app, ctx)

	coins := sdk.Coins{sdk.NewInt64Coin(ptypes.Elys, 1000_000_000_000_000)}
	err := app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	require.NoError(t, err)
	err = app.BankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.ModuleName, coins)
	require.NoError(t, err)

	buyer := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	coins = sdk.Coins{sdk.NewInt64Coin(ptypes.BaseCurrency, 1100_000_000_000)}
	err = app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	require.NoError(t, err)
	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, buyer, coins)
	require.NoError(t, err)

	now := time.Now()
	ctx = ctx.WithBlockTime(now)
	params := k.GetParams(ctx)
	params.LaunchpadStarttime = uint64(now.Unix())
	k.SetParams(ctx, params)

	msgServer := keeper.NewMsgServerImpl(k)
	cacheCtx, _ := ctx.CacheContext()
	response, err := msgServer.BuyElys(cacheCtx, &types.MsgBuyElys{
		Sender:        buyer.String(),
		SpendingToken: ptypes.BaseCurrency,
		TokenAmount:   sdk.NewInt(1000_000),
	})
	require.NoError(t, err)
	require.Len(t, response.OrderIds, 1)
	require.Equal(t, response.OrderIds[0], uint64(1))
	orders := k.GetAllOrders(cacheCtx)
	require.Len(t, orders, 1)
	require.Equal(t, orders[0].OrderId, uint64(1))
	require.Equal(t, orders[0].OrderMaker, buyer.String())
	require.Equal(t, orders[0].SpendingToken, ptypes.BaseCurrency)
	require.Equal(t, orders[0].TokenAmount.String(), "1000000")
	require.Equal(t, orders[0].ElysAmount.String(), "1333333")
	require.Equal(t, orders[0].ReturnedElysAmount.String(), "0")
	require.Equal(t, orders[0].BonusAmount.String(), "1333333")
	require.Equal(t, orders[0].VestingStarted, false)

	cacheCtx, _ = ctx.CacheContext()
	response, err = msgServer.BuyElys(cacheCtx, &types.MsgBuyElys{
		Sender:        buyer.String(),
		SpendingToken: ptypes.BaseCurrency,
		TokenAmount:   sdk.NewInt(1000_000_000_000),
	})
	require.NoError(t, err)
	require.Len(t, response.OrderIds, 2)
	require.Equal(t, response.OrderIds[0], uint64(1))
	require.Equal(t, response.OrderIds[1], uint64(2))
	orders = k.GetAllOrders(cacheCtx)
	require.Len(t, orders, 2)
	require.Equal(t, orders[0].OrderId, uint64(1))
	require.Equal(t, orders[0].OrderMaker, buyer.String())
	require.Equal(t, orders[0].SpendingToken, ptypes.BaseCurrency)
	require.Equal(t, orders[0].TokenAmount.String(), "675000000000")
	require.Equal(t, orders[0].ElysAmount.String(), "900000000000")
	require.Equal(t, orders[0].ReturnedElysAmount.String(), "0")
	require.Equal(t, orders[0].BonusAmount.String(), "900000000000")
	require.Equal(t, orders[0].VestingStarted, false)
	require.Equal(t, orders[1].OrderId, uint64(2))
	require.Equal(t, orders[1].OrderMaker, buyer.String())
	require.Equal(t, orders[1].SpendingToken, ptypes.BaseCurrency)
	require.Equal(t, orders[1].TokenAmount.String(), "325000000000")
	require.Equal(t, orders[1].ElysAmount.String(), "433333333333")
	require.Equal(t, orders[1].ReturnedElysAmount.String(), "0")
	require.Equal(t, orders[1].BonusAmount.String(), "389999999999")
	require.Equal(t, orders[1].VestingStarted, false)
}

// TODO:
// func (k Keeper) IsEnabledToken(ctx sdk.Context, spendingToken string) bool {
// func (k Keeper) GenerateOrder(ctx sdk.Context, orderMaker string, spendingToken string, elysAmount math.Int, bonusRate sdk.Dec, price sdk.Dec) types.Purchase {
// func (k Keeper) CalcBuyElysResult(ctx sdk.Context, sender string, spendingToken string, tokenAmount math.Int) (math.Int, []types.Purchase, error) {
// func (k Keeper) CalcReturnElysResult(ctx sdk.Context, orderId uint64, returnElysAmount math.Int) (math.Int, error) {
// func (k msgServer) ReturnElys(goCtx context.Context, msg *types.MsgReturnElys) (*types.MsgReturnElysResponse, error) {
// func (k msgServer) DepositElysToken(goCtx context.Context, msg *types.MsgDepositElysToken) (*types.MsgDepositElysTokenResponse, error) {
// func (k msgServer) WithdrawRaised(goCtx context.Context, msg *types.MsgWithdrawRaised) (*types.MsgWithdrawRaisedResponse, error) {
