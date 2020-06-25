package keeper

import (
	"crypto/sha1"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// GetPost returns post if one exists
func (k Keeper) GetPost(
	ctx sdk.Context, vendorID uint32, postID string) (post types.Post, found bool) {

	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postID)
	value := store.Get(key)
	if value == nil {
		return post, false
	}
	k.cdc.MustUnmarshalBinaryBare(value, &post)

	return post, true
}

// CreatePost registers a post on-chain and starts the curation period
func (k Keeper) CreatePost(
	ctx sdk.Context, vendorID uint32, postID, body string, deposit sdk.Coin,
	creator, rewardAccount sdk.AccAddress) (post types.Post, err error) {

	if rewardAccount == nil {
		rewardAccount = creator
	}

	bodyHash, err := encodeBody(body)
	if err != nil {
		return post, err
	}

	err = k.lockDeposit(ctx, creator, deposit)
	if err != nil {
		return post, err
	}

	curationWindow := k.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	post = types.NewPost(bodyHash, creator, rewardAccount, deposit, curationEndTime)

	store := ctx.KVStore(k.storeKey)
	key := types.PostKey(vendorID, postID)
	value := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(key, value)

	k.InsertCurationQueue(ctx, vendorID, postID, curationEndTime)

	return post, nil
}

// InsertCurationQueue inserts a VPPair into the right timeslot in the curation queue
func (k Keeper) InsertCurationQueue(
	ctx sdk.Context, vendorID uint32, postID string, curationEndTime time.Time) {
	vpPair := types.VPPair{vendorID, postID}

	timeSlice := k.GetCurationQueueTimeSlice(ctx, curationEndTime)
	if len(timeSlice) == 0 {
		k.SetCurationQueueTimeSlice(ctx, curationEndTime, []types.VPPair{vpPair})
	} else {
		timeSlice := append(timeSlice, vpPair)
		k.SetCurationQueueTimeSlice(ctx, curationEndTime, timeSlice)
	}
}

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

func (k Keeper) SetCurationQueueTimeSlice(
	ctx sdk.Context, timestamp time.Time, vps []types.VPPair) {

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&types.VPPairs{vps})
	store.Set(types.CurationQueueByTimeKey(timestamp), bz)
}

func encodeBody(body string) ([]byte, error) {
	h := sha1.New()
	_, err := h.Write([]byte(body))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
