package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v3 "github.com/public-awesome/stargaze/v11/x/alloc/migrations/v3"
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

// Migrate1to2 migrates x/alloc state from consensus version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return nil
}

// Migrate2to3 migrates x/alloc state from consensus version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc, m.keeper.paramstore)
}
