package keeper

import (
	"context"
	"testing"

	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/globalfee/keeper"
	"github.com/public-awesome/stargaze/v17/x/globalfee/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	storemetrics "cosmossdk.io/store/metrics"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
)

// GlobalFeeKeeper creates a testing keeper for the x/global module
func GlobalFeeKeeper(tb testing.TB) (keeper.Keeper, sdk.Context) {
	tb.Helper()
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	tStoreKey := storetypes.NewTransientStoreKey("t_globalfee")

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewTestLogger(tb), storemetrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(tStoreKey, storetypes.StoreTypeTransient, db)
	require.NoError(tb, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()

	wk := MockWasmKeeper{
		HasContractInfoFn: func(_ context.Context, contractAddr sdk.AccAddress) bool {
			switch contractAddr.String() {
			case "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du":
				return true
			case "cosmos1hfml4tzwlc3mvynsg6vtgywyx00wfkhrtpkx6t":
				return true
			case "cosmos144sh8vyv5nqfylmg4mlydnpe3l4w780jsrmf4k":
				return true
			}
			return false
		},
		GetCodeInfoFn: func(_ context.Context, codeID uint64) *wasmtypes.CodeInfo {
			if codeID == 1 {
				return &wasmtypes.CodeInfo{
					Creator: "cosmos144sh8vyv5nqfylmg4mlydnpe3l4w780jsrmf4k",
				}
			}
			if codeID == 2 {
				return &wasmtypes.CodeInfo{
					Creator: "cosmos1hfml4tzwlc3mvynsg6vtgywyx00wfkhrtpkx6t",
				}
			}
			if codeID == 3 {
				return &wasmtypes.CodeInfo{
					Creator: "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du",
				}
			}
			return nil
		},
	}

	k := keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		runtime.NewKVStoreService(storeKey),
		wk,
		"cosmos1a48wdtjn3egw7swhfkeshwdtjvs6hq9nlyrwut", // random addr for gov module
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	params := types.Params{PrivilegedAddresses: []string{}}
	_ = k.SetParams(ctx, params)

	return k, ctx
}
