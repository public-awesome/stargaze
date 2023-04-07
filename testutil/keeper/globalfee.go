package keeper

import (
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/v9/app"
	"github.com/public-awesome/stargaze/v9/x/globalfee/keeper"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// GlobalFeeKeeper creates a testing keeper for the x/global module
func GlobalFeeKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	tStoreKey := storetypes.NewTransientStoreKey("t_globalfee")

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(tStoreKey, sdk.StoreTypeTransient, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	appCodec := encoding.Marshaler

	paramsKeeper := paramskeeper.NewKeeper(appCodec, encoding.Amino, storeKey, tStoreKey)
	paramsKeeper.Subspace(types.ModuleName).WithKeyTable(types.ParamKeyTable())
	subspace, _ := paramsKeeper.GetSubspace(types.ModuleName)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"GlobalFeeParams",
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
		GetCodeInfoFn: func(ctx sdk.Context, codeID uint64) *wasmtypes.CodeInfo {
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
		storeKey,
		paramsSubspace,
		wk,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	params := types.Params{PrivilegedAddress: []string{}}
	k.SetParams(ctx, params)

	return k, ctx
}
