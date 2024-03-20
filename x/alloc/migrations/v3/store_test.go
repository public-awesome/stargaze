package v3_test

import (
	"testing"

	"cosmossdk.io/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stargazeapp "github.com/public-awesome/stargaze/v14/app"
	v3 "github.com/public-awesome/stargaze/v14/x/alloc/migrations/v3"
	"github.com/public-awesome/stargaze/v14/x/alloc/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"

	storemetrics "cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
)

func TestStoreMigration(t *testing.T) {
	encodingConfig := stargazeapp.MakeEncodingConfig()
	allocKey := storetypes.NewKVStoreKey(types.StoreKey)
	transientAllocKey := storetypes.NewTransientStoreKey("alloc_transient")
	ctx := DefaultContext(t, allocKey, transientAllocKey)
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
func DefaultContext(t *testing.T, key storetypes.StoreKey, tkey storetypes.StoreKey) sdk.Context {
	t.Helper()
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db, log.NewTestLogger(t), storemetrics.NewNoOpMetrics())
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tkey, storetypes.StoreTypeTransient, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return ctx
}
