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
	Inflation        = "inflation"
	GenesisTime      = "genesis_time"
	GenesisInflation = "genesis_inflation"
	ReductionFactor  = "reduction_factor"
)

// GenInflation randomized Inflation
func GenInflation(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(99)), 2)
}

// GenGenesisTime randomized genesis time
func GenGenesisTime(r *rand.Rand) time.Time {
	return simtypes.RandTimestamp(r)
}

// GenGenesisInflation randomized GenesisInflation
func GenGenesisInflation(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(99)), 2)
}

// GenReductionFactor randomized InflationMax
func GenReductionFactor(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(99)), 2)
}

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	// minter
	var inflation sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, Inflation, &inflation, simState.Rand,
		func(r *rand.Rand) { inflation = GenInflation(r) },
	)

	// params
	var genesisTime time.Time
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GenesisTime, &genesisTime, simState.Rand,
		func(r *rand.Rand) { genesisTime = GenGenesisTime(r) },
	)

	var genesisInflation sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GenesisInflation, &genesisInflation, simState.Rand,
		func(r *rand.Rand) { genesisInflation = GenGenesisInflation(r) },
	)

	var reductionFactor sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, ReductionFactor, &reductionFactor, simState.Rand,
		func(r *rand.Rand) { reductionFactor = GenReductionFactor(r) },
	)

	mintDenom := sdk.DefaultBondDenom
	blocksPerYear := uint64(60 * 60 * 8766 / 5)
	params := types.NewParams(mintDenom, genesisTime, genesisInflation, reductionFactor, blocksPerYear)

	mintGenesis := types.NewGenesisState(types.InitialMinter(inflation), params)

	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}
