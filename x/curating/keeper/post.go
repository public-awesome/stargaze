package keeper

import (
	"crypto/sha256"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

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
	k.cdc.MustUnmarshalBinaryBare(value, &post)

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

	// hash postID to avoid non-determinism
	postIDHash, err := hash(postID)
	if err != nil {
		return err
	}

	bodyHash, err := hash(body)
	if err != nil {
		return err
	}

	curationWindow := k.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	post := types.NewPost(
		vendorID, postIDHash, bodyHash, creator, rewardAccount, curationEndTime)

	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postIDHash)
	value := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(key, value)

	k.InsertCurationQueue(ctx, vendorID, postIDHash, curationEndTime)

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

// md5 is used over sha256 because it's faster and produces a more compact result.
// Collisions are unlikely since it's always paired with another id (vendor_id) or
// only used to verify content bodies.
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
func (k Keeper) CurationQueueIterator(
	ctx sdk.Context, endTime time.Time) sdk.Iterator {

	store := ctx.KVStore(k.storeKey)
	return store.Iterator(
		types.KeyPrefixCurationQueue,
		sdk.PrefixEndBytes(types.CurationQueueByTimeKey(endTime)))
}
