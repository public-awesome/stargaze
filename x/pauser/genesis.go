package pauser

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/pauser/keeper"
	"github.com/public-awesome/stargaze/v17/x/pauser/types"
)

// InitGenesis initializes the module genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	params := genState.GetParams()
	if err := k.SetParams(ctx, params); err != nil {
		panic(err)
	}

	for _, pc := range genState.GetPausedContracts() {
		if err := k.SetPausedContract(ctx, pc); err != nil {
			panic(err)
		}
	}

	for _, pc := range genState.GetPausedCodeIds() {
		if err := k.SetPausedCodeID(ctx, pc); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis exports the module genesis for the current block.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}
	genesis.Params = params
	k.IteratePausedContracts(ctx, func(pc types.PausedContract) bool {
		genesis.PausedContracts = append(genesis.PausedContracts, pc)
		return false
	})
	k.IteratePausedCodeIDs(ctx, func(pc types.PausedCodeID) bool {
		genesis.PausedCodeIds = append(genesis.PausedCodeIds, pc)
		return false
	})
	return genesis
}
