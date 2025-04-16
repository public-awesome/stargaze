package alloc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/alloc/keeper"
	"github.com/public-awesome/stargaze/v17/x/alloc/types"
)

// InitGenesis initializes the alloc module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
	k.GetModuleAccount(ctx, types.FairburnPoolName)
	k.GetModuleAccount(ctx, types.SupplementPoolName)
	err := k.FundCommunityPool(ctx)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the alloc module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}
	return &types.GenesisState{
		Params: params,
	}
}
