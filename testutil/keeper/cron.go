package keeper

import (
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/v12/x/cron/keeper"
	"github.com/public-awesome/stargaze/v12/x/cron/types"
	"github.com/stretchr/testify/require"
)

// CronKeeper creates a testing keeper for the x/cron module
func CronKeeper(tb testing.TB) (keeper.Keeper, sdk.Context) {
	tb.Helper()
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(tb, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CronParams",
	)
	wk := MockWasmKeeper{
		HasContractInfoFn: func(ctx sdk.Context, contractAddr sdk.AccAddress) bool {
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
		SudoFn: func(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error) {
			return nil, nil
		},
	}

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		wk,
		"cosmos1a48wdtjn3egw7swhfkeshwdtjvs6hq9nlyrwut", // random addr for gov module
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}

type MockWasmKeeper struct {
	HasContractInfoFn func(ctx sdk.Context, contractAddr sdk.AccAddress) bool
	SudoFn            func(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
	GetCodeInfoFn     func(ctx sdk.Context, codeID uint64) *wasmtypes.CodeInfo
	GetContractInfoFn func(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}

func (k MockWasmKeeper) HasContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) bool {
	if k.HasContractInfoFn == nil {
		panic("not supposed to be called!")
	}
	return k.HasContractInfoFn(ctx, contractAddress)
}

func (k MockWasmKeeper) Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error) {
	if k.SudoFn == nil {
		panic("not supposed to be called!")
	}
	return k.SudoFn(ctx, contractAddress, msg)
}

func (k MockWasmKeeper) GetCodeInfo(ctx sdk.Context, codeID uint64) *wasmtypes.CodeInfo {
	if k.GetCodeInfoFn == nil {
		panic("not supposed to be called!")
	}
	return k.GetCodeInfoFn(ctx, codeID)
}

func (k MockWasmKeeper) GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo {
	if k.GetContractInfoFn == nil {
		panic("not supposed to be called!")
	}
	return k.GetContractInfoFn(ctx, contractAddress)
}
