package curating

import (
	"fmt"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/keeper"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// BeginBlocker to fund reward pool on every begin block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	if err := k.InflateRewardPool(ctx); err != nil {
		panic(fmt.Sprintf("Error funding reward pool: %s", err.Error()))
	}

}

// EndBlocker called every block, update validator set
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
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

		rewardAccount, err := sdk.AccAddressFromBech32(post.RewardAccount)
		if err != nil {
			panic(err)
		}
		err = k.RewardCreator(ctx, rewardAccount, qv.MatchPool())
		if err != nil {
			panic(err)
		}

		curatorVotingReward := qv.VoterReward()
		curatorMatchReward := qv.MatchReward()

		k.IterateUpvotes(ctx, post.VendorID, post.PostIDHash,
			func(upvote types.Upvote) (stop bool) {
				rewardAccount, err := sdk.AccAddressFromBech32(upvote.RewardAccount)
				if err != nil {
					panic(err)
				}
				// distribute quadratic voting per capita reward from voting pool
				err = k.SendVotingReward(ctx, rewardAccount, curatorVotingReward)
				if err != nil {
					panic(err)
				}

				// distribute quadratic funding reward from protocol reward pool
				err = k.SendMatchingReward(ctx, rewardAccount, curatorMatchReward)
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
	return []abci.ValidatorUpdate{}
}
