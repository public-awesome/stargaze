package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgBuy{}

// MsgBuy - struct for unjailing jailed validator
type MsgBuy struct {
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
}

// NewMsgBuy creates a new MsgBuy instance
func NewMsgBuy(validatorAddr sdk.ValAddress) MsgBuy {
	return MsgBuy{
		ValidatorAddr: validatorAddr,
	}
}

const BuyConst = "Buy"

// nolint
func (msg MsgBuy) Route() string { return RouterKey }
func (msg MsgBuy) Type() string  { return BuyConst }
func (msg MsgBuy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgBuy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgBuy) ValidateBasic() error {
	if msg.ValidatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	return nil
}
