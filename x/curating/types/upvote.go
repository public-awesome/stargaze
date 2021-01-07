package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewUpvote fills and Upvote struct
func NewUpvote(
	vendorID uint32, postID PostID, curator, rewardAccount sdk.AccAddress,
	voteAmount sdk.Coin, curatedTime time.Time) Upvote {

	return Upvote{
		VendorID:      vendorID,
		PostID:        postID,
		Curator:       curator.String(),
		RewardAccount: rewardAccount.String(),
		VoteAmount:    voteAmount,
		CuratedTime:   curatedTime,
	}
}

// Upvotes is a collection of Upvote objects
type Upvotes []Upvote
