package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/curating/types"
)

// CreateUpvote performs an upvote operation
func (k Keeper) CreateUpvote(
	ctx sdk.Context, vendorID uint32, postID types.PostID, curator,
	rewardAccount sdk.AccAddress, voteNum int32) error {

	err := k.validateVendorID(ctx, vendorID)
	if err != nil {
		return err
	}
	if rewardAccount.Empty() {
		rewardAccount = curator
	}

	voteAmt := k.voteAmount(ctx, int64(voteNum))

	// check if there is already an upvote
	upvote, found, err := k.GetUpvote(ctx, vendorID, postID, curator)
	if err != nil {
		return err
	}
	if found {
		voteNumNew := voteNum + upvote.VoteNum
		voteAmtNew := k.voteAmount(ctx, int64(voteNumNew))

		// shadow voteAmt with the delta
		voteAmt = voteAmtNew.Sub(upvote.VoteAmount)
		// every additional upvote reward goes to the original reward account
		rewardAccount, err = sdk.AccAddressFromBech32(upvote.RewardAccount)
		if err != nil {
			return err
		}
		upvote = types.NewUpvote(
			vendorID,
			postID,
			curator,
			rewardAccount,
			voteNumNew,
			voteAmtNew,
			upvote.CuratedTime,
			ctx.BlockTime(),
		)
	} else {
		upvote = types.NewUpvote(vendorID, postID, curator, rewardAccount, voteNum, voteAmt, ctx.BlockTime(), ctx.BlockTime())
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
		// no deposit is locked
		// this curator gets both creator + curator rewards (sent to reward_account)
		_, err = k.CreatePost(
			ctx,
			vendorID,
			&postID,
			types.BodyHash{},
			"",
			nil,
			rewardAccount,
			"",
			nil,
			"",
			nil,
		)
		if err != nil {
			return err
		}
	}

	// update post totals
	post.TotalVotes += uint64(voteNum)
	post.TotalVoters++
	if !post.TotalAmount.IsValid() {
		post.TotalAmount = sdk.NewInt64Coin(voteAmt.Denom, 0)
	}
	post.TotalAmount = post.TotalAmount.Add(voteAmt)

	k.SetPost(ctx, post)

	k.SetUpvote(ctx, upvote, curator)

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
			sdk.NewAttribute(types.AttributeKeyPostID, postID.String()),
			sdk.NewAttribute(types.AttributeKeyCurator, curator.String()),
			sdk.NewAttribute(types.AttributeKeyRewardAccount, rewardAccount.String()),
			sdk.NewAttribute(types.AttributeKeyVoteNumber, fmt.Sprintf("%d", voteNum)),
			sdk.NewAttribute(types.AttributeKeyVoteAmount, voteAmt.String()),
		),
	})

	return nil
}

// SetUpvote sets a upvote in the store
func (k Keeper) SetUpvote(ctx sdk.Context, upvote types.Upvote, curator sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.UpvoteKey(upvote.VendorID, upvote.PostID, curator)
	value := k.MustMarshalUpvote(upvote)
	store.Set(key, value)
}

// GetUpvote returns an upvote if one exists
func (k Keeper) GetUpvote(
	ctx sdk.Context, vendorID uint32, postID types.PostID,
	curator sdk.AccAddress) (upvote types.Upvote, found bool, err error) {

	store := ctx.KVStore(k.storeKey)
	if err != nil {
		return upvote, false, err
	}

	key := types.UpvoteKey(vendorID, postID, curator)
	value := store.Get(key)
	if value == nil {
		return upvote, false, nil
	}
	k.cdc.MustUnmarshal(value, &upvote)

	return upvote, true, nil
}

// DeleteUpvote removes an upvote
func (k Keeper) DeleteUpvote(ctx sdk.Context, vendorID uint32, postID types.PostID, upvote types.Upvote) error {
	err := k.validateVendorID(ctx, vendorID)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	curator, err := sdk.AccAddressFromBech32(upvote.Curator)
	if err != nil {
		return err
	}
	key := types.UpvoteKey(vendorID, postID, curator)

	store.Delete(key)
	return nil
}

// voteAmount does the quadratic voting calculation
func (k Keeper) voteAmount(ctx sdk.Context, voteNum int64) sdk.Coin {
	amtPerVote := k.GetParams(ctx).VoteAmount

	amt := amtPerVote.Amount.
		MulRaw(voteNum).
		MulRaw(voteNum)

	return sdk.NewCoin(amtPerVote.Denom, amt)
}

// IterateUpvotes performs a callback function for each upvoter on a post
func (k Keeper) IterateUpvotes(
	ctx sdk.Context, vendorID uint32, postID types.PostID, cb func(upvote types.Upvote) (stop bool)) {

	store := ctx.KVStore(k.storeKey)

	// iterator over upvoters on a post
	it := sdk.KVStorePrefixIterator(store, types.UpvotePrefixKey(vendorID, postID))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var upvote types.Upvote
		k.cdc.MustUnmarshal(it.Value(), &upvote)
		if cb(upvote) {
			break
		}
	}
}
