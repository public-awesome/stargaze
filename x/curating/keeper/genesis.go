package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// InitGenesis initializes the curating module state
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	k.SetParams(ctx, state.Params)

	if k.GetRewardPoolBalance(ctx).IsZero() && ctx.BlockHeight() == 0 {
		err := k.InitializeRewardPool(ctx, k.GetParams(ctx).InitialRewardPool)
		if err != nil {
			panic(err)
		}
	}

	if k.GetCreditPoolBalance(ctx).IsZero() && ctx.BlockHeight() == 0 {
		err := k.InitializeCreditPool(ctx, k.GetParams(ctx).InitialCreditPool)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis exports the curating module state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
