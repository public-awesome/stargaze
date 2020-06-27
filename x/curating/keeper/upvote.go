package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// CreateUpvote performs an upvote operation
func (k Keeper) CreateUpvote(
	ctx sdk.Context, vendorID uint32, postID string, curator,
	rewardAccount sdk.AccAddress, voteNum int32, deposit sdk.Coin) error {

	err := k.validateVendorID(ctx, vendorID)
	if err != nil {
		return err
	}
	if rewardAccount.Empty() {
		rewardAccount = curator
	}
	ud := k.GetParams(ctx).UpvoteDeposit
	if !deposit.IsEqual(ud) {
		return sdkerrors.Wrap(
			sdkerrors.ErrInsufficientFunds, fmt.Sprintf("%v != %v", deposit, ud))
	}

	// hash postID to avoid non-determinism
	postIDHash, err := hash(postID)
	if err != nil {
		return err
	}

	// check if post exist, if not, create it and start the curation period
	_, found, err := k.GetPost(ctx, vendorID, postID)
	if err != nil {
		return err
	}

	if !found {
		// pass the deposit along to the post to be locked
		// this curator gets both creator + curator rewards (sent to reward_account)
		err = k.CreatePost(ctx, vendorID, postID, "", sdk.Coin{}, nil, rewardAccount)
		if err != nil {
			return err
		}
		// shadow deposit as its no longer available
		deposit = deposit.Sub(deposit)
	} else {
		// lock deposit only if post already exists
		err = k.lockDeposit(ctx, curator, deposit)
		if err != nil {
			return err
		}
	}

	voteAmt := k.voteAmount(ctx, int64(voteNum))
	upvote := types.NewUpvote(curator, rewardAccount, voteAmt, deposit)

	store := ctx.KVStore(k.storeKey)
	key := types.UpvoteKey(vendorID, postIDHash, curator)
	value := k.cdc.MustMarshalBinaryBare(&upvote)
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
			sdk.NewAttribute(types.AttributeKeyDeposit, deposit.String()),
		),
	})

	return nil
}

// GetUpvote returns an upvote if one exists
func (k Keeper) GetUpvote(
	ctx sdk.Context, vendorID uint32, postID string,
	curator sdk.AccAddress) (upvote types.Upvote, found bool, err error) {

	store := ctx.KVStore(k.storeKey)
	postIDHash, err := hash(postID)
	if err != nil {
		return upvote, false, err
	}

	key := types.UpvoteKey(vendorID, postIDHash, curator)
	value := store.Get(key)
	if value == nil {
		return upvote, false, nil
	}
	k.cdc.MustUnmarshalBinaryBare(value, &upvote)

	return upvote, true, nil
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
	ctx sdk.Context, vendorID uint32, postIDHash []byte, cb func(upvote types.Upvote) (stop bool)) {

	store := ctx.KVStore(k.storeKey)

	// iterator over upvoters on a post
	it := sdk.KVStorePrefixIterator(store, types.UpvotePrefixKey(vendorID, postIDHash))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var upvote types.Upvote
		k.cdc.MustUnmarshalBinaryBare(it.Value(), &upvote)
		if cb(upvote) {
			break
		}
	}
}
