package keeper

import (
	"context"
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	dbm "github.com/cosmos/cosmos-db"

	storemetrics "cosmossdk.io/store/metrics"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/public-awesome/stargaze/v13/x/cron/keeper"
	"github.com/public-awesome/stargaze/v13/x/cron/types"
	"github.com/stretchr/testify/require"
)

// CronKeeper creates a testing keeper for the x/cron module
func CronKeeper(tb testing.TB) (keeper.Keeper, sdk.Context) {
	tb.Helper()
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	tStoreKey := storetypes.NewTransientStoreKey("t_cron")

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewTestLogger(tb), storemetrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(tStoreKey, storetypes.StoreTypeTransient, db)
	require.NoError(tb, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsKeeper := paramskeeper.NewKeeper(cdc, types.Amino, storeKey, tStoreKey)
	paramsKeeper.Subspace(types.ModuleName).WithKeyTable(types.ParamKeyTable())
	subspace, _ := paramsKeeper.GetSubspace(types.ModuleName)

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
		SudoFn: func(_ context.Context, _ sdk.AccAddress, _ []byte) ([]byte, error) {
			return nil, nil
		},
	}

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		subspace,
		wk,
		"cosmos1a48wdtjn3egw7swhfkeshwdtjvs6hq9nlyrwut", // random addr for gov module
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	params := types.Params{AdminAddresses: []string{}}
	_ = k.SetParams(ctx, params)

	return k, ctx
}

type MockWasmKeeper struct {
	HasContractInfoFn func(ctx context.Context, contractAddr sdk.AccAddress) bool
	SudoFn            func(ctx context.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
	GetCodeInfoFn     func(ctx context.Context, codeID uint64) *wasmtypes.CodeInfo
	GetContractInfoFn func(ctx context.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}

func (k MockWasmKeeper) HasContractInfo(ctx context.Context, contractAddress sdk.AccAddress) bool {
	if k.HasContractInfoFn == nil {
		panic("not supposed to be called!")
	}
	return k.HasContractInfoFn(ctx, contractAddress)
}

func (k MockWasmKeeper) Sudo(ctx context.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error) {
	if k.SudoFn == nil {
		panic("not supposed to be called!")
	}
	return k.SudoFn(ctx, contractAddress, msg)
}

func (k MockWasmKeeper) GetCodeInfo(ctx context.Context, codeID uint64) *wasmtypes.CodeInfo {
	if k.GetCodeInfoFn == nil {
		panic("not supposed to be called!")
	}
	return k.GetCodeInfoFn(ctx, codeID)
}

func (k MockWasmKeeper) GetContractInfo(ctx context.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo {
	if k.GetContractInfoFn == nil {
		panic("not supposed to be called!")
	}
	return k.GetContractInfoFn(ctx, contractAddress)
}
