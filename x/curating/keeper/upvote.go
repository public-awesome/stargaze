package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Perform a delegation
func (k Keeper) CreateUpvote(ctx sdk.Context, vendorID uint32, postID string, delAddr sdk.AccAddress, valAddress sdk.ValAddress, amount sdk.Coin) (err error) {
	// check if post exist, if not, create it and begin the voting period
	_, found := k.GetPost(ctx, vendorID, postID)
	if !found {
		// post = k.CreatePost(ctx, postID, vendorID, "", k.VotingPeriod(ctx))
	}

	return nil
}

// func (k Keeper) getDelegation(ctx sdk.Context, endTime time.Time, vendorID, postID, stakeID uint64) stakingtypes.Delegation {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.VotingDelegationQueueKey(endTime, vendorID, postID, stakeID)
// 	value := store.Get(key)
// 	var delegation stakingtypes.Delegation
// 	k.cdc.UnmarshalBinaryBare(value, &delegation)

// 	return delegation
// }

// func (k Keeper) InsertVotingDelegationQueue(ctx sdk.Context, vendorID, postID uint64,
// 	delegation stakingtypes.Delegation, completionTime time.Time) {
// 	// get current stake index
// 	store := ctx.KVStore(k.storeKey)
// 	value := store.Get(types.KeyIndexStakeID)
// 	stakeID := types.StakeIndexFromKey(value)

// 	// create voting delegation queue key
// 	queueKey := types.VotingDelegationQueueKey(completionTime, vendorID, postID, stakeID)

// 	value = store.Get(queueKey)
// 	if len(value) == 0 {
// 		// add to queue at the right time
// 		k.setVotingDelegationQueue(ctx, queueKey, delegation)
// 	}

// 	// store incremented index
// 	store.Set(types.KeyIndexStakeID, types.StakeIndexToKey(stakeID+1))
// }

// func (k Keeper) RemoveFromVotingDelegationQueue(ctx sdk.Context, endTime time.Time, vendorID, postID, stakeID uint64) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.VotingDelegationQueueKey(endTime, vendorID, postID, stakeID)
// 	store.Delete(key)
// }

func (k Keeper) setVotingDelegationQueue(ctx sdk.Context, key []byte, delegation stakingtypes.Delegation) {
	store := ctx.KVStore(k.storeKey)
	value := k.cdc.MustMarshalBinaryBare(&delegation)
	store.Set(key, value)
}

// func (k Keeper) IterateVotingDelegationQueue(ctx sdk.Context, endTime time.Time,
// 	cb func(endTime time.Time, vendorID, postID, stakeID uint64, delegation stakingtypes.Delegation) (stop bool)) {

// 	iterator := k.VotingDelegationQueueIterator(ctx, endTime)

// 	defer iterator.Close()
// 	for ; iterator.Valid(); iterator.Next() {
// 		spew.Dump(iterator.Key())
// 		endTime, vendorID, postID, stakeID := types.SplitVotingDelegationQueueKey(iterator.Key())
// 		// spew.Dump("vendorID, postID, stakeID", vendorID, postID, stakeID)
// 		var delegation stakingtypes.Delegation
// 		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &delegation)

// 		if cb(endTime, vendorID, postID, stakeID, delegation) {
// 			break
// 		}
// 	}
// }

// Returns all delegation timeslices from time 0 until completion time
// func (k Keeper) VotingDelegationQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
// 	store := ctx.KVStore(k.storeKey)
// 	return store.Iterator(types.KeyPrefixPostCurationQueue,
// 		sdk.InclusiveEndBytes(types.VotingDelegationQueueTimeKeyPrefix(endTime)))
// }
