package claim

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/claim/keeper"
)

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	params, err := k.Params(ctx)
	if err != nil {
		panic(err)
	}

	if !params.AirdropEnabled {
		return
	}
	// End Airdrop
	// goneTime := ctx.BlockTime().Sub(params.AirdropStartTime)
	// if goneTime > params.DurationUntilDecay+params.DurationOfDecay {
	// 	// airdrop time passed
	// 	err := k.EndAirdrop(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}
