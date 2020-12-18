package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// CreateUpvote performs an upvote operation
func (k Keeper) CreateUpvote(ctx sdk.Context, vendorID uint32, postID string,
	curator, rewardAccount sdk.AccAddress, voteNum int32) error {

	err := k.validateVendorID(ctx, vendorID)
	if err != nil {
		return err
	}
	if rewardAccount.Empty() {
		rewardAccount = curator
	}

	voteAmt := k.voteAmount(ctx, int64(voteNum))

	postIDBz, err := postIDBytes(postID)
	if err != nil {
		return err
	}

	// check if post exist, if not, create it and start the curation period
	post, found, err := k.GetPost(ctx, vendorID, postID)
	if err != nil {
		return err
	}
	if found && ctx.BlockTime().After(post.CuratingEndTime) {
		return types.ErrPostExpired
	}

	if !found {
		// this curator gets both creator + curator rewards (sent to reward_account)
		err = k.CreatePost(ctx, vendorID, postID, "", nil, rewardAccount)
		if err != nil {
			return err
		}
	}

	// check if there is already an upvote, and append vote num
	upvote, found, err := k.GetUpvote(ctx, vendorID, postID, curator)
	if err != nil {
		return err
	}
	if found {
		voteNumNew := voteNum + upvote.VoteNum
		voteAmtNew := k.voteAmount(ctx, int64(voteNumNew))
		// every additional upvote reward goes to the original reward account
		rewardAccount, err = sdk.AccAddressFromBech32(upvote.RewardAccount)
		if err != nil {
			return err
		}
		upvote = types.NewUpvote(curator, rewardAccount, voteNumNew, voteAmtNew, upvote.CuratedTime, ctx.BlockTime())
	} else {
		upvote = types.NewUpvote(curator, rewardAccount, voteNum, voteAmt, ctx.BlockTime(), ctx.BlockTime())
	}

	store := ctx.KVStore(k.storeKey)
	key := types.UpvoteKey(vendorID, postIDBz, curator)
	value := k.MustMarshalUpvote(upvote)
	store.Set(key, value)

	// add vote amount to the voting pool
	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, curator, types.VotingPoolName, sdk.NewCoins(voteAmt))
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpvote,
			sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
			sdk.NewAttribute(types.AttributeKeyPostID, postID),
			sdk.NewAttribute(types.AttributeKeyCurator, curator.String()),
			sdk.NewAttribute(types.AttributeKeyRewardAccount, rewardAccount.String()),
			sdk.NewAttribute(types.AttributeKeyVoteNumber, fmt.Sprintf("%d", voteNum)),
			sdk.NewAttribute(types.AttributeKeyVoteAmount, voteAmt.String()),
		),
	})

	return nil
}

// GetUpvote returns an upvote if one exists
func (k Keeper) GetUpvote(
	ctx sdk.Context, vendorID uint32, postID string,
	curator sdk.AccAddress) (upvote types.Upvote, found bool, err error) {

	store := ctx.KVStore(k.storeKey)
	postIDBz, err := postIDBytes(postID)
	if err != nil {
		return upvote, false, err
	}

	key := types.UpvoteKey(vendorID, postIDBz, curator)
	value := store.Get(key)
	if value == nil {
		return upvote, false, nil
	}
	k.cdc.MustUnmarshalBinaryBare(value, &upvote)

	return upvote, true, nil
}

// DeleteUpvote removes an upvote
func (k Keeper) DeleteUpvote(ctx sdk.Context, vendorID uint32, postIDBz []byte, upvote types.Upvote) error {
	err := k.validateVendorID(ctx, vendorID)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	curator, err := sdk.AccAddressFromBech32(upvote.Curator)
	if err != nil {
		return err
	}
	key := types.UpvoteKey(vendorID, postIDBz, curator)

	store.Delete(key)
	return nil
}

// IterateUpvotes performs a callback function for each upvoter on a post
func (k Keeper) IterateUpvotes(
	ctx sdk.Context, vendorID uint32, postIDBz []byte, cb func(upvote types.Upvote) (stop bool)) {

	store := ctx.KVStore(k.storeKey)

	// iterator over upvoters on a post
	it := sdk.KVStorePrefixIterator(store, types.UpvotePrefixKey(vendorID, postIDBz))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var upvote types.Upvote
		k.cdc.MustUnmarshalBinaryBare(it.Value(), &upvote)
		if cb(upvote) {
			break
		}
	}
}

// voteAmount does the quadratic voting calculation
func (k Keeper) voteAmount(ctx sdk.Context, voteNum int64) sdk.Coin {
	amtPerVote := k.GetParams(ctx).VoteAmount

	amt := amtPerVote.Amount.
		MulRaw(voteNum).
		MulRaw(voteNum)

	return sdk.NewCoin(amtPerVote.Denom, amt)
}
