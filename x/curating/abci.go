package curating

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker to fund reward pool on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
	k.InflateRewardPool(ctx)
}

// EndBlocker called after each block to process rewards
// Iterates all expired posts, processing each upvote twice:
// First upvote iteration: refund deposits, collect QV data
// Second upvote iteration: distribute QV rewards
func EndBlocker(ctx sdk.Context, k Keeper) {
	k.IterateExpiredPosts(ctx, func(post types.Post) bool {
		k.Logger(ctx).Info(
			fmt.Sprintf("Processing vendor %d post %v", post.VendorID, post.PostIDHash))

		// return creator deposit
		err := k.RefundDeposit(ctx, post.Creator, post.Deposit)
		if err != nil {
			panic(err)
		}

		qv := NewQVFData()

		// iterate upvoters, returning deposits, and tallying upvotes
		k.IterateUpvotes(ctx, post.VendorID, post.PostIDHash,
			func(upvote types.Upvote) (stop bool) {
				// return curator deposit
				err := k.RefundDeposit(ctx, upvote.Curator, upvote.Deposit)
				if err != nil {
					panic(err)
				}

				qv, err = qv.TallyVote(upvote.VoteAmount.Amount)
				if err != nil {
					panic(err)
				}

				return false
			})

		err = k.RewardCreator(ctx, post.RewardAccount, qv.MatchPool())
		if err != nil {
			panic(err)
		}

		curatorVotingReward := qv.VoterReward()
		curatorMatchReward := qv.MatchReward()

		k.IterateUpvotes(ctx, post.VendorID, post.PostIDHash,
			func(upvote types.Upvote) (stop bool) {
				// distribute quadratic voting per capita reward from voting pool
				err := k.SendVotingReward(ctx, upvote.RewardAccount, curatorVotingReward)
				if err != nil {
					panic(err)
				}

				// distribute quadratic funding reward from protocol reward pool
				err = k.SendMatchingReward(ctx, upvote.RewardAccount, curatorMatchReward)
				if err != nil {
					panic(err)
				}

				return false
			})

		return false
	})
}
