package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitialClaim{}

// msg types
const (
	TypeMsgInitialClaim = "initial_claim"
)

func NewMsgInitialClaim(sender string) *MsgInitialClaim {
	return &MsgInitialClaim{
		Sender: sender,
	}
}

func (msg MsgInitialClaim) Route() string {
	return RouterKey
}

func (msg MsgInitialClaim) Type() string {
	return TypeMsgInitialClaim
}

func (msg MsgInitialClaim) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgInitialClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgInitialClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
