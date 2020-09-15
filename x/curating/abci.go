package curating

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker to fund reward pool on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
	if err := k.InflateRewardPool(ctx); err != nil {
		panic(fmt.Sprintf("Error funding reward pool: %s", err.Error()))
	}
}

// EndBlocker called after each block to process rewards
// Iterates all expired posts, processing each upvote twice:
// First upvote iteration: refund deposits, collect QV data
// Second upvote iteration: distribute QV rewards
func EndBlocker(ctx sdk.Context, k Keeper) {
	endTimes := make(map[time.Time]bool)
	k.IterateExpiredPosts(ctx, func(post types.Post) bool {
		k.Logger(ctx).Info(
			fmt.Sprintf("Processing vendor %d post %v", post.VendorID, post.PostIDHash))

		qv := NewQVFData(ctx, k)

		// iterate upvoters, returning deposits, and tallying upvotes
		k.IterateUpvotes(ctx, post.VendorID, post.PostIDHash,
			func(upvote types.Upvote) (stop bool) {
				var err error
				qv, err = qv.TallyVote(upvote.VoteAmount.Amount)
				if err != nil {
					panic(err)
				}

				return false
			})

		err := k.RewardCreator(ctx, post.RewardAccount, qv.MatchPool())
		if err != nil {
			panic(err)
		}

		curatorVotingReward := qv.VoterReward()
		curatorMatchReward := qv.MatchReward()

		k.IterateUpvotes(ctx, post.VendorID, post.PostIDHash,
			func(upvote types.Upvote) (stop bool) {
				// distribute quadratic voting per capita reward from voting pool
				err = k.SendVotingReward(ctx, upvote.RewardAccount, curatorVotingReward)
				if err != nil {
					panic(err)
				}

				// distribute quadratic funding reward from protocol reward pool
				err = k.SendMatchingReward(ctx, upvote.RewardAccount, curatorMatchReward)
				if err != nil {
					panic(err)
				}

				// Remove upvote
				err = k.DeleteUpvote(ctx, post.VendorID, post.PostIDHash, upvote)
				if err != nil {
					panic(err)
				}

				return false
			})

		endTimes[post.GetCuratingEndTime()] = true
		// [NOTE]: not deleting posts until we store a historical record of them (SSV)
		// https://github.com/public-awesome/stakebird/issues/194
		// err = k.DeletePost(ctx, post.VendorID, post.PostIDHash)
		// if err != nil {
		// panic(err)
		// }

		return false
	})

	// remove processed curationEndtime from queue
	for t := range endTimes {
		k.RemoveFromCurationQueue(ctx, t)
	}
}
