package stake

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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

	endTime := ctx.BlockTime()
	k.IterateVotingDelegationQueue(ctx, endTime, func(vendorID, postID uint64, delegation stakingtypes.Delegation) bool {
		spew.Dump(vendorID, postID, delegation)
		return true
	})
}
