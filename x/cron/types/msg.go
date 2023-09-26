package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgPromoteToPrivilegedContract{}
	_ sdk.Msg = &MsgDemoteFromPrivilegedContract{}
)

// msg types
const (
	TypeMsgPromoteToPrivilegedContract = "promote_to_privileged_contract"
	TypeMsgRemoveCodeAuthorization     = "demote_from_privileged_contract"
)

func NewMsgPromoteToPrivilegedContract(sender string, contractAddr string) *MsgPromoteToPrivilegedContract {
	return &MsgPromoteToPrivilegedContract{
		Authority: sender,
		Contract:  contractAddr,
	}
}

func (msg MsgPromoteToPrivilegedContract) Route() string {
	return RouterKey
}

func (msg MsgPromoteToPrivilegedContract) Type() string {
	return TypeMsgRemoveCodeAuthorization
}

func (msg MsgPromoteToPrivilegedContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgPromoteToPrivilegedContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgPromoteToPrivilegedContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Contract)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address (%s)", err)
	}
	return nil
}

func NewMsgDemoteFromPrivilegedContract(sender string, contractAddr string) *MsgDemoteFromPrivilegedContract {
	return &MsgDemoteFromPrivilegedContract{
		Authority: sender,
		Contract:  contractAddr,
	}
}

func (msg MsgDemoteFromPrivilegedContract) Route() string {
	return RouterKey
}

func (msg MsgDemoteFromPrivilegedContract) Type() string {
	return TypeMsgRemoveCodeAuthorization
}

func (msg MsgDemoteFromPrivilegedContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgDemoteFromPrivilegedContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDemoteFromPrivilegedContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Contract)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address (%s)", err)
	}
	return nil
}
