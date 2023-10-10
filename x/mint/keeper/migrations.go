package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/public-awesome/stargaze/v12/x/mint/migrations/v2"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

// Migrate1to2 migrates the x/mint module state from the consensus
// version 1 to version 2. Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/mint
// module state.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.paramSpace, m.keeper.cdc)
}
