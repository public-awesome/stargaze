package types

// DefaultGenesisState returns a default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

// Validate perform object fields validation.
func (m GenesisState) Validate() error {
	return nil
}
