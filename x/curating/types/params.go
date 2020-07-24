package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Default parameter namespace
const (
	DefaultParamspace                   = ModuleName
	DefaultCurationWindow time.Duration = time.Hour * 24 * 3
	DefaultMaxNumVotes    uint32        = 5
	DefaultMaxVendors     uint32        = 1
)

// 50%
var (
	DefaultPostDeposit          = sdk.NewInt64Coin(DefaultStakeDenom, 1000000)
	DefaultUpvoteDeposit        = sdk.NewInt64Coin(DefaultStakeDenom, 1000000)
	DefaultModerateDeposit      = sdk.NewInt64Coin(DefaultStakeDenom, 1000000)
	DefaultVoteAmount           = sdk.NewInt64Coin(DefaultStakeDenom, 1000000)
	DefaultRewardPoolAllocation = sdk.NewDecWithPrec(50, 2)
	DefaultCreatorAllocation    = sdk.NewDecWithPrec(50, 2)
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
	KeyCreatorAllocation    = []byte("CreatorAllocation")
)

// Params - used for initializing default parameter for stake at genesis
type Params struct {
	CurationWindow       time.Duration `json:"curation_window" yaml:"curation_window"`
	PostDeposit          sdk.Coin      `json:"post_deposit" yaml:"post_deposit"`
	UpvoteDeposit        sdk.Coin      `json:"upvote_deposit" yaml:"upvote_depsoit"`
	ModerateDeposit      sdk.Coin      `json:"moderate_deposit" yaml:"moderate_deposit"`
	VoteAmount           sdk.Coin      `json:"vote_amount" yaml:"vote_amount"`
	MaxNumVotes          uint32        `json:"max_num_votes" yaml:"max_num_votes"`
	MaxVendors           uint32        `json:"max_vendors" yaml:"max_vendors"`
	RewardPoolAllocation sdk.Dec       `json:"reward_pool_allocation" yaml:"reward_pool_allocation"`
	CreatorAllocation    sdk.Dec       `json:"creator_allocation" yaml:"creator_allocation"`
}

// NewParams creates a new Params object
func NewParams(
	curationWindow time.Duration, postDeposit, upvoteDeposit, moderateDeposit, voteAmount sdk.Coin,
	maxNumVotes, maxVendors uint32, rewardPoolAllocation, creatorAllocation sdk.Dec) Params {

	return Params{
		CurationWindow:       curationWindow,
		PostDeposit:          postDeposit,
		UpvoteDeposit:        upvoteDeposit,
		ModerateDeposit:      moderateDeposit,
		VoteAmount:           voteAmount,
		MaxNumVotes:          maxNumVotes,
		MaxVendors:           maxVendors,
		RewardPoolAllocation: rewardPoolAllocation,
		CreatorAllocation:    creatorAllocation,
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
		paramtypes.NewParamSetPair(KeyPostDeposit, &p.PostDeposit, validatePostDeposit),
		paramtypes.NewParamSetPair(KeyUpvoteDeposit, &p.UpvoteDeposit, validateUpvoteDeposit),
		paramtypes.NewParamSetPair(KeyModerateDeposit, &p.ModerateDeposit, validateModerateDeposit),
		paramtypes.NewParamSetPair(KeyVoteAmount, &p.VoteAmount, validateVoteAmount),
		paramtypes.NewParamSetPair(KeyMaxNumVotes, &p.MaxNumVotes, validateMaxNumVotes),
		paramtypes.NewParamSetPair(KeyMaxVendors, &p.MaxVendors, validateMaxVendors),
		paramtypes.NewParamSetPair(KeyRewardPoolAllocation, &p.RewardPoolAllocation, validateRewardPoolAllocation),
		paramtypes.NewParamSetPair(KeyCreatorAllocation, &p.CreatorAllocation, validateRewardPoolAllocation),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultCurationWindow, DefaultPostDeposit, DefaultUpvoteDeposit, DefaultModerateDeposit,
		DefaultVoteAmount, DefaultMaxNumVotes, DefaultMaxVendors, DefaultRewardPoolAllocation,
		DefaultCreatorAllocation)
}

// Validate validates all params
func (p Params) Validate() error {
	if err := validateCurationWindow(p.CurationWindow); err != nil {
		return err
	}
	if err := validatePostDeposit(p.PostDeposit); err != nil {
		return err
	}
	if err := validateUpvoteDeposit(p.UpvoteDeposit); err != nil {
		return err
	}
	if err := validateModerateDeposit(p.ModerateDeposit); err != nil {
		return err
	}
	if err := validateVoteAmount(p.VoteAmount); err != nil {
		return err
	}
	if err := validateMaxNumVotes(p.MaxNumVotes); err != nil {
		return err
	}
	if err := validateMaxVendors(p.MaxVendors); err != nil {
		return err
	}
	if err := validateCurationWindow(p.CurationWindow); err != nil {
		return err
	}
	if err := validateRewardPoolAllocation(p.RewardPoolAllocation); err != nil {
		return err
	}
	if err := validateCreatorAllocation(p.CreatorAllocation); err != nil {
		return err
	}

	return nil
}

func validatePostDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid post deposit: %s", v)
	}

	return nil
}

func validateUpvoteDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid upvote deposit: %s", v)
	}

	return nil
}

func validateModerateDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid moderate deposit: %s", v)
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

func validateCreatorAllocation(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsZero() {
		return fmt.Errorf("creator allocation can't be zero: %d", v)
	}

	return nil
}
