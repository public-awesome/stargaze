package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// Stake delegates an amount to a validator and associates a post
func (k Keeper) Stake(ctx sdk.Context, vendorID uint32, postID string, delAddr sdk.AccAddress, valAddr sdk.ValAddress,
	amount sdk.Coin) error {

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return stakingtypes.ErrNoValidatorFound
	}

	_, err := k.stakingKeeper.Delegate(ctx, delAddr, amount.Amount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

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
