package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgDelegate{}
var _ sdk.Msg = &MsgPost{}

type MsgPost struct {
	*MsgDelegate

	Body         string        `json:"body" yaml:"body"`
	VotingPeriod time.Duration `json:"voting_period" yaml:"voting_period"`
}

func NewMsgPost(vendorID, postID uint64, body string, delAddr sdk.AccAddress,
	valAddr sdk.ValAddress, amount sdk.Coin, votingPeriod time.Duration) MsgPost {

	return MsgPost{
		MsgDelegate: &MsgDelegate{
			VendorID:      vendorID,
			PostID:        postID,
			DelegatorAddr: delAddr,
			ValidatorAddr: valAddr,
			Amount:        amount,
		},
		Body:         body,
		VotingPeriod: votingPeriod,
	}
}

// nolint
func (msg MsgPost) Route() string { return RouterKey }
func (msg MsgPost) Type() string  { return "post" }
func (msg MsgPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.DelegatorAddr}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgPost) ValidateBasic() error {
	if msg.PostID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid post id")
	}
	if msg.VendorID == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vendor id")
	}
	// TODO: skip body for now
	if msg.VotingPeriod < 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid voting period")
	}
	if msg.DelegatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing delegator address")
	}
	return nil
}

// MsgDelegate - struct for delegating to a validator
type MsgDelegate struct {
	VendorID      uint64         `json:"vendor_id" yaml:"vendor_id"`
	PostID        uint64         `json:"post_id" yaml:"post_id"`
	DelegatorAddr sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddr sdk.ValAddress `json:"validator_address" yaml:"validator_address"`
	VotingPeriod  time.Duration  `json:"voting_period" yaml:"voting_period"`
	Amount        sdk.Coin       `json:"amount" yaml:"amount"`
}

// NewMsgDelegate creates a new MsgDelegate instance
func NewMsgDelegate(vendorID, postID uint64, delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress, amount sdk.Coin) MsgDelegate {

	return MsgDelegate{
		VendorID:      vendorID,
		PostID:        postID,
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
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
