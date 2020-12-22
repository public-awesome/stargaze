package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

// InitGenesis initializes the curating module state
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
}

// ExportGenesis exports the curating module state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{}
}
