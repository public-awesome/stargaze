package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	curatingtypes "github.com/public-awesome/stargaze/x/curating/types"
	"gopkg.in/yaml.v2"
)

// Default parameter namespace
const (
	DefaultParamspace string = ModuleName
	DefaultVouchCount uint32 = 3
)

// Default vars
var (
	DefaultThresholdAmount sdk.Coins = sdk.NewCoins(sdk.NewInt64Coin(curatingtypes.DefaultStakeDenom, 1000000))
)

// Parameter store keys
var (
	KeyThresholdAmount = []byte("ThresholdAmount")
	KeyVouchCount      = []byte("VouchCount")
)

// NewParams creates a new Params object
func NewParams(thresholdAmount sdk.Coins, vouchCount uint32) Params {

	return Params{
		ThresholdAmount: thresholdAmount,
		VouchCount:      vouchCount,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	out, err := yaml.Marshal(p)
	if err != nil {
		return ""
	}
	return string(out)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyThresholdAmount, &p.ThresholdAmount, validateThresholdAmount),
		paramtypes.NewParamSetPair(KeyVouchCount, &p.VouchCount, validateVouchCount),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultThresholdAmount,
		DefaultVouchCount,
	)
}

// Validate validates all params
func (p Params) Validate() error {
	if err := validateThresholdAmount(p.ThresholdAmount); err != nil {
		return err
	}
	if err := validateVouchCount(p.VouchCount); err != nil {
		return err
	}

	return nil
}

func validateThresholdAmount(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid threshold amount: %s", v)
	}

	return nil
}

func validateVouchCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("vouch count must be greater than or equal to 1: %d", v)
	}

	return nil
}
