package keeper

import (
	"strconv"

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

	if deposit.IsLT(k.GetParams(ctx).UpvoteDeposit) {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, deposit.String())
	}

	if rewardAccount == nil {
		rewardAccount = curator
	}

	// check if post exist, if not, create it and start the curation period
	_, found := k.GetPost(ctx, vendorID, postID)
	if !found {
		// pass the deposit along to the post to be locked
		// this curator gets both creator + curator rewards
		err = k.CreatePost(ctx, vendorID, postID, "", deposit, curator, rewardAccount)
		if err != nil {
			return err
		}
	} else {
		// lock deposit only if post already exists
		err = k.lockDeposit(ctx, curator, deposit)
		if err != nil {
			return err
		}
		// deposit is no longer available
		deposit = deposit.Sub(deposit)
	}

	voteAmt := k.voteAmount(ctx, int64(voteNum))
	upvote := types.NewUpvote(curator, rewardAccount, voteAmt, deposit)

	store := ctx.KVStore(k.storeKey)
	key := types.UpvoteKey(vendorID, postID, curator)
	value := k.cdc.MustMarshalBinaryBare(&upvote)
	store.Set(key, value)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpvote,
			sdk.NewAttribute(types.AttributeKeyVendorID, strconv.FormatUint(uint64(vendorID), 10)),
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
	curator sdk.AccAddress) (upvote types.Upvote, found bool) {

	store := ctx.KVStore(k.storeKey)
	key := types.UpvoteKey(vendorID, postID, curator)
	value := store.Get(key)
	if value == nil {
		return upvote, false
	}
	k.cdc.MustUnmarshalBinaryBare(value, &upvote)

	return upvote, true
}

// voteAmount does the quadratic voting calculation
func (k Keeper) voteAmount(ctx sdk.Context, voteNum int64) sdk.Coin {
	amtPerVote := k.GetParams(ctx).VoteAmount

	amt := amtPerVote.Amount.
		MulRaw(voteNum).
		MulRaw(voteNum)

	return sdk.NewCoin(amtPerVote.Denom, amt)
}
