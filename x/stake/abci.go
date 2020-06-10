package stake

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/rocket-protocol/stakebird/x/stake/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker to fund reward pool on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
	k.InflateRewardPool(ctx)
}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, k Keeper) {
	// TODO
	// peek the head of the voting delegation queue
	//
	// if its greater or equal to block time, pop each one until time is different..
	// ..collate by vendor_id
	// ..collate by post_id
	// ..iterate all stakes to calculate final rewards
	// ..distribute rewards

	endTime := ctx.BlockTime()
	k.IterateVotingDelegationQueue(ctx, endTime, func(endTime time.Time, vendorID, postID, stakeID uint64, delegation stakingtypes.Delegation) bool {
		// undelegate from validator
		k.Undelegate(ctx, endTime, vendorID, postID, stakeID)

		k.RemoveFromVotingDelegationQueue(ctx, endTime, vendorID, postID, stakeID)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeVoteEnd,
				sdk.NewAttribute(sdk.AttributeKeyAmount, delegation.Shares.String()),
				sdk.NewAttribute(types.AttributeKeyVendorID, strconv.FormatUint(vendorID, 10)),
				sdk.NewAttribute(types.AttributeKeyPostID, strconv.FormatUint(postID, 10)),
				sdk.NewAttribute(types.AttributeKeyDelegator, delegation.DelegatorAddress.String()),
			),
		)

		return true
	})
}
