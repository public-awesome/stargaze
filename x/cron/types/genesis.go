package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PrivilegedContractAddresses: []string{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for _, addr := range gs.GetPrivilegedContractAddresses() {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
