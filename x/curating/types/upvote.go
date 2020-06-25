package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewUpvote(
	curator, rewardAccount sdk.AccAddress, voteAmount, deposit sdk.Coin) Upvote {

	return Upvote{
		Curator:       curator,
		RewardAccount: rewardAccount,
		VoteAmount:    voteAmount,
		Deposit:       deposit,
	}
}
