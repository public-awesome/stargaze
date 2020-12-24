package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	curatingtypes "github.com/public-awesome/stakebird/x/curating/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

// GetStakes returns all stakes for a post
func (k Keeper) GetStakes(ctx sdk.Context, vendorID uint32, postID []byte) (stakes []types.Stake) {
	k.iterateStakes(ctx, vendorID, postID, func(stake types.Stake) bool {
		stakes = append(stakes, stake)
		return false
	})
	return
}

// GetStake returns an existing stake from storage
func (k Keeper) GetStake(ctx sdk.Context, vendorID uint32, postID []byte,
	delAddr sdk.AccAddress) (stake types.Stake, found bool, err error) {

	store := ctx.KVStore(k.storeKey)
	key := types.StakeKey(vendorID, postID, delAddr)
	value := store.Get(key)
	if value == nil {
		return stake, false, nil
	}
	k.MustUnmarshalStake(value, &stake)

	return stake, true, nil
}

// PerformStake delegates an amount to a validator and associates a post
func (k Keeper) PerformStake(ctx sdk.Context, vendorID uint32, postID []byte, delAddr sdk.AccAddress,
	valAddr sdk.ValAddress, amount sdk.Int) error {

	_, found, err := k.curatingKeeper.GetPostZ(ctx, vendorID, postID)
	if !found {
		return curatingtypes.ErrPostNotFound
	}
	if err != nil {
		return err
	}

	// TODO: check if post has expired

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return stakingtypes.ErrNoValidatorFound
	}

	_, err = k.stakingKeeper.Delegate(ctx, delAddr, amount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	stake, found, err := k.GetStake(ctx, vendorID, postID, delAddr)
	if err != nil {
		return err
	}
	key := types.StakeKey(vendorID, postID, delAddr)
	amt := amount
	if found {
		amt = stake.Amount.Add(amount)
		// shadow valAddr so we don't mix validators when adding stake
		valAddr, err = sdk.ValAddressFromBech32(stake.Validator)
		if err != nil {
			return err
		}
	}
	value := k.MustMarshalStake(types.NewStake(valAddr, amt))
	store.Set(key, value)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeStake,
			sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
			// sdk.NewAttribute(types.AttributeKeyPostID, postID),
			sdk.NewAttribute(types.AttributeKeyDelegator, delAddr.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	})

	return nil
}

// PerformUnstake delegates an amount to a validator and associates a post
func (k Keeper) PerformUnstake(ctx sdk.Context, vendorID uint32, postID []byte,
	delAddr sdk.AccAddress, amount sdk.Int) error {

	_, found, err := k.curatingKeeper.GetPostZ(ctx, vendorID, postID)
	if err != nil {
		return err
	}
	if !found {
		return curatingtypes.ErrPostNotFound
	}

	stake, found, err := k.GetStake(ctx, vendorID, postID, delAddr)
	if err != nil {
		return err
	}
	if !found {
		return types.ErrStaketNotFound
	}
	if amount.GT(stake.Amount) {
		return types.ErrAmountTooLarge
	}
	valAddr, err := sdk.ValAddressFromBech32(stake.Validator)
	if err != nil {
		return err
	}

	_, err = k.stakingKeeper.Undelegate(ctx, delAddr, valAddr, amount.ToDec())
	if err != nil {
		return err
	}

	if amount.Equal(stake.Amount) {
		k.deleteStake(ctx, vendorID, postID, delAddr)
	} else {
		store := ctx.KVStore(k.storeKey)
		key := types.StakeKey(vendorID, postID, delAddr)
		value := k.MustMarshalStake(types.NewStake(valAddr, stake.Amount.Sub(amount)))
		store.Set(key, value)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnstake,
			sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
			// sdk.NewAttribute(types.AttributeKeyPostID, postID),
			sdk.NewAttribute(types.AttributeKeyDelegator, delAddr.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	})

	return nil
}

// iterateStakes iterates over the stakes and performs a callback function
func (k Keeper) iterateStakes(ctx sdk.Context, vendorID uint32, postID []byte, cb func(post types.Stake) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostKey(vendorID, postID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var stake types.Stake
		k.MustUnmarshalStake(iterator.Value(), &stake)

		if cb(stake) {
			break
		}
	}
}

func (k Keeper) deleteStake(ctx sdk.Context, vendorID uint32, postID []byte, delAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.StakeKey(vendorID, postID, delAddr)
	store.Delete(key)
}
