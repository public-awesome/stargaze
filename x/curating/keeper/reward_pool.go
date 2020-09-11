package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// GetRewardPool returns the reward pool account.
func (k Keeper) GetRewardPool(ctx sdk.Context) (rewardPool authexported.ModuleAccountI) {
	return k.accountKeeper.GetModuleAccount(ctx, types.RewardPoolName)
}

// GetRewardPoolBalance returns the reward pool balance
func (k Keeper) GetRewardPoolBalance(ctx sdk.Context) sdk.Coin {
	return k.bankKeeper.GetBalance(ctx, k.GetRewardPool(ctx).GetAddress(), types.DefaultStakeDenom)
}

// InflateRewardPool process the designated inflation to the reward pool
func (k Keeper) InflateRewardPool(ctx sdk.Context) error {
	blockInflationAddr := k.accountKeeper.GetModuleAccount(ctx, auth.FeeCollectorName).GetAddress()
	blockInflation := k.bankKeeper.GetBalance(ctx, blockInflationAddr, types.DefaultStakeDenom)
	rewardPoolAllocation := k.GetParams(ctx).RewardPoolAllocation

	blockInflationDec := sdk.NewDecFromInt(blockInflation.Amount)
	rewardAmount := blockInflationDec.Mul(rewardPoolAllocation)
	rewardCoin := sdk.NewCoin(types.DefaultStakeDenom, rewardAmount.TruncateInt())

	return k.bankKeeper.SendCoinsFromModuleToModule(
		ctx, auth.FeeCollectorName, types.RewardPoolName, sdk.NewCoins(rewardCoin))
}

// InitializeRewardPool sets up the reward pool from genesis
func (k Keeper) InitializeRewardPool(ctx sdk.Context, funds sdk.Coin) error {
	return k.bankKeeper.MintCoins(ctx, types.RewardPoolName, sdk.NewCoins(funds))
}
