package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgStake{}
var _ sdk.Msg = &MsgUnstake{}
var _ sdk.Msg = &MsgBuyCreatorCoin{}

// msg types
const (
	TypeMsgStake          = "stake_stake"
	TypeMsgUnstake        = "stake_unstake"
	TypeMsgBuyCreatorCoin = "buy_creator_coin"
)

// NewMsgStake creates a new MsgStake instance
func NewMsgStake(
	vendorID uint32,
	postID string,
	delegator sdk.AccAddress,
	validator sdk.ValAddress,
	amount sdk.Int,
) *MsgStake {
	return &MsgStake{
		VendorID:  vendorID,
		PostID:    postID,
		Delegator: delegator.String(),
		Validator: validator.String(),
		Amount:    amount,
	}
}

// Route implements sdk.Msg
func (msg MsgStake) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgStake) Type() string { return TypeMsgStake }

// GetSigners implements sdk.Msg
func (msg MsgStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgStake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgStake) ValidateBasic() error {
	if strings.TrimSpace(msg.PostID) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty post_id")
	}
	if strings.TrimSpace(msg.Delegator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty delegator")
	}
	if strings.TrimSpace(msg.Validator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty validator")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}
	return nil
}

// NewMsgUnstake creates a new MsgUnstake instance
func NewMsgUnstake(
	vendorID uint32,
	postID string,
	delegator sdk.AccAddress,
	amount sdk.Int,
) *MsgUnstake {
	return &MsgUnstake{
		VendorID:  vendorID,
		PostID:    postID,
		Delegator: delegator.String(),
		Amount:    amount,
	}
}

// Route implements sdk.Msg
func (msg MsgUnstake) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgUnstake) Type() string { return TypeMsgUnstake }

// GetSigners implements sdk.Msg
func (msg MsgUnstake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUnstake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgUnstake) ValidateBasic() error {
	if strings.TrimSpace(msg.PostID) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty post_id")
	}
	if strings.TrimSpace(msg.Delegator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty delegator")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}
	return nil
}

// NewMsgBuyCreatorCoin creates a new NewMsgBuyCreatorCoin instance
func NewMsgBuyCreatorCoin(
	username string,
	creator sdk.AccAddress,
	buyer sdk.AccAddress,
	validator sdk.ValAddress,
	amount sdk.Int,
) *MsgBuyCreatorCoin {
	return &MsgBuyCreatorCoin{
		Username:  username,
		Creator:   creator.String(),
		Buyer:     buyer.String(),
		Validator: validator.String(),
		Amount:    amount,
	}
}

// Route implements sdk.Msg
func (msg MsgBuyCreatorCoin) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgBuyCreatorCoin) Type() string { return TypeMsgStake }

// GetSigners implements sdk.Msg
func (msg MsgBuyCreatorCoin) GetSigners() []sdk.AccAddress {
	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{buyer}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgBuyCreatorCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgBuyCreatorCoin) ValidateBasic() error {
	if len(strings.TrimSpace(msg.Username)) <= 3 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "username too short")
	}
	if strings.TrimSpace(msg.Creator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty creator")
	}
	if strings.TrimSpace(msg.Buyer) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty buyer")
	}
	if strings.TrimSpace(msg.Validator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty validator")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}
	return nil
}

// NewMsgSellCreatorCoin creates a new NewMsgSellCreatorCoin instance
func NewMsgSellCreatorCoin(
	username string,
	creator sdk.AccAddress,
	seller sdk.AccAddress,
	validator sdk.ValAddress,
	amount sdk.Int,
) *MsgSellCreatorCoin {
	return &MsgSellCreatorCoin{
		Username:  username,
		Creator:   creator.String(),
		Seller:    seller.String(),
		Validator: validator.String(),
		Amount:    amount,
	}
}

// Route implements sdk.Msg
func (msg MsgSellCreatorCoin) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgSellCreatorCoin) Type() string { return TypeMsgStake }

// GetSigners implements sdk.Msg
func (msg MsgSellCreatorCoin) GetSigners() []sdk.AccAddress {
	seller, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{seller}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgSellCreatorCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSellCreatorCoin) ValidateBasic() error {
	if len(strings.TrimSpace(msg.Username)) <= 3 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "username too short")
	}
	if strings.TrimSpace(msg.Creator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty creator")
	}
	if strings.TrimSpace(msg.Seller) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty seller")
	}
	if strings.TrimSpace(msg.Validator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty validator")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}
	return nil
}
