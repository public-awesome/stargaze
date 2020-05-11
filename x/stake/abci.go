package stake

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/davecgh/go-spew/spew"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
	// 	TODO: fill out if your application requires beginblock, if not you can delete this function
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

	spew.Dump("voting period", k.VotingPeriod(ctx))

	endTime := ctx.BlockTime()
	spew.Dump(endTime)
	k.IterateVotingDelegationQueue(ctx, endTime, func(endTime time.Time, vendorID, postID, stakeID uint64, delegation types.Delegation) bool {
		// spew.Dump(vendorID, postID, delegation)

		// undelegate from validator
		k.Undelegate(ctx, endTime, vendorID, postID, stakeID)

		k.RemoveFromVotingDelegationQueue(ctx, endTime, vendorID, postID, stakeID)

		return true
	})
}
