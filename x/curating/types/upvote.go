package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewUpvote fills and Upvote struct
func NewUpvote(
	curator, rewardAccount sdk.AccAddress,
	voteNum int32,
	voteAmount sdk.Coin,
	curatedTime, updatedTime time.Time) Upvote {

	return Upvote{
		Curator:       curator.String(),
		RewardAccount: rewardAccount.String(),
		VoteNum:       voteNum,
		VoteAmount:    voteAmount,
		CuratedTime:   curatedTime,
		UpdatedTime:   updatedTime,
	}
}
