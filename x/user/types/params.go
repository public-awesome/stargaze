package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	curatingtypes "github.com/public-awesome/stakebird/x/curating/types"
	"gopkg.in/yaml.v2"
)

// Default parameter namespace
const (
	DefaultParamspace        = ModuleName
	DefaultVouchCount uint32 = 3
)

// Default vars
var (
	DefaultThresholdAmount sdk.Coins = sdk.NewCoins(sdk.NewInt64Coin(curatingtypes.DefaultStakeDenom, 1000000))
)

// Parameter store keys
var (
	ThresholdAmount = []byte("ThresholdAmount")
	VouchCount      = []byte("VouchCount")
)

// Params - used for initializing default parameter for stake at genesis
type Params struct {
	ThresholdAmount sdk.Coins `json:"threshold_amount" yaml:"threshold_amount"`
	VouchCount      uint32    `json:"vouch_count" yaml:"vouch_count"`
}

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
		paramtypes.NewParamSetPair(ThresholdAmount, &p.ThresholdAmount, validateThresholdAmount),
		paramtypes.NewParamSetPair(VouchCount, &p.VouchCount, validateVouchCount),
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
