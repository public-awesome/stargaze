package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgSetContractAuthorization{}
	_ sdk.Msg = &MsgRemoveContractAuthorization{}
)

// msg types
const (
	TypeMsgSetContractAuthorization    = "set_contract_authorization"
	TypeMsgRemoveContractAuthorization = "remove_contract_authorization"
)

func NewMsgSetContractAuthorization(sender string, contractAddress string, methods []string) *MsgSetContractAuthorization {
	return &MsgSetContractAuthorization{
		Sender: sender,
		ContractAuthorization: &ContractAuthorization{
			ContractAddress: contractAddress,
			Methods:         methods,
		},
	}
}

func (msg MsgSetContractAuthorization) Route() string {
	return RouterKey
}

func (msg MsgSetContractAuthorization) Type() string {
	return TypeMsgRemoveCodeAuthorization
}

func (msg MsgSetContractAuthorization) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgSetContractAuthorization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetContractAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return msg.ContractAuthorization.Validate()
}

func NewMsgRemoveContractAuthorization(sender string, contractAddress string) *MsgRemoveContractAuthorization {
	return &MsgRemoveContractAuthorization{
		Sender:          sender,
		ContractAddress: contractAddress,
	}
}

func (msg MsgRemoveContractAuthorization) Route() string {
	return RouterKey
}

func (msg MsgRemoveContractAuthorization) Type() string {
	return TypeMsgRemoveCodeAuthorization
}

func (msg MsgRemoveContractAuthorization) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgRemoveContractAuthorization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgRemoveContractAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

// Validate validates the values of contract authorizations
func (ca ContractAuthorization) Validate() error {
	_, err := sdk.AccAddressFromBech32(ca.GetContractAddress())
	if err != nil {
		return err
	}

	return validateMethods(ca.Methods)
}
