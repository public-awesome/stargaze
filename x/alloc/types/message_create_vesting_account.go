package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateVestingAccount{}

// NewMsgCreateVestingAccount returns a reference to a NewMsgCreateVestingAccount
func NewMsgCreateVestingAccount(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins, startTime, endTime int64, delayed bool) *MsgCreateVestingAccount {
	return &MsgCreateVestingAccount{
		FromAddress: fromAddr.String(),
		ToAddress:   toAddr.String(),
		Amount:      amount,
		StartTime:   startTime,
		EndTime:     endTime,
		Delayed:     delayed,
	}
}

// ValidateBasic Implements Msg.
func (msg MsgCreateVestingAccount) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.FromAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid 'from' address: %s", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ToAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid 'to' address: %s", err)
	}

	if !msg.Amount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsAllPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if msg.StartTime <= 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid start time")
	}

	if msg.EndTime <= 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid end time")
	}

	if msg.StartTime >= msg.EndTime {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid start time")
	}

	return nil
}
