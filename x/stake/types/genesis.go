package types

// DefaultGenesisState - default GenesisState
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// Validate validates the x/curating genesis parameters
func (gs GenesisState) Validate() error {
	return nil
}
