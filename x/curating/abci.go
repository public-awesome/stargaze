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
func EndBlocker(ctx sdk.Context, k Keeper) {
	k.IterateExpiredPosts(ctx, func(post types.Post) bool {
		k.Logger(ctx).Info(
			fmt.Sprintf("Processing vendor %d post %v", post.VendorID, post.PostIDHash))

		// return creator deposit
		err := k.RefundDeposit(ctx, post.Creator, post.Deposit)
		if err != nil {
			panic(err)
		}

		numVotes := 0
		postVotingPool := sdk.NewCoin(types.DefaultStakeDenom, sdk.ZeroInt())

		// iterate upvoters, and return their deposits
		k.IterateUpvotes(ctx, post.VendorID, post.PostIDHash, func(upvote types.Upvote) (stop bool) {
			// return curator deposit
			err := k.RefundDeposit(ctx, upvote.Curator, upvote.Deposit)
			if err != nil {
				panic(err)
			}

			numVotes++
			postVotingPool = postVotingPool.Add(upvote.VoteAmount)

			return false
		})

		curatorRewardAmount := sdk.NewCoin(
			types.DefaultStakeDenom, postVotingPool.Amount.QuoRaw(int64(numVotes)))

		// distribute quadratic voting per capita reward
		k.IterateUpvotes(ctx, post.VendorID, post.PostIDHash, func(upvote types.Upvote) (stop bool) {
			err := k.RewardAccount(ctx, upvote.RewardAccount, curatorRewardAmount)
			if err != nil {
				panic(err)
			}

			return false
		})

		return false
	})
}
