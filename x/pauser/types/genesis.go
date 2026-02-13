package types

// DefaultGenesis returns a default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:          DefaultParams(),
		PausedContracts: []PausedContract{},
		PausedCodeIds:   []PausedCodeID{},
	}
}

// Validate performs genesis state validation.
func (m GenesisState) Validate() error {
	return m.Params.Validate()
}
