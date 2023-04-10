package types

// DefaultGenesisState returns a default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: &Params{
			PrivilegedAddress: []string{},
		},
		CodeAuthorizations:     []*CodeAuthorization{},
		ContractAuthorizations: []*ContractAuthorization{},
	}
}

// Validate perform object fields validation.
func (m GenesisState) Validate() error {
	err := m.Params.Validate()
	if err != nil {
		return err
	}

	for _, ca := range m.GetCodeAuthorizations() {
		err := ca.Validate()
		if err != nil {
			return err
		}
	}

	for _, ca := range m.GetContractAuthorizations() {
		err := ca.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
