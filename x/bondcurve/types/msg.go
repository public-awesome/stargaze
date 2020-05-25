package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuy{}
var _ sdk.Msg = &MsgSell{}

// NewMsgBuy creates a new MsgBuy instance
func NewMsgBuy(amount sdk.Coin, sender sdk.AccAddress) MsgBuy {
	return MsgBuy{
		Amount: amount,
		Sender: sender,
	}
}

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
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid coins")
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	return nil
}

// NewMsgSell creates a new MsgBuy instance
func NewMsgSell(amount sdk.Coin, sender sdk.AccAddress) MsgSell {
	return MsgSell{
		Amount: amount,
		Sender: sender,
	}
}

func (msg MsgSell) Route() string { return RouterKey }
func (msg MsgSell) Type() string  { return "sell" }
func (msg MsgSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgSell) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSell) ValidateBasic() error {
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid coins")
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	return nil
}
