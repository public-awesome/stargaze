package v3_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stargazeapp "github.com/public-awesome/stargaze/v12/app"
	v3 "github.com/public-awesome/stargaze/v12/x/alloc/migrations/v3"
	"github.com/public-awesome/stargaze/v12/x/alloc/types"
	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/libs/log"

	dbm "github.com/cometbft/cometbft-db"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

func TestStoreMigration(t *testing.T) {
	encodingConfig := stargazeapp.MakeEncodingConfig()
	allocKey := sdk.NewKVStoreKey(types.StoreKey)
	transientAllocKey := sdk.NewTransientStoreKey("alloc_transient")
	ctx := DefaultContext(allocKey, transientAllocKey)
	paramstore := paramtypes.NewSubspace(encodingConfig.Codec, encodingConfig.Amino, allocKey, transientAllocKey, types.StoreKey)

	// check it doesn't exist before
	require.False(t, paramstore.Has(ctx, types.KeyIncentiveRewardsReceiver))
	require.False(t, paramstore.Has(ctx, types.KeySupplementAmount))

	err := v3.MigrateStore(ctx, allocKey, encodingConfig.Codec, paramstore)

	require.NoError(t, err)

	require.True(t, paramstore.Has(ctx, types.KeyIncentiveRewardsReceiver))
	require.True(t, paramstore.Has(ctx, types.KeySupplementAmount))
}

// DefaultContext creates a sdk.Context with a fresh MemDB that can be used in tests.
func DefaultContext(key storetypes.StoreKey, tkey storetypes.StoreKey) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tkey, storetypes.StoreTypeTransient, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return ctx
}
