package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default parameter namespace
const (
	DefaultParamspace string = ModuleName
	DefaultFunder     string = "stars1czlu4tvr3dg3ksuf8zak87eafztr2u004zyh5a"
)

// Parameter store keys
var (
	KeyFunder = []byte("Funder")
)

// NewParams creates a new Params object
func NewParams(
	funder string,
) Params {
	return Params{
		Funder: funder,
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultFunder,
	)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyFunder, &p.Funder, validateFunder),
	}
}

// Validate validates all params
func (p Params) Validate() error {
	if err := validateFunder(p.Funder); err != nil {
		return err
	}

	return nil
}

func validateFunder(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("funder can't be empty")
	}

	return nil
}
