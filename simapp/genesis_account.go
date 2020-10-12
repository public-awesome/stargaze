package simapp

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var _ authtypes.GenesisAccount = (*SimGenesisAccount)(nil)

// SimGenesisAccount defines a type that implements the GenesisAccount interface
// to be used for simulation accounts in the genesis state.
type SimGenesisAccount struct {
	*authtypes.BaseAccount

	// vesting account fields

	// total vesting coins upon initialization
	OriginalVesting sdk.Coins `json:"original_vesting" yaml:"original_vesting"` // total vesting coins upon initialization
	// delegated vested coins at time of delegation
	DelegatedFree sdk.Coins `json:"delegated_free" yaml:"delegated_free"`
	// delegated vesting coins at time of delegation
	DelegatedVesting sdk.Coins `json:"delegated_vesting" yaml:"delegated_vesting"`
	// vesting start time (UNIX Epoch time)
	StartTime int64 `json:"start_time" yaml:"start_time"`
	// vesting end time (UNIX Epoch time)
	EndTime int64 `json:"end_time" yaml:"end_time"`

	// module account fields

	// name of the module account
	ModuleName string `json:"module_name" yaml:"module_name"`
	// permissions of module account
	ModulePermissions []string `json:"module_permissions" yaml:"module_permissions"`
}

// Validate checks for errors on the vesting and module account parameters
func (sga SimGenesisAccount) Validate() error {
	if !sga.OriginalVesting.IsZero() {
		if sga.StartTime >= sga.EndTime {
			return errors.New("vesting start-time cannot be before end-time")
		}
	}

	if sga.ModuleName != "" {
		ma := authtypes.ModuleAccount{
			BaseAccount: sga.BaseAccount, Name: sga.ModuleName, Permissions: sga.ModulePermissions,
		}
		if err := ma.Validate(); err != nil {
			return err
		}
	}

	return sga.BaseAccount.Validate()
}
