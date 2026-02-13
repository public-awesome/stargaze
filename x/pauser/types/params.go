package types

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		PrivilegedAddresses: []string{},
	}
}

// Validate validates the params.
func (p Params) Validate() error {
	return validatePrivilegedAddresses(p.PrivilegedAddresses)
}

func validatePrivilegedAddresses(i interface{}) error {
	addresses, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, addr := range addresses {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateParams{}

const (
	TypeMsgUpdateParams = "update_params"
)

func (msg MsgUpdateParams) Route() string { return RouterKey }
func (msg MsgUpdateParams) Type() string  { return TypeMsgUpdateParams }

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgUpdateParams) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgUpdateParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return msg.Params.Validate()
}
