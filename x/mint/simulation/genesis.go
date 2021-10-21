package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/public-awesome/stargaze/x/mint/types"
)

// Simulation parameter constants
const (
	StartTime       = "start_time"
	StartProvisions = "start_provisions"
	ReductionFactor = "reduction_factor"
)

// GenStartTime randomized start time
func GenStartTime(r *rand.Rand) time.Time {
	return simtypes.RandTimestamp(r)
}

// GenStartProvisions randomized StartProvisions
func GenStartProvisions(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(99)), 2)
}

// GenReductionFactor randomized InflationMax
func GenReductionFactor(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(99)), 2)
}

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	// params
	var startTime time.Time
	simState.AppParams.GetOrGenerate(
		simState.Cdc, StartTime, &startTime, simState.Rand,
		func(r *rand.Rand) { startTime = GenStartTime(r) },
	)

	var startProvisions sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, StartProvisions, &startProvisions, simState.Rand,
		func(r *rand.Rand) { startProvisions = GenStartProvisions(r) },
	)

	var reductionFactor sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, ReductionFactor, &reductionFactor, simState.Rand,
		func(r *rand.Rand) { reductionFactor = GenReductionFactor(r) },
	)

	mintDenom := sdk.DefaultBondDenom
	blocksPerYear := uint64(60 * 60 * 8766 / 5)
	params := types.NewParams(mintDenom, startTime, startProvisions, reductionFactor, blocksPerYear)

	mintGenesis := types.NewGenesisState(types.InitialMinter(), params)

	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}
