package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v18/app/upgrades"
	"github.com/public-awesome/stargaze/v18/internal/sdkutil"
)

// init validates every registered Fork at startup so a misconfigured entry
// (missing ChainID, missing handler, etc.) prevents the binary from booting
// instead of silently failing to fire at UpgradeHeight on chain.
func init() {
	for _, f := range Forks {
		if err := f.Validate(); err != nil {
			panic(err)
		}
	}
}

// BeginBlockForks runs in BeginBlock and applies a registered Fork when the
// current block height matches its UpgradeHeight and its ChainID matches the
// running chain (or is the wildcard upgrades.ChainIDAny). Only one fork can
// fire per block. Fork logic must be deterministic across all validators.
//
// The fork's BeginForkLogic runs inside a cache context: if it returns an
// error or panics, state changes are discarded and the error is logged so
// the chain keeps producing blocks instead of halting.
func BeginBlockForks(ctx sdk.Context, app *App) {
	for _, fork := range Forks {
		if ctx.BlockHeight() != fork.UpgradeHeight {
			continue
		}
		if fork.ChainID != upgrades.ChainIDAny && fork.ChainID != ctx.ChainID() {
			continue
		}
		err := sdkutil.ApplyFuncIfNoError(ctx, func(cacheCtx sdk.Context) error {
			return fork.BeginForkLogic(cacheCtx, app.Keepers)
		})
		if err != nil {
			ctx.Logger().Error(
				"fork failed; state changes discarded",
				"name", fork.UpgradeName,
				"chain_id", fork.ChainID,
				"height", fork.UpgradeHeight,
				"err", err,
			)
		}
		return
	}
}
