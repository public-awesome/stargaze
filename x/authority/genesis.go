package authority

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v13/x/authority/keeper"
	"github.com/public-awesome/stargaze/v13/x/authority/types"
)

// InitGenesis initializes the authority module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	params := genState.Params
	err := k.SetParams(ctx, params)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the authority module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	return &types.GenesisState{
		Params: params,
	}
}
