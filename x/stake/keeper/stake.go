package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

// GetStake returns an existing stake from storage
func (k Keeper) GetStake(ctx sdk.Context, vendorID uint32, postID []byte, delAddr sdk.AccAddress) (stake types.Stake, found bool, err error) {
	store := ctx.KVStore(k.storeKey)
	key := types.StakeKey(vendorID, postID, delAddr)
	value := store.Get(key)
	if value == nil {
		return stake, false, nil
	}
	k.MustUnmarshalStake(value, &stake)

	return stake, true, nil
}

// CreateStake delegates an amount to a validator and associates a post
func (k Keeper) CreateStake(ctx sdk.Context, vendorID uint32, postID []byte, delAddr sdk.AccAddress,
	valAddr sdk.ValAddress, amount sdk.Int) error {

	_, found, err := k.curatingKeeper.GetPostZ(ctx, vendorID, postID)
	if !found {
		return types.ErrPostNotFound
	}
	if err != nil {
		return err
	}

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
	key := types.StakeKey(vendorID, postID, delAddr)
	amt := amount
	if found {
		amt = stake.Amount.Add(amount)
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
