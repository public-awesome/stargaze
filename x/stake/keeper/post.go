package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rocket-protocol/stakebird/x/stake/types"
)

// Return post if one exists for (vendor_id | post_id)
func (k Keeper) GetPost(ctx sdk.Context, vendorID uint32, postID uint64) (post types.Post, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postID)
	value := store.Get(key)
	if value == nil {
		return post, false
	}
	types.ModuleCdc.MustUnmarshalBinaryLengthPrefixed(value, post)

	return post, true
}

func (k Keeper) Post(ctx sdk.Context, postID uint64, vendorID uint32, body string, votingPeriod time.Duration) {
	post := types.NewPost(postID, vendorID, body, votingPeriod, ctx.BlockTime())
	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(post)
	store.Set(key, value)
}
