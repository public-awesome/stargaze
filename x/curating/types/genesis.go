package types

// NewGenesisState creates a new GenesisState object
func NewGenesisState(params Params) GenesisState {
	return GenesisState{
		Params: params,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}

// ValidateGenesis validates the x/stake genesis parameters
func ValidateGenesis(data GenesisState) error {
	err := data.Params.Validate()
	if err != nil {
		return err
	}

	return nil
}
