package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgPost{}

// msg types
const (
	TypeMsgPost   = "curating_post"
	TypeMsgUpvote = "curating_upvote"
)

// NewMsgPost creates a new MsgPost instance
func NewMsgPost(
	vendorID uint32,
	postID string,
	creator,
	rewardAccount sdk.AccAddress,
	body string,
) *MsgPost {
	return &MsgPost{
		VendorID:      vendorID,
		PostID:        postID,
		Creator:       creator.String(),
		RewardAccount: rewardAccount.String(),
		Body:          body,
	}
}

// Route implements sdk.Msg
func (msg MsgPost) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgPost) Type() string { return TypeMsgPost }

// GetSigners implements sdk.Msg
func (msg MsgPost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgPost) ValidateBasic() error {
	if msg.Body == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty body")
	}
	if strings.TrimSpace(msg.Creator) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty creator")
	}
	if strings.TrimSpace(msg.PostID) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty post_id")
	}
	if msg.VendorID < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid vendor_id")
	}

	return nil
}
