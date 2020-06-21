package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
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

func (k Keeper) CreatePost(ctx sdk.Context, postID, vendorID uint64, hash string, stake sdk.Coin, creator sdk.AccAddress) types.Post {
	curationWindow := k.CurationWindow(ctx)
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	post := types.NewPost(postID, vendorID, hash, creator, stake, curationEndTime)
	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postID)
	value := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(key, value)

	return post
}
