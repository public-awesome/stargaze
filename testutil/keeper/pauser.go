package keeper

import (
	"context"
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v18/x/pauser/keeper"
	"github.com/public-awesome/stargaze/v18/x/pauser/types"
	"github.com/stretchr/testify/require"
)

// Known test addresses used by MockWasmKeeper
const (
	TestContract1 = "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du"
	TestContract2 = "cosmos1hfml4tzwlc3mvynsg6vtgywyx00wfkhrtpkx6t"
	TestContract3 = "cosmos144sh8vyv5nqfylmg4mlydnpe3l4w780jsrmf4k"
)

// PauserKeeper creates a testing keeper for the x/pauser module
func PauserKeeper(tb testing.TB) (keeper.Keeper, sdk.Context) {
	tb.Helper()
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	tStoreKey := storetypes.NewTransientStoreKey("t_pauser")

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewTestLogger(tb), storemetrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tStoreKey, storetypes.StoreTypeTransient, db)
	require.NoError(tb, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	wk := MockWasmKeeper{
		HasContractInfoFn: func(_ context.Context, contractAddr sdk.AccAddress) bool {
			switch contractAddr.String() {
			case TestContract1, TestContract2, TestContract3:
				return true
			}
			return false
		},
		GetCodeInfoFn: func(_ context.Context, codeID uint64) *wasmtypes.CodeInfo {
			switch codeID {
			case 1:
				return &wasmtypes.CodeInfo{Creator: TestContract3}
			case 2:
				return &wasmtypes.CodeInfo{Creator: TestContract2}
			case 3:
				return &wasmtypes.CodeInfo{Creator: TestContract1}
			}
			return nil
		},
		GetContractInfoFn: func(_ context.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo {
			switch contractAddress.String() {
			case TestContract1:
				return &wasmtypes.ContractInfo{CodeID: 1}
			case TestContract2:
				return &wasmtypes.ContractInfo{CodeID: 2}
			case TestContract3:
				return &wasmtypes.ContractInfo{CodeID: 3}
			}
			return nil
		},
	}

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		"cosmos1a48wdtjn3egw7swhfkeshwdtjvs6hq9nlyrwut", // random addr for gov module
	)
	k.SetWasmKeeper(wk)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	params := types.Params{PrivilegedAddresses: []string{}}
	_ = k.SetParams(ctx, params)

	return k, ctx
}
