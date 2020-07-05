package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuy{}

// NewMsgBuy creates a new MsgBuy instance
func NewMsgBuy(maxAmount sdk.Coin, quantity uint64, sender sdk.AccAddress) MsgBuy {
	return MsgBuy{
		MaxAmount: maxAmount,
		Quantity:  quantity,
		Sender:    sender,
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
	if !msg.MaxAmount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid coins")
	}
	if msg.Quantity < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid quantity")
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	return nil
}
