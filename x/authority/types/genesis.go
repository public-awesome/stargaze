package types

// DefaultGenesisState returns a default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate perform object fields validation.
func (g GenesisState) Validate() error {
	return g.Params.Validate()
}
