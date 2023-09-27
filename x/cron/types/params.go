package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyAdminAddress = []byte("AdminAddress")
)

var _ paramtypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// default module parameters
func DefaultParams() Params {
	return Params{
		AdminAddresses: []string{},
	}
}

func NewParams(addresses []string) Params {
	return Params{
		AdminAddresses: addresses,
	}
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAdminAddress, &p.AdminAddresses, validateAdminAddress),
	}
}

func (p Params) Validate() error {
	return validateAdminAddress(p.AdminAddresses)
}

func validateAdminAddress(i interface{}) error {
	privilegedAddress, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, addr := range privilegedAddress {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
