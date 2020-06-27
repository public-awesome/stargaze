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
			fmt.Sprintf("Processing vendor %d post %v", post.VendorID, post.PostID))

		// return creator deposit
		err := k.RefundDeposit(ctx, post.Creator, post.Deposit)
		if err != nil {
			panic(err)
		}

		// iterate upvoters, and return their deposits
		k.IterateUpvotes(ctx, post.VendorID, post.PostID, func(upvote types.Upvote) (stop bool) {
			// return curator deposit
			err := k.RefundDeposit(ctx, upvote.Curator, upvote.Deposit)
			if err != nil {
				panic(err)
			}

			return false
		})

		return false
	})
}
