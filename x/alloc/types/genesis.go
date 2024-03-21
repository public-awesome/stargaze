package types

import (
	"encoding/json"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default alloc genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: Params{
			DistributionProportions: DistributionProportions{
				NftIncentives:    math.LegacyNewDecWithPrec(45, 2), // 45%
				DeveloperRewards: math.LegacyNewDecWithPrec(15, 2), // 15%
				CommunityPool:    math.LegacyNewDecWithPrec(5, 2),  // 5%
			},
			WeightedDeveloperRewardsReceivers:  []WeightedAddress{},
			WeightedIncentivesRewardsReceivers: []WeightedAddress{},
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

// GetGenesisStateFromAppState return GenesisState
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
