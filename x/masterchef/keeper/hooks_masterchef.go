package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ammtypes "github.com/elys-network/elys/x/amm/types"
	"github.com/elys-network/elys/x/masterchef/types"
)

func (k Keeper) GetPoolTotalSupply(ctx sdk.Context, poolId uint64) sdk.Int {
	pool, found := k.amm.GetPool(ctx, poolId)
	if !found {
		return sdk.ZeroInt()
	}

	return pool.TotalShares.Amount
}

func (k Keeper) GetPoolBalance(ctx sdk.Context, poolId uint64, user string) sdk.Int {
	commitments := k.cmk.GetCommitments(ctx, user)

	return commitments.GetCommittedAmountForDenom(ammtypes.GetPoolShareDenom(poolId))
}

func (k Keeper) UpdateAccPerShare(ctx sdk.Context, poolId uint64, rewardDenom string, amount sdk.Int) {
	poolRewardInfo, found := k.GetPoolRewardInfo(ctx, poolId, rewardDenom)

	if !found {
		poolRewardInfo = types.PoolRewardInfo{
			PoolId:                poolId,
			RewardDenom:           rewardDenom,
			PoolAccRewardPerShare: sdk.NewDec(0),
			LastUpdatedBlock:      0,
		}
	}

	poolRewardInfo.PoolAccRewardPerShare = poolRewardInfo.PoolAccRewardPerShare.Add(
		math.LegacyNewDecFromInt(amount).Quo(math.LegacyNewDecFromInt(k.GetPoolTotalSupply(ctx, poolId))),
	)
	poolRewardInfo.LastUpdatedBlock = uint64(ctx.BlockHeight())

	k.SetPoolRewardInfo(ctx, poolRewardInfo)
}

func (k Keeper) UpdateUserRewardPending(ctx sdk.Context, poolId uint64, rewardDenom string, user string, isDeposit bool, amount sdk.Int) {
	poolRewardInfo, found := k.GetPoolRewardInfo(ctx, poolId, rewardDenom)

	if !found {
		poolRewardInfo = types.PoolRewardInfo{
			PoolId:                poolId,
			RewardDenom:           rewardDenom,
			PoolAccRewardPerShare: sdk.NewDec(0),
			LastUpdatedBlock:      0,
		}
	}

	userRewardInfo, found := k.GetUserRewardInfo(ctx, user, poolId, rewardDenom)

	if !found {
		userRewardInfo = types.UserRewardInfo{
			User:          user,
			PoolId:        poolId,
			RewardDenom:   rewardDenom,
			RewardDebt:    sdk.NewDec(0),
			RewardPending: sdk.NewDec(0),
		}
	}

	// need to consider balance change on deposit/withdraw
	userBalance := k.GetPoolBalance(ctx, poolId, user)
	if isDeposit {
		userBalance = userBalance.Sub(amount)
	} else {
		userBalance = userBalance.Add(amount)
	}

	userRewardInfo.RewardPending = userRewardInfo.RewardPending.Add(
		poolRewardInfo.PoolAccRewardPerShare.Mul(
			math.LegacyDec(userBalance),
		).Sub(userRewardInfo.RewardDebt).Quo(sdk.NewDecFromIntWithPrec(math.OneInt(), 18)),
	)

	k.SetUserRewardInfo(ctx, userRewardInfo)
}

func (k Keeper) UpdateUserRewardDebt(ctx sdk.Context, poolId uint64, rewardDenom string, user string) {
	poolRewardInfo, found := k.GetPoolRewardInfo(ctx, poolId, rewardDenom)

	if !found {
		poolRewardInfo = types.PoolRewardInfo{
			PoolId:                poolId,
			RewardDenom:           rewardDenom,
			PoolAccRewardPerShare: sdk.NewDec(0),
			LastUpdatedBlock:      0,
		}
	}

	userRewardInfo, found := k.GetUserRewardInfo(ctx, user, poolId, rewardDenom)

	if !found {
		userRewardInfo = types.UserRewardInfo{
			User:          user,
			PoolId:        poolId,
			RewardDenom:   rewardDenom,
			RewardDebt:    sdk.NewDec(0),
			RewardPending: sdk.NewDec(0),
		}
	}

	userRewardInfo.RewardDebt = poolRewardInfo.PoolAccRewardPerShare.Mul(
		math.LegacyDec(k.GetPoolBalance(ctx, poolId, user)),
	)

	k.SetUserRewardInfo(ctx, userRewardInfo)
}
