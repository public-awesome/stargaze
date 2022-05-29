package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFundFairburnPool{}

// msg types
const (
	TypeMsgFundFairburnPool = "fund_fairburn_pool"
)

func NewMsgFundFairburnPool(sender sdk.AccAddress, amount sdk.Coins) *MsgFundFairburnPool {
	return &MsgFundFairburnPool{
		Sender: sender.String(),
		Amount: amount,
	}
}

func (msg MsgFundFairburnPool) Route() string {
	return RouterKey
}

func (msg MsgFundFairburnPool) Type() string {
	return TypeMsgFundFairburnPool
}

func (msg MsgFundFairburnPool) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgFundFairburnPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgFundFairburnPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
