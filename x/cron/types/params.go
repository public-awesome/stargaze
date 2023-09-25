package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyPrivilegedAddresses = []byte("PrivilegedAddresses")
)

var _ paramtypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// default module parameters
func DefaultParams() Params {
	return Params{
		PrivilegedAddresses: []string{},
	}
}

func NewParams(addresses []string) Params {
	return Params{
		PrivilegedAddresses: addresses,
	}
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPrivilegedAddresses, &p.PrivilegedAddresses, validatePriviligedAddresses),
	}
}

func (p Params) Validate() error {
	return validatePriviligedAddresses(p.PrivilegedAddresses)
}

func validatePriviligedAddresses(i interface{}) error {
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
