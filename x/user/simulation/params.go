package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

const (
	keyThresholdAmount = "threshold_amount"
)

// GetThresholdAmount randomized RewardPoolAllocation
func GetThresholdAmount(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(25))+25, 2)
}

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyThresholdAmount,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GetThresholdAmount(r))
			},
		),
	}
}
