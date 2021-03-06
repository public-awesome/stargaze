package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Default parameter namespace
const (
	DefaultParamspace        string        = ModuleName
	DefaultCurationWindow    time.Duration = time.Minute * 10
	DefaultMaxNumVotes       uint32        = 5
	DefaultMaxVendors        uint32        = 1
	DefaultMaxPostBodyLength uint32        = 280
)

// Default vars
var (
	DefaultVoteAmount                      = sdk.NewInt64Coin(DefaultVoteDenom, 1_000_000)
	DefaultInitialRewardPool               = sdk.NewInt64Coin(DefaultStakeDenom, 21_000_000_000_000)
	DefaultRewardPoolAllocation            = sdk.NewDecWithPrec(50, 2) // from inflation
	DefaultRewardPoolCurationMaxAllocation = sdk.NewDecWithPrec(1, 3)  // .001 (0.1%)
)

// Parameter store keys
var (
	KeyCurationWindow             = []byte("CurationWindow")
	KeyVoteAmount                 = []byte("VoteAmount")
	KeyMaxNumVotes                = []byte("MaxNumVotes")
	KeyMaxVendors                 = []byte("MaxVendors")
	KeyMaxPostBodyLength          = []byte("MaxPostBodyLength")
	KeyRewardPoolAllocation       = []byte("RewardPoolAllocation")
	KeyRewardPoolCurationMaxAlloc = []byte("RewardPoolCurationMaxAlloc")
	KeyInitialRewardPool          = []byte("InitialRewardPool")
	KeyStakeDenom                 = []byte("StakeDenom")
)

// NewParams creates a new Params object
func NewParams(
	curationWindow time.Duration,
	voteAmount,
	initialRewardPool sdk.Coin,
	maxNumVotes,
	maxVendors uint32,
	maxPostBodyLength uint32,
	rewardPoolAllocation,
	rewardPoolCurationMaxAllocation sdk.Dec,
	stakeDenom string,
) Params {
	return Params{
		CurationWindow:             curationWindow,
		VoteAmount:                 voteAmount,
		InitialRewardPool:          initialRewardPool,
		MaxNumVotes:                maxNumVotes,
		MaxVendors:                 maxVendors,
		MaxPostBodyLength:          maxPostBodyLength,
		RewardPoolAllocation:       rewardPoolAllocation,
		RewardPoolCurationMaxAlloc: rewardPoolCurationMaxAllocation,
		StakeDenom:                 stakeDenom,
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
		paramtypes.NewParamSetPair(KeyCurationWindow, &p.CurationWindow, validateCurationWindow),
		paramtypes.NewParamSetPair(KeyVoteAmount, &p.VoteAmount, validateVoteAmount),
		paramtypes.NewParamSetPair(KeyInitialRewardPool, &p.InitialRewardPool, validateRewardPoolAmount),
		paramtypes.NewParamSetPair(KeyMaxNumVotes, &p.MaxNumVotes, validateMaxNumVotes),
		paramtypes.NewParamSetPair(KeyMaxVendors, &p.MaxVendors, validateMaxVendors),
		paramtypes.NewParamSetPair(KeyMaxPostBodyLength, &p.MaxPostBodyLength, validateMaxPostBodyLength),
		paramtypes.NewParamSetPair(KeyRewardPoolAllocation, &p.RewardPoolAllocation, validateRewardPoolAlloc),
		paramtypes.NewParamSetPair(KeyRewardPoolCurationMaxAlloc, &p.RewardPoolCurationMaxAlloc,
			validateRewardPoolCurationMaxAllocation),
		paramtypes.NewParamSetPair(KeyStakeDenom, &p.StakeDenom, validateDenom),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultCurationWindow,
		DefaultVoteAmount,
		DefaultInitialRewardPool,
		DefaultMaxNumVotes,
		DefaultMaxVendors,
		DefaultMaxPostBodyLength,
		DefaultRewardPoolAllocation,
		DefaultRewardPoolCurationMaxAllocation,
		DefaultStakeDenom,
	)
}

// Validate validates all params
func (p Params) Validate() error {
	if err := validateCurationWindow(p.CurationWindow); err != nil {
		return err
	}
	if err := validateVoteAmount(p.VoteAmount); err != nil {
		return err
	}
	if err := validateRewardPoolAmount(p.InitialRewardPool); err != nil {
		return err
	}
	if err := validateMaxNumVotes(p.MaxNumVotes); err != nil {
		return err
	}
	if err := validateMaxVendors(p.MaxVendors); err != nil {
		return err
	}
	if err := validateMaxPostBodyLength(p.MaxPostBodyLength); err != nil {
		return err
	}
	if err := validateCurationWindow(p.CurationWindow); err != nil {
		return err
	}
	if err := validateRewardPoolAlloc(p.RewardPoolAllocation); err != nil {
		return err
	}
	if err := validateRewardPoolCurationMaxAllocation(p.RewardPoolCurationMaxAlloc); err != nil {
		return err
	}

	return nil
}

func validateDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return fmt.Errorf("invalid denom: %s", v)
	}

	return nil
}
func validateVoteAmount(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid vote amount: %s", v)
	}

	return nil
}

func validateRewardPoolAmount(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid reward pool amount: %s", v)
	}

	return nil
}

func validateMaxNumVotes(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max num votes must be greater than or equal to 1: %d", v)
	}

	return nil
}

func validateMaxVendors(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max vendors must be greater than or equal to 1: %d", v)
	}

	return nil
}

func validateMaxPostBodyLength(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max post body length must be greater than or equal to 1: %d", v)
	}

	return nil
}

func validateCurationWindow(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("curation window must be positive: %d", v)
	}

	return nil
}

func validateRewardPoolAlloc(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsZero() {
		return fmt.Errorf("reward pool allocation can't be zero: %d", v)
	}

	return nil
}

func validateRewardPoolCurationMaxAllocation(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsZero() {
		return fmt.Errorf("reward pool curation max allocation can't be zero: %d", v)
	}

	return nil
}
