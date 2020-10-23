package keeper

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
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
	ctx sdk.Context, vendorID uint32, postID string) (post types.Post, found bool, err error) {

	postIDHash, err := hash(postID)
	if err != nil {
		return post, false, err
	}

	return k.GetPostZ(ctx, vendorID, postIDHash)
}

// GetPostZ returns post if one exists
func (k Keeper) GetPostZ(
	ctx sdk.Context, vendorID uint32, postIDHash []byte) (post types.Post, found bool, err error) {

	store := ctx.KVStore(k.storeKey)

	key := types.PostKey(vendorID, postIDHash)
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
	ctx sdk.Context, vendorID uint32, postID, body string, creator, rewardAccount sdk.AccAddress) error {

	_, found, err := k.GetPost(ctx, vendorID, postID)
	if err != nil {
		return err
	}
	if found {
		return types.ErrDuplicatePost
	}

	err = k.validateVendorID(ctx, vendorID)
	if err != nil {
		return err
	}
	if rewardAccount.Empty() {
		rewardAccount = creator
	}

	postIDBz, err := postIDBytes(postID)
	if err != nil {
		return err
	}

	bodyHash, err := hash(body)
	if err != nil {
		return err
	}

	curationWindow := k.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	post := types.NewPost(vendorID, postIDBz, bodyHash, creator, rewardAccount, curationEndTime)

	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postIDBz)
	value := k.MustMarshalPost(post)
	store.Set(key, value)

	k.InsertCurationQueue(ctx, vendorID, postIDBz, curationEndTime)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePost,
			sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
			sdk.NewAttribute(types.AttributeKeyPostID, postID),
			sdk.NewAttribute(types.AttributeKeyCreator, creator.String()),
			sdk.NewAttribute(types.AttributeKeyRewardAccount, rewardAccount.String()),
			sdk.NewAttribute(types.AttributeKeyBody, body),
			sdk.NewAttribute(types.AttributeCurationEndTime, curationEndTime.Format(time.RFC3339)),
			sdk.NewAttribute(types.AttributeKeyVoteDenom, types.DefaultVoteDenom),
		),
	})

	return nil
}

// DeletePost removes a post
func (k Keeper) DeletePost(ctx sdk.Context, vendorID uint32, postIDHash []byte) error {
	err := k.validateVendorID(ctx, vendorID)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postIDHash)

	store.Delete(key)
	return nil
}

// InsertCurationQueue inserts a VPPair into the right timeslot in the curation queue
func (k Keeper) InsertCurationQueue(
	ctx sdk.Context, vendorID uint32, postID []byte, curationEndTime time.Time) {
	vpPair := types.VPPair{VendorID: vendorID, PostIDHash: postID}

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
	k.cdc.MustUnmarshalBinaryBare(bz, &vps)

	return vps.Pairs
}

// SetCurationQueueTimeSlice sets a slice of Vendor/PostIDs in the curation queue
func (k Keeper) SetCurationQueueTimeSlice(
	ctx sdk.Context, timestamp time.Time, vps []types.VPPair) {

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&types.VPPairs{Pairs: vps})
	store.Set(types.CurationQueueByTimeKey(timestamp), bz)
}

func hash(body string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(body))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// IterateExpiredPosts iterates over posts that have finished their
// curation period, and performs a callback fuction.
func (k Keeper) IterateExpiredPosts(
	ctx sdk.Context, cb func(post types.Post) (stop bool)) {

	it := k.CurationQueueIterator(ctx, ctx.BlockTime())
	defer it.Close()

	for ; it.Valid(); it.Next() {
		vps := types.VPPairs{}
		k.cdc.MustUnmarshalBinaryBare(it.Value(), &vps)
		for _, vp := range vps.Pairs {
			post, found, err := k.GetPostZ(ctx, vp.VendorID, vp.PostIDHash)
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

// postIDBytes returns the byte representation of a postID
func postIDBytes(postID string) ([]byte, error) {
	postIDInt64, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		return nil, err
	}

	postIDBz := make([]byte, 8)
	binary.BigEndian.PutUint64(postIDBz, uint64(postIDInt64))

	return postIDBz, nil
}
