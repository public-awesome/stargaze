package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgStake{}
var _ sdk.Msg = &MsgUnstake{}

// msg types
const (
	TypeMsgStake   = "stake_stake"
	TypeMsgUnstake = "stake_unstake"
)

// NewMsgStake creates a new MsgStake instance
func NewMsgStake(
	vendorID uint32,
	postID string,
	delegator sdk.AccAddress,
	amount sdk.Int,
) *MsgStake {
	return &MsgStake{
		VendorID:  vendorID,
		PostID:    postID,
		Delegator: delegator.String(),
		Amount:    amount,
	}
}

// Route implements sdk.Msg
func (msg MsgStake) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgStake) Type() string { return TypeMsgStake }

// GetSigners implements sdk.Msg
func (msg MsgStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
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
	if strings.TrimSpace(msg.PostID) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty post_id")
	}
	if msg.VendorID < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vendor_id")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	return nil
}

// NewMsgUnstake creates a new MsgUnstake instance
func NewMsgUnstake(
	vendorID uint32,
	postID string,
	delegator sdk.AccAddress,
	amount sdk.Int,
) *MsgUnstake {
	return &MsgUnstake{
		VendorID:  vendorID,
		PostID:    postID,
		Delegator: delegator.String(),
		Amount:    amount,
	}
}

// Route implements sdk.Msg
func (msg MsgUnstake) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgUnstake) Type() string { return TypeMsgUnstake }

// GetSigners implements sdk.Msg
func (msg MsgUnstake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUnstake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgUnstake) ValidateBasic() error {
	if strings.TrimSpace(msg.Delegator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty delegator")
	}
	if strings.TrimSpace(msg.PostID) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty post_id")
	}
	if msg.VendorID < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vendor_id")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	return nil
}
