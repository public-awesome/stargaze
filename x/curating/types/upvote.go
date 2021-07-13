package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewUpvote fills and Upvote struct
func NewUpvote(
	vendorID uint32, postID PostID, curator, rewardAccount sdk.AccAddress, voteNum int32,
	voteAmount sdk.Coin, curatedTime time.Time, updatedTime time.Time) Upvote {

	return Upvote{
		VendorID:      vendorID,
		PostID:        postID,
		Curator:       curator.String(),
		RewardAccount: rewardAccount.String(),
		VoteNum:       voteNum,
		VoteAmount:    voteAmount,
		CuratedTime:   curatedTime,
		UpdatedTime:   updatedTime,
	}
}

// Upvotes is a collection of Upvote objects
type Upvotes []Upvote
