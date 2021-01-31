package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	curatingtypes "github.com/public-awesome/stargaze/x/curating/types"
	"github.com/public-awesome/stargaze/x/stake/types"
)

// InitGenesis initializes the curating module state
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	for _, stake := range state.Stakes {
		delAddr, err := sdk.AccAddressFromBech32(stake.Delegator)
		if err != nil {
			panic(err)
		}

		k.SetStake(ctx, delAddr, stake)
	}
}

// ExportGenesis exports the curating module state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var stakes types.Stakes

	// since we only have twitter = 1 for now
	vendorID := uint32(1)
	k.curatingKeeper.IteratePosts(ctx, vendorID, func(post curatingtypes.Post) bool {
		stakes = append(stakes, k.GetStakes(ctx, vendorID, post.PostID)...)

		return false
	})

	return &types.GenesisState{
		Stakes: stakes,
	}
}
