package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// InitGenesis initializes the curating module state
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	k.SetParams(ctx, state.Params)
	if k.GetRewardPoolBalance(ctx).IsZero() {
		err := k.InitializeRewardPool(ctx, k.GetParams(ctx).InitialRewardPool)
		if err != nil {
			panic(err)
		}
	}

	// [TODO]
	// set posts
	// set curation queue
	// set upvotes

	for _, post := range state.Posts {
		k.SetPost(ctx, post)
		// if post is before block time, insert into queue
	}
}

// ExportGenesis exports the curating module state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
		// TODO: add reward pool
	}

	// TODO: append posts
	// TODO: append upvotes
	// TODO: append curation queue
}
