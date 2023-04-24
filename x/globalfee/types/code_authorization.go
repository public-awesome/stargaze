package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetCodeAuthorization{}
var _ sdk.Msg = &MsgRemoveCodeAuthorization{}

// msg types
const (
	TypeMsgSetCodeAuthorization    = "set_code_authorization"
	TypeMsgRemoveCodeAuthorization = "remove_code_authorization"
)

func NewMsgSetCodeAuthorization(sender string, codeId uint64, methods []string) *MsgSetCodeAuthorization {
	return &MsgSetCodeAuthorization{
		Sender: sender,
		CodeAuthorization: &CodeAuthorization{
			CodeId:  codeId,
			Methods: methods,
		},
	}
}

func (msg MsgSetCodeAuthorization) Route() string {
	return RouterKey
}

func (msg MsgSetCodeAuthorization) Type() string {
	return TypeMsgRemoveCodeAuthorization
}

func (msg MsgSetCodeAuthorization) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgSetCodeAuthorization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetCodeAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return msg.CodeAuthorization.Validate()
}

func NewMsgRemoveCodeAuthorization(sender string, codeId uint64) *MsgRemoveCodeAuthorization {
	return &MsgRemoveCodeAuthorization{
		Sender: sender,
		CodeId: codeId,
	}
}

func (msg MsgRemoveCodeAuthorization) Route() string {
	return RouterKey
}

func (msg MsgRemoveCodeAuthorization) Type() string {
	return TypeMsgRemoveCodeAuthorization
}

func (msg MsgRemoveCodeAuthorization) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgRemoveCodeAuthorization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgRemoveCodeAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

// Validate validates the values of code authorizations
func (ca CodeAuthorization) Validate() error {
	return validateMethods(ca.Methods)
}
