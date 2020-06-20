package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/public-awesome/stakebird/x/stake/types"
)

func (k Keeper) GetRewardPool(ctx sdk.Context) (rewardPool authexported.ModuleAccountI) {
	return k.accountKeeper.GetModuleAccount(ctx, types.RewardPoolName)
}

func (k Keeper) InflateRewardPool(ctx sdk.Context) {
	blockInflationAddr := k.accountKeeper.GetModuleAccount(ctx, auth.FeeCollectorName).GetAddress()
	blockInflation := k.bankKeeper.GetBalance(ctx, blockInflationAddr, types.StakeDenom)
	rewardPoolAllocation := k.GetParams(ctx).RewardPoolAllocation

	blockInflationDec := sdk.NewDecFromInt(blockInflation.Amount)
	rewardAmount := blockInflationDec.Mul(rewardPoolAllocation)
	rewardCoin := sdk.NewCoin(types.StakeDenom, rewardAmount.TruncateInt())

	err := k.bankKeeper.SendCoinsFromModuleToModule(
		ctx, auth.FeeCollectorName, types.RewardPoolName, sdk.NewCoins(rewardCoin))
	if err != nil {
		panic(fmt.Sprintf("Error funding reward pool: %s", err.Error()))
	}
}
