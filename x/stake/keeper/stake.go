package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	curatingtypes "github.com/public-awesome/stargaze/x/curating/types"
	"github.com/public-awesome/stargaze/x/stake/types"
)

// GetStakes returns all stakes for a post
func (k Keeper) GetStakes(ctx sdk.Context, vendorID uint32, postID curatingtypes.PostID) (stakes []types.Stake) {
	k.iterateStakes(ctx, vendorID, postID, func(stake types.Stake) bool {
		stakes = append(stakes, stake)
		return false
	})
	return
}

// GetStake returns an existing stake from storage
func (k Keeper) GetStake(ctx sdk.Context, vendorID uint32, postID curatingtypes.PostID,
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

// BuyCreatorCoin delegates an amount to a validator and associates a post
func (k Keeper) PerformBuyCreatorCoin(
	ctx sdk.Context,
	username string,
	creator sdk.AccAddress,
	buyer sdk.AccAddress,
	valAddr sdk.ValAddress,
	amount sdk.Int) error {

	coin := sdk.NewCoin(fmt.Sprintf("@%s/%s", username, creator.String()), amount)

	stake, found, err := k.GetStake(ctx, 0, curatingtypes.PostID{}, buyer)
	if err != nil {
		return err
	}
	amt := amount
	if found {
		amt = stake.Amount.Add(amount)
		// shadow valAddr so we don't mix validators when adding stake
		valAddr, err = sdk.ValAddressFromBech32(stake.Validator)
		if err != nil {
			return err
		}
	}

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return stakingtypes.ErrNoValidatorFound
	}

	stake = types.NewStake(0, curatingtypes.PostID{}, buyer, valAddr, amt)
	k.SetStake(ctx, buyer, stake)

	_, err = k.stakingKeeper.Delegate(ctx, buyer, amount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

	if err := k.bankKeeper.MintCoins(
		ctx, types.ModuleName, sdk.NewCoins(coin),
	); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, buyer, sdk.NewCoins(coin)); err != nil {
		panic(fmt.Sprintf("unable to send coins from module to account despite previously minting coins to module account: %v", err))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBuyCreatorCoin,
			sdk.NewAttribute(types.AttributeKeyUsername, username),
			sdk.NewAttribute(types.AttributeKeyCreator, creator.String()),
			sdk.NewAttribute(types.AttributeKeyBuyer, buyer.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	})

	return nil
}

// PerformStake delegates an amount to a validator and associates a post
func (k Keeper) PerformStake(ctx sdk.Context, vendorID uint32, postID curatingtypes.PostID, delAddr sdk.AccAddress,
	valAddr sdk.ValAddress, amount sdk.Int) error {

	p, found, err := k.curatingKeeper.GetPost(ctx, vendorID, postID)
	if !found {
		return curatingtypes.ErrPostNotFound
	}
	if err != nil {
		return err
	}

	if ctx.BlockTime().Before(p.CuratingEndTime) {
		return types.ErrCurationNotExpired
	}

	stake, found, err := k.GetStake(ctx, vendorID, postID, delAddr)
	if err != nil {
		return err
	}
	amt := amount
	if found {
		amt = stake.Amount.Add(amount)
		// shadow valAddr so we don't mix validators when adding stake
		valAddr, err = sdk.ValAddressFromBech32(stake.Validator)
		if err != nil {
			return err
		}
	}

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return stakingtypes.ErrNoValidatorFound
	}

	stake = types.NewStake(vendorID, postID, delAddr, valAddr, amt)
	k.SetStake(ctx, delAddr, stake)

	_, err = k.stakingKeeper.Delegate(ctx, delAddr, amount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeStake,
			sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
			sdk.NewAttribute(types.AttributeKeyPostID, postID.String()),
			sdk.NewAttribute(types.AttributeKeyDelegator, delAddr.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amt.String()),
		),
	})

	return nil
}

// PerformUnstake delegates an amount to a validator and associates a post
func (k Keeper) PerformUnstake(ctx sdk.Context, vendorID uint32, postID curatingtypes.PostID,
	delAddr sdk.AccAddress, amount sdk.Int) error {

	_, found, err := k.curatingKeeper.GetPost(ctx, vendorID, postID)
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

	amt := stake.Amount.Sub(amount)
	if amt.Equal(sdk.ZeroInt()) {
		k.deleteStake(ctx, vendorID, postID, delAddr)
	} else {
		stake = types.NewStake(vendorID, postID, delAddr, valAddr, amt)
		k.SetStake(ctx, delAddr, stake)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnstake,
			sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
			sdk.NewAttribute(types.AttributeKeyPostID, postID.String()),
			sdk.NewAttribute(types.AttributeKeyDelegator, delAddr.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amt.String()),
		),
	})

	return nil
}

// SetStake sets the stake in the store
func (k Keeper) SetStake(ctx sdk.Context, delAddr sdk.AccAddress, s types.Stake) {
	store := ctx.KVStore(k.storeKey)
	key := types.StakeKey(s.VendorID, s.PostID, delAddr)
	value := k.MustMarshalStake(s)
	store.Set(key, value)
}

// iterateStakes iterates over the stakes and performs a callback function
func (k Keeper) iterateStakes(
	ctx sdk.Context, vendorID uint32, postID curatingtypes.PostID,
	cb func(post types.Stake) (stop bool)) {

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

func (k Keeper) deleteStake(ctx sdk.Context, vendorID uint32, postID curatingtypes.PostID, delAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.StakeKey(vendorID, postID, delAddr)
	store.Delete(key)
}
