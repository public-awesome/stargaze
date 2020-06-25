package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

// Default parameter namespace
const (
	DefaultParamspace                   = ModuleName
	DefaultCurationWindow time.Duration = time.Hour * 24 * 3
	DefaultMaxNumVotes    uint16        = 5
	DefaultMaxVendors     uint16        = 1
)

// 50%
var (
	DefaultPostDeposit          = sdk.NewInt64Coin(DefaultStakeDenom, 1000000)
	DefaultUpvoteDeposit        = sdk.NewInt64Coin(DefaultStakeDenom, 1000000)
	DefaultModerateDeposit      = sdk.NewInt64Coin(DefaultStakeDenom, 1000000)
	DefaultVoteAmount           = sdk.NewInt64Coin(DefaultStakeDenom, 1000000)
	DefaultRewardPoolAllocation = sdk.NewDecWithPrec(50, 2)
)

// Parameter store keys
var (
	KeyCurationWindow       = []byte("CurationWindow")
	KeyPostDeposit          = []byte("PostDeposit")
	KeyUpvoteDeposit        = []byte("UpvoteDeposit")
	KeyModerateDeposit      = []byte("ModerateDeposit")
	KeyVoteAmount           = []byte("VoteAmount")
	KeyMaxNumVotes          = []byte("MaxNumVotes")
	KeyMaxVendors           = []byte("MaxVendors")
	KeyRewardPoolAllocation = []byte("RewardPoolAllocation")
)

// Params - used for initializing default parameter for stake at genesis
type Params struct {
	CurationWindow       time.Duration `json:"curation_window" yaml:"curation_window"`
	PostDeposit          sdk.Coin      `json:"post_deposit" yaml:"post_deposit"`
	UpvoteDeposit        sdk.Coin      `json:"upvote_deposit" yaml:"upvote_depsoit"`
	ModerateDeposit      sdk.Coin      `json:"moderate_deposit" yaml:"moderate_deposit"`
	VoteAmount           sdk.Coin      `json:"vote_amount" yaml:"vote_amount"`
	MaxNumVotes          uint16        `json:"max_num_votes" yaml:"max_num_votes"`
	MaxVendors           uint16        `json:"max_vendors" yaml:"max_vendors"`
	RewardPoolAllocation sdk.Dec       `json:"reward_pool_allocation" yaml:"reward_pool_allocation"`
}

// NewParams creates a new Params object
func NewParams(
	curationWindow time.Duration, postDeposit, upvoteDeposit, moderateDeposit, voteAmount sdk.Coin,
	maxNumVotes, maxVendors uint16, rewardPoolAllocation sdk.Dec) Params {

	return Params{
		CurationWindow:       curationWindow,
		PostDeposit:          postDeposit,
		UpvoteDeposit:        upvoteDeposit,
		ModerateDeposit:      moderateDeposit,
		VoteAmount:           voteAmount,
		MaxNumVotes:          maxNumVotes,
		MaxVendors:           maxVendors,
		RewardPoolAllocation: rewardPoolAllocation,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCurationWindow, &p.CurationWindow, validateCurationWindow),
		paramtypes.NewParamSetPair(KeyPostDeposit, &p.PostDeposit, validateCoin),
		paramtypes.NewParamSetPair(KeyUpvoteDeposit, &p.UpvoteDeposit, validateCoin),
		paramtypes.NewParamSetPair(KeyModerateDeposit, &p.ModerateDeposit, validateCoin),
		paramtypes.NewParamSetPair(KeyVoteAmount, &p.VoteAmount, validateCoin),
		paramtypes.NewParamSetPair(KeyMaxNumVotes, &p.MaxNumVotes, validateUint16),
		paramtypes.NewParamSetPair(KeyMaxVendors, &p.MaxVendors, validateUint16),
		paramtypes.NewParamSetPair(KeyRewardPoolAllocation, &p.RewardPoolAllocation, validateRewardPoolAllocation),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultCurationWindow, DefaultPostDeposit, DefaultUpvoteDeposit, DefaultModerateDeposit,
		DefaultVoteAmount, DefaultMaxNumVotes, DefaultMaxVendors, DefaultRewardPoolAllocation)
}

func (p Params) Validate() error {
	if err := validateCurationWindow(p.CurationWindow); err != nil {
		return err
	}
	if err := validateCoin(p.PostDeposit); err != nil {
		return err
	}
	if err := validateCoin(p.UpvoteDeposit); err != nil {
		return err
	}
	if err := validateCoin(p.ModerateDeposit); err != nil {
		return err
	}
	if err := validateUint16(p.MaxNumVotes); err != nil {
		return err
	}
	if err := validateUint16(p.MaxVendors); err != nil {
		return err
	}
	if err := validateCurationWindow(p.CurationWindow); err != nil {
		return err
	}

	return nil
}

func validateCoin(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	spew.Dump(v)

	if v.IsValid() {
		return fmt.Errorf("invalid coins: %v", v)
	}

	// if v.IsPositive() {
	// 	return fmt.Errorf("invalid coins: %v", v)
	// }

	return nil
}

func validateUint16(i interface{}) error {
	v, ok := i.(uint16)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("value must be greater than or equal to 1: %d", v)
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

func validateRewardPoolAllocation(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsZero() {
		return fmt.Errorf("reward pool allocation can't be zero: %d", v)
	}

	return nil
}
