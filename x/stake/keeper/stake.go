package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

// CreateStake delegates an amount to a validator and associates a post
func (k Keeper) CreateStake(ctx sdk.Context, vendorID uint32, postID string, delAddr sdk.AccAddress,
	valAddr sdk.ValAddress, amount sdk.Int) error {

	_, found, err := k.curationKeeper.GetPost(ctx, vendorID, postID)
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

	_, err := k.stakingKeeper.Delegate(ctx, delAddr, amount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

	// Stake := types.NewStake(vendorID, StakeIDBz, bodyHash, creator, rewardAccount, curationEndTime)
	// store := ctx.KVStore(k.storeKey)
	// key := types.StakeKey(vendorID, StakeIDBz)
	// value := k.MustMarshalStake(Stake)
	// store.Set(key, value)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePost,
			sdk.NewAttribute(types.AttributeKeyVendorID, fmt.Sprintf("%d", vendorID)),
			sdk.NewAttribute(types.AttributeKeyPostID, postID),
			sdk.NewAttribute(types.AttributeKeyDelegator, delAddr.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, valAddr.String()),
			sdk.NewAttribute(types.AttributeKeyStakeAmount, amount.String()),
		),
	})

	return nil
}

// DeleteStake removes a Stake
// func (k Keeper) DeleteStake(ctx sdk.Context, vendorID uint32, StakeIDBz []byte) error {
// 	err := k.validateVendorID(ctx, vendorID)
// 	if err != nil {
// 		return err
// 	}

// 	store := ctx.KVStore(k.storeKey)
// 	key := types.StakeKey(vendorID, StakeIDBz)

// 	store.Delete(key)
// 	return nil
// }
