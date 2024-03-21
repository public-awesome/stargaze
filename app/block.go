package app

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PreBlocker application updates every pre block
func (a *App) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	return a.ModuleManager.PreBlock(ctx)
}

// Precommiter application updates every commit
func (a *App) Precommiter(ctx sdk.Context) {
	if err := a.ModuleManager.Precommit(ctx); err != nil {
		panic(err)
	}
}

// PrepareCheckStater application updates every commit
func (a *App) PrepareCheckStater(ctx sdk.Context) {
	if err := a.ModuleManager.PrepareCheckState(ctx); err != nil {
		panic(err)
	}
}
