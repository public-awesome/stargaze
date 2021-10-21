package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/public-awesome/stargaze/x/mint/types"
)

const (
	keyStartTime       = "StartTime"
	keyStartProvisions = "StartProvisions"
	keyReductionFactor = "ReductionFactor"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyStartTime,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenStartTime(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyStartProvisions,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenStartProvisions(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyReductionFactor,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenReductionFactor(r))
			},
		),
	}
}
