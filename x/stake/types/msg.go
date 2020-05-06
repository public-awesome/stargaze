package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgDelegate{}
var _ sdk.Msg = &MsgPost{}

type MsgPost struct {
	ID              uint64                  `json:"id" yaml:"id"`
	VendorID        uint32                  `json:"vendor_id" yaml:"vendor_id"`
	Body            string                  `json:"body" yaml:"body"`
	VotingPeriod    time.Duration           `json:"voting_period" yaml:"voting_period"`
	VotingStartTime time.Time               `json:"voting_start_time" yaml:"voting_start_time"`
	Delegation      stakingtypes.Delegation `json:"delegation" yaml:"delegation"`
}

func NewPostMsg(id uint64, vendorID uint32, body string, votingPeriod time.Duration, votingStartTime time.Time,
	delegation stakingtypes.Delegation) MsgPost {

	return MsgPost{
		ID:              id,
		VendorID:        vendorID,
		Body:            body,
		VotingPeriod:    votingPeriod,
		VotingStartTime: votingStartTime,
		Delegation:      delegation,
	}
}

// nolint
func (msg MsgPost) Route() string { return RouterKey }
func (msg MsgPost) Type() string  { return "post" }
func (msg MsgPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Delegation.DelegatorAddress}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgPost) ValidateBasic() error {
	if msg.ID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid post id")
	}
	if msg.VendorID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vendor id")
	}
	// TODO: skip body for now
	if msg.VotingPeriod < 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid voting period")
	}
	if msg.VotingStartTime.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid voting start time")
	}
	if msg.Delegation.DelegatorAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing delegator address")
	}
	return nil
}

// MsgDelegate - struct for delegating to a validator
type MsgDelegate struct {
	VendorID      uint32         `json:"vendor_id" yaml:"vendor_id"`
	PostID        uint64         `json:"post_id" yaml:"post_id"`
	DelegatorAddr sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddr sdk.ValAddress `json:"validator_address" yaml:"validator_address"`
	VotingPeriod  time.Duration  `json:"voting_period" yaml:"voting_period"`
	Amount        sdk.Coin       `json:"amount" yaml:"amount"`
}

// NewMsgDelegate creates a new MsgDelegate instance
func NewMsgDelegate(vendorID uint32, postID uint64, delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress,
	votingPeriod time.Duration, amount sdk.Coin) MsgDelegate {

	return MsgDelegate{
		VendorID:      vendorID,
		PostID:        postID,
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
		VotingPeriod:  votingPeriod,
		Amount:        amount,
	}
}

// nolint
func (msg MsgDelegate) Route() string { return RouterKey }
func (msg MsgDelegate) Type() string  { return "delegate" }
func (msg MsgDelegate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.DelegatorAddr}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgDelegate) ValidateBasic() error {
	if msg.DelegatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing delegator address")
	}
	if msg.ValidatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	if msg.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid amount")
	}
	return nil
}
