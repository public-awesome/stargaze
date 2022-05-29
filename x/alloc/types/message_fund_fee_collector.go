package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFundFeeCollector{}

// msg types
const (
	TypeMsgFundFeeCollector = "fund_fee_collector"
)

func NewMsgFundFeeCollector(sender string) *MsgFundFeeCollector {
	return &MsgFundFeeCollector{
		Sender: sender,
	}
}

func (msg MsgFundFeeCollector) Route() string {
	return RouterKey
}

func (msg MsgFundFeeCollector) Type() string {
	return TypeMsgFundFeeCollector
}

func (msg MsgFundFeeCollector) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgFundFeeCollector) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgFundFeeCollector) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
