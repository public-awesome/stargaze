package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/stake/types"
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
	return &types.GenesisState{}

	// TODO
	// append stakes
}
