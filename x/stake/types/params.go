package types

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/x/params"
	"gopkg.in/yaml.v2"
)

// Default parameter namespace
const (
	DefaultParamspace                 = ModuleName
	DefaultVotingPeriod time.Duration = time.Hour * 24 * 3
)

// Parameter store keys
var (
	KeyVotingPeriod = []byte("VotingPeriod")
)

// ParamKeyTable for x/stake module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// Params - used for initializing default parameter for x/stake at genesis
type Params struct {
	VotingPeriod time.Duration `json:"voting_period" yaml:"voting_period"`
}

// NewParams creates a new Params object
func NewParams(votingPeriod time.Duration) Params {
	return Params{
		VotingPeriod: votingPeriod,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeyVotingPeriod, &p.VotingPeriod, validateVotingPeriod),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(DefaultVotingPeriod)
}

func (p Params) Validate() error {
	if err := validateVotingPeriod(p.VotingPeriod); err != nil {
		return err
	}

	return nil
}

func validateVotingPeriod(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("voting period must be positive: %d", v)
	}

	return nil
}
