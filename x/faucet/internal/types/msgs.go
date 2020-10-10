package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgMint{}
var _ sdk.Msg = &MsgFaucetKey{}

// msg types
const (
	TypeMsgMint      = "mint"
	TypeMsgFaucetKey = "faucet_key"
)

// NewMsgMint is a constructor function for NewMsgMint
func NewMsgMint(sender sdk.AccAddress, minter sdk.AccAddress, mTime int64, denom string) *MsgMint {
	return &MsgMint{Sender: sender.String(), Minter: minter.String(), Time: mTime, Denom: denom}
}

// Route should return the name of the module
func (msg MsgMint) Route() string { return RouterKey }

// Type should return the action
func (msg MsgMint) Type() string { return TypeMsgMint }

// ValidateBasic runs stateless checks on the message
func (msg MsgMint) ValidateBasic() error {
	if strings.TrimSpace(msg.Minter) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Minter)
	}
	if strings.TrimSpace(msg.Sender) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// NewMsgFaucetKey is a constructor function for MsgFaucetKey
func NewMsgFaucetKey(sender sdk.AccAddress, armor string) *MsgFaucetKey {
	return &MsgFaucetKey{Sender: sender.String(), Armor: armor}
}

// Route should return the name of the module
func (msg MsgFaucetKey) Route() string { return RouterKey }

// Type should return the action
func (msg MsgFaucetKey) Type() string { return TypeMsgFaucetKey }

// ValidateBasic runs stateless checks on the message
func (msg MsgFaucetKey) ValidateBasic() error {
	if strings.TrimSpace(msg.Sender) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if strings.TrimSpace(msg.Armor) == "" {
		return ErrFaucetKeyEmpty
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgFaucetKey) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgFaucetKey) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
