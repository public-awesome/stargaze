package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgClaimFor{}

// msg types
const (
	TypeMsgClaimFor = "claim_for"
)

func NewMsgClaimFor(sender string, address string, action Action) *MsgClaimFor {
	return &MsgClaimFor{
		Sender:  sender,
		Address: address,
		Action:  action,
	}
}

func (msg *MsgClaimFor) Route() string {
	return RouterKey
}

func (msg *MsgClaimFor) Type() string {
	return TypeMsgClaimFor
}

func (msg *MsgClaimFor) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgClaimFor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimFor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
