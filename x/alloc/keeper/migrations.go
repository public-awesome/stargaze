package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
func (m Migrator) Migrate1to2(_ sdk.Context) error {
	return nil
}

// Migrate2to3 migrates x/alloc state from consensus version 2 to 3.
func (m Migrator) Migrate2to3(_ sdk.Context) error {
	return nil // v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc, m.keeper.paramstore)
}

// Migrate3to4 migrates the x/alloc module state from the consensus
// version 3 to version 4. Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/alloc
// module state.
func (m Migrator) Migrate3to4(_ sdk.Context) error {
	return nil // v4.MigrateStore(ctx, m.keeper.storeKey, m.keeper.paramstore, m.keeper.cdc)
}
