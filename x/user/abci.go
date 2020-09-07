package user

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker ...
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
}

// EndBlocker ...
func EndBlocker(ctx sdk.Context, k Keeper) {
}
