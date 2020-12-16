package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgPost{}
var _ sdk.Msg = &MsgUpvote{}
var _ sdk.Msg = &MsgStake{}

// msg types
const (
	TypeMsgPost   = "curating_post"
	TypeMsgUpvote = "curating_upvote"
	TypeMsgStake  = "curating_stake"
)

// NewMsgPost creates a new MsgPost instance
func NewMsgPost(
	vendorID uint32,
	postID string,
	creator,
	rewardAccount sdk.AccAddress,
	body string,
) *MsgPost {
	return &MsgPost{
		VendorID:      vendorID,
		PostID:        postID,
		Creator:       creator.String(),
		RewardAccount: rewardAccount.String(),
		Body:          body,
	}
}

// Route implements sdk.Msg
func (msg MsgPost) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgPost) Type() string { return TypeMsgPost }

// GetSigners implements sdk.Msg
func (msg MsgPost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgPost) ValidateBasic() error {
	if msg.Body == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty body")
	}
	if strings.TrimSpace(msg.Creator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty creator")
	}
	if strings.TrimSpace(msg.PostID) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty post_id")
	}
	if msg.VendorID < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vendor_id")
	}

	return nil
}

// NewMsgUpvote creates a new MsgUpvote instance
func NewMsgUpvote(vendorID uint32, postID string, curator, rewardAccount sdk.AccAddress, voteNum int32) *MsgUpvote {
	return &MsgUpvote{
		VendorID:      vendorID,
		PostID:        postID,
		Curator:       curator.String(),
		RewardAccount: rewardAccount.String(),
		VoteNum:       voteNum,
	}
}

// Route implements sdk.Msg
func (msg MsgUpvote) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgUpvote) Type() string { return TypeMsgUpvote }

// GetSigners implements sdk.Msg
func (msg MsgUpvote) GetSigners() []sdk.AccAddress {
	curator, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{curator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUpvote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgUpvote) ValidateBasic() error {
	if strings.TrimSpace(msg.Curator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty curator")
	}
	if strings.TrimSpace(msg.PostID) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty post_id")
	}
	if msg.VendorID < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vendor_id")
	}
	if msg.VoteNum < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vote_num")
	}

	return nil
}

// NewMsgStake creates a new MsgStake instance
func NewMsgStake(vendorID uint32, postID string, delegator sdk.AccAddress, validator sdk.ValAddress, amount sdk.Int) *MsgStake {
	return &MsgStake{
		VendorID:  vendorID,
		PostID:    postID,
		Delegator: delegator.String(),
		Validator: validator.String(),
		Amount:    amount.Int64(),
	}
}

// Route implements sdk.Msg
func (msg MsgStake) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgStake) Type() string { return TypeMsgUpvote }

// GetSigners implements sdk.Msg
func (msg MsgStake) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgStake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgStake) ValidateBasic() error {
	if strings.TrimSpace(msg.Delegator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty delegator")
	}
	if strings.TrimSpace(msg.Validator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty validator")
	}
	if strings.TrimSpace(msg.PostID) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty post_id")
	}
	if msg.VendorID < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vendor_id")
	}
	if msg.Amount < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	return nil
}
