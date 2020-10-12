package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgVouch{}

// msg types
const (
	TypeMsgVouch = "user_vouch"
)

// NewMsgVouch returns a new instance of MsgVouch
func NewMsgVouch(
	voucher sdk.AccAddress,
	vouched sdk.AccAddress,
	comment string,
) MsgVouch {
	return MsgVouch{
		Voucher: voucher.String(),
		Vouched: vouched.String(),
		Comment: comment,
	}
}

// Route returns the route key
func (msg MsgVouch) Route() string { return RouterKey }

// Type returns the type
func (msg MsgVouch) Type() string { return TypeMsgVouch }

// GetSigners returns the signers need to sign the msg
func (msg MsgVouch) GetSigners() []sdk.AccAddress {
	voucher, err := sdk.AccAddressFromBech32(msg.Voucher)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{voucher}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgVouch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgVouch) ValidateBasic() error {
	if strings.TrimSpace(msg.Voucher) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty voucher")
	}
	if strings.TrimSpace(msg.Vouched) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty vouched")
	}

	return nil
}
