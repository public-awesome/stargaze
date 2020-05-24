package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuy{}

// NewMsgBuy creates a new MsgBuy instance
func NewMsgBuy(amount sdk.Coin, sender sdk.AccAddress) MsgBuy {
	return MsgBuy{
		Amount: amount,
		Sender: sender,
	}
}

// nolint
func (msg MsgBuy) Route() string { return RouterKey }
func (msg MsgBuy) Type() string  { return "buy" }
func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgBuy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgBuy) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	return nil
}
