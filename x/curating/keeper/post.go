package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/curating/types"
)

// GetPosts returns all posts on chain based on vendor id
func (k Keeper) GetPosts(ctx sdk.Context, vendorID uint32) (posts []types.Post) {
	k.IteratePosts(ctx, vendorID, func(post types.Post) bool {
		posts = append(posts, post)
		return false
	})
	return
}

// GetPost returns post if one exists
func (k Keeper) GetPost(
	ctx sdk.Context, vendorID uint32, postID types.PostID) (post types.Post, found bool, err error) {

	store := ctx.KVStore(k.storeKey)

	key := types.PostKey(vendorID, postID)
	value := store.Get(key)
	if value == nil {
		return post, false, nil
	}
	k.MustUnmarshalPost(value, &post)

	return post, true, nil
}

// CreatePost registers a post on-chain and starts the curation period.
// It can be called from CreateUpvote() when a post doesn't exist yet.
func (k Keeper) CreatePost(
	ctx sdk.Context,
	vendorID uint32,
	postID *types.PostID,
	bodyHash types.BodyHash,
	body string,
	creator, rewardAccount sdk.AccAddress,
	chainID string,
	contractAddress sdk.AccAddress,
	metadata string,
	parentID *types.PostID,
) (post types.Post, err error) {

	err = k.validateVendorID(ctx, vendorID)
	if err != nil {
		return post, err
	}
	err = k.validatePostBodyLength(ctx, body)
	if err != nil {
		return post, err
	}
	err = k.validatePostBodyLength(ctx, body)
	if err != nil {
		return post, err
	}
	if rewardAccount.Empty() {
		rewardAccount = creator
	}

	if postID != nil {
		_, found, err := k.GetPost(ctx, vendorID, *postID)
		if err != nil {
			return post, err
		}
		if found {
			return post, types.ErrDuplicatePost
		}
	}
	if vendorID == 0 {
		rawPostID := k.GetPostID(ctx) + 1
		k.SetPostID(ctx, rawPostID)
		id := types.PostIDFromInt64(int64(rawPostID))
		postID = &id
	}
	if postID == nil {
		return post, types.ErrInvalidPostID
	}

	curationWindow := k.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	post = types.NewPost(
		vendorID,
		*postID,
		bodyHash,
		body,
		creator,
		rewardAccount,
		curationEndTime,
		chainID,
		creator,
		contractAddress,
		metadata,
		false,
		parentID,
	)
	k.SetPost(ctx, post)
	k.InsertCurationQueue(ctx, vendorID, *postID, curationEndTime)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePost,
			sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
			sdk.NewAttribute(types.AttributeKeyPostID, postID.String()),
			sdk.NewAttribute(types.AttributeKeyCreator, creator.String()),
			sdk.NewAttribute(types.AttributeKeyRewardAccount, rewardAccount.String()),
			sdk.NewAttribute(types.AttributeKeyBodyHash, bodyHash.String()),
			sdk.NewAttribute(types.AttributeKeyBody, body),
			sdk.NewAttribute(types.AttributeCurationEndTime, curationEndTime.Format(time.RFC3339)),
			sdk.NewAttribute(types.AttributeKeyVoteDenom, types.DefaultVoteDenom),
			sdk.NewAttribute(types.AttributeKeyChainID, chainID),
			sdk.NewAttribute(types.AttributeKeyContractAddress, contractAddress.String()),
			sdk.NewAttribute(types.AttributeKeyMetadata, metadata),
			sdk.NewAttribute(types.AttributeKeyLocked, "false"),
		),
	})

	if parentID != nil {
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypePost,
				sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
				sdk.NewAttribute(types.AttributeKeyPostID, postID.String()),
				sdk.NewAttribute(types.AttributeKeyParentID, parentID.String()),
			),
		})
	}

	return post, nil
}

// SetPost sets a post in the store
func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(post.VendorID, post.PostID)
	value := k.MustMarshalPost(post)
	store.Set(key, value)
}

// DeletePost removes a post
func (k Keeper) DeletePost(ctx sdk.Context, vendorID uint32, postID types.PostID) error {
	err := k.validateVendorID(ctx, vendorID)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postID)

	store.Delete(key)
	return nil
}

// InsertCurationQueue inserts a VPPair into the right timeslot in the curation queue
func (k Keeper) InsertCurationQueue(
	ctx sdk.Context, vendorID uint32, postID types.PostID, curationEndTime time.Time) {
	vpPair := types.VPPair{VendorID: vendorID, PostID: postID}

	timeSlice := k.GetCurationQueueTimeSlice(ctx, curationEndTime)
	if len(timeSlice) == 0 {
		k.SetCurationQueueTimeSlice(ctx, curationEndTime, []types.VPPair{vpPair})

		return
	}

	timeSlice = append(timeSlice, vpPair)
	k.SetCurationQueueTimeSlice(ctx, curationEndTime, timeSlice)
}

// RemoveFromCurationQueue will remove an entire VPPair set from the queue
func (k Keeper) RemoveFromCurationQueue(ctx sdk.Context, curationEndTime time.Time) {
	ctx.KVStore(k.storeKey).Delete(types.CurationQueueByTimeKey(curationEndTime))
}

// GetCurationQueueTimeSlice returns a slice of Vendor/PostID pairs for a give time
func (k Keeper) GetCurationQueueTimeSlice(
	ctx sdk.Context, timestamp time.Time) (vpPairs []types.VPPair) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.CurationQueueByTimeKey(timestamp))
	if bz == nil {
		return []types.VPPair{}
	}

	vps := types.VPPairs{}
	k.cdc.MustUnmarshal(bz, &vps)

	return vps.Pairs
}

// SetCurationQueueTimeSlice sets a slice of Vendor/PostIDs in the curation queue
func (k Keeper) SetCurationQueueTimeSlice(
	ctx sdk.Context, timestamp time.Time, vps []types.VPPair) {

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.VPPairs{Pairs: vps})
	store.Set(types.CurationQueueByTimeKey(timestamp), bz)
}

// IterateExpiredPosts iterates over posts that have finished their
// curation period, and performs a callback fuction.
func (k Keeper) IterateExpiredPosts(
	ctx sdk.Context, cb func(post types.Post) (stop bool)) {

	it := k.CurationQueueIterator(ctx, ctx.BlockTime())
	defer it.Close()

	for ; it.Valid(); it.Next() {
		vps := types.VPPairs{}
		k.cdc.MustUnmarshal(it.Value(), &vps)
		for _, vp := range vps.Pairs {
			post, found, err := k.GetPost(ctx, vp.VendorID, vp.PostID)
			if err != nil {
				// Do want to panic here because if a post doesn't exist for an upvote
				// it means there's some kind of consensus failure, so halt the chain.
				panic(err)
			}
			if found {
				cb(post)
			}
		}
	}
}

// CurationQueueIterator returns an sdk.Iterator for all the posts
// in the queue that expire by endTime
func (k Keeper) CurationQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(
		types.KeyPrefixCurationQueue,
		sdk.PrefixEndBytes(types.CurationQueueByTimeKey(endTime)))
}

// IteratePosts iterates over the all the posts by vendor_id and performs a callback function
func (k Keeper) IteratePosts(ctx sdk.Context, vendorID uint32, cb func(post types.Post) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PostsKey(vendorID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.MustUnmarshalPost(iterator.Value(), &post)

		if cb(post) {
			break
		}
	}
}

// ----- The PostID for native posts (vendor = 0)

// GetPostID gets the highest postl ID
func (k Keeper) GetPostID(ctx sdk.Context) (postID uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PostIDKey)
	if bz == nil {
		k.SetPostID(ctx, 0)
		return 0
	}

	postID = types.GetPostIDFromBytes(bz)
	return postID
}

// SetPostID sets the new post ID to the store
func (k Keeper) SetPostID(ctx sdk.Context, postID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PostIDKey, types.GetPostIDBytes(postID))
}
