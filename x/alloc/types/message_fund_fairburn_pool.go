package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFundFairburnPool{}

func NewMsgFundFairburnPool(sender sdk.AccAddress, amount sdk.Coins) *MsgFundFairburnPool {
	return &MsgFundFairburnPool{
		Sender: sender.String(),
		Amount: amount,
	}
}

func (msg MsgFundFairburnPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
