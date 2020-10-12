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
	keyRewardPoolAllocation = "reward_pool_allocation"
)

// GetRewardPoolAllocation randomized RewardPoolAllocation
func GetRewardPoolAllocation(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(25))+25, 2)
}

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyRewardPoolAllocation,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GetRewardPoolAllocation(r))
			},
		),
	}
}
