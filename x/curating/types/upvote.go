package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewUpvote fills and Upvote struct
func NewUpvote(
	curator, rewardAccount sdk.AccAddress, voteAmount sdk.Coin, curatedTime time.Time) Upvote {

	return Upvote{
		Curator:       curator,
		RewardAccount: rewardAccount,
		VoteAmount:    voteAmount,
		CuratedTime:   curatedTime,
	}
}
