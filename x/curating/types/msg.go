package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Implements the interface of `sdk.Msg`
// verify interface at compile time
var _ sdk.Msg = &MsgPost{}
var _ sdk.Msg = &MsgUpvote{}

func NewMsgPost(vendorID uint32, postID string, creator, rewardAccount sdk.AccAddress, body string, deposit sdk.Coin) MsgPost {

	return MsgPost{
		VendorID:      vendorID,
		PostID:        postID,
		Creator:       creator,
		RewardAccount: rewardAccount,
		Body:          body,
		Deposit:       deposit,
	}
}

// nolint
func (msg MsgPost) Route() string { return RouterKey }
func (msg MsgPost) Type() string  { return "curating_post" }
func (msg MsgPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgPost) ValidateBasic() error {
	// TODO: fill
	return nil
}

// ---

// NewMsgUpvote fills a MsgUpvote struct
func NewMsgUpvote(
	vendorID uint32, postID string, curator, rewardAccount sdk.AccAddress,
	voteNum int32, deposit sdk.Coin) MsgUpvote {

	return MsgUpvote{
		VendorID:      vendorID,
		PostID:        postID,
		Curator:       curator,
		RewardAccount: rewardAccount,
		VoteNum:       voteNum,
		Deposit:       deposit,
	}
}

// nolint
func (msg MsgUpvote) Route() string { return RouterKey }
func (msg MsgUpvote) Type() string  { return "curating_upvote" }
func (msg MsgUpvote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Curator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUpvote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgUpvote) ValidateBasic() error {
	// TODO: fill
	return nil
}
