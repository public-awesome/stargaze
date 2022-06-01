package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/public-awesome/stargaze/v6/x/claim/types"
)

// RandomizedGenState generates a random GenesisState  for claim
func RandomizedGenState(simState *module.SimulationState) {
	claimGenesis := types.DefaultGenesis()
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(claimGenesis)
}
