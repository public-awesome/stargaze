package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rocket-protocol/stakebird/x/stake/types"
)

// Return post if one exists for (vendor_id | post_id)
func (k Keeper) GetPost(ctx sdk.Context, vendorID, postID uint64) (post types.Post, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postID)
	value := store.Get(key)
	if value == nil {
		return post, false
	}
	k.cdc.MustUnmarshalBinaryBare(value, &post)

	return post, true
}

func (k Keeper) CreatePost(ctx sdk.Context, postID, vendorID uint64, body string, votingPeriod time.Duration) types.Post {
	// use default voting period if not specified
	if votingPeriod == 0 {
		votingPeriod = k.VotingPeriod(ctx)
	}
	voteEnd := ctx.BlockTime().Add(votingPeriod)
	post := types.NewPost(postID, vendorID, body, voteEnd)
	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postID)
	value := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(key, value)

	return post
}
