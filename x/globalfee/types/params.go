package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyPrivilegedAddresses = []byte("PrivilegedAddresses")
	KeyMinGasPrices        = []byte("MinimumGasPricesParam")
)

var _ paramtypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPrivilegedAddresses, &p.PrivilegedAddress, validatePriviligedAddresses),
		paramtypes.NewParamSetPair(KeyMinGasPrices, &p.MinimumGasPrices, validateMinimumGasPrices),
	}
}

func (p Params) Validate() error {
	if err := validatePriviligedAddresses(p.PrivilegedAddress); err != nil {
		return err
	}
	if err := validateMinimumGasPrices(p.MinimumGasPrices); err != nil {
		return err
	}
	return nil
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

func validateMinimumGasPrices(i interface{}) error {
	v, ok := i.(sdk.DecCoins)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "type: %T, expected sdk.DecCoins", i)
	}

	return v.Validate()
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
