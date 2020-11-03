package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// GetRewardPool returns the reward pool account.
func (k Keeper) GetRewardPool(ctx sdk.Context) (rewardPool authtypes.ModuleAccountI) {
	return k.accountKeeper.GetModuleAccount(ctx, types.RewardPoolName)
}

// GetRewardPoolBalance returns the reward pool balance
func (k Keeper) GetRewardPoolBalance(ctx sdk.Context) sdk.Coin {
	return k.bankKeeper.GetBalance(ctx, k.GetRewardPool(ctx).GetAddress(), k.GetParams(ctx).StakeDenom)
}

// InflateRewardPool process the designated inflation to the reward pool
func (k Keeper) InflateRewardPool(ctx sdk.Context) error {
	blockInflationAddr := k.accountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName).GetAddress()
	blockInflation := k.bankKeeper.GetBalance(ctx, blockInflationAddr, k.GetParams(ctx).StakeDenom)
	rewardPoolAllocation := k.GetParams(ctx).RewardPoolAllocation

	blockInflationDec := sdk.NewDecFromInt(blockInflation.Amount)
	rewardAmount := blockInflationDec.Mul(rewardPoolAllocation)
	rewardCoin := sdk.NewCoin(k.GetParams(ctx).StakeDenom, rewardAmount.TruncateInt())

	return k.bankKeeper.SendCoinsFromModuleToModule(
		ctx, authtypes.FeeCollectorName, types.RewardPoolName, sdk.NewCoins(rewardCoin))
}

// InitializeRewardPool sets up the reward pool from genesis
func (k Keeper) InitializeRewardPool(ctx sdk.Context, funds sdk.Coin) error {
	return k.bankKeeper.MintCoins(ctx, types.RewardPoolName, sdk.NewCoins(funds))
}
