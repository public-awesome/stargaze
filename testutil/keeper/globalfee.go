package keeper

import (
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/v9/x/globalfee/keeper"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// GlobalFeeKeeper creates a testing keeper for the x/global module
func GlobalFeeKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

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
			return nil
		},
	}

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		paramsSubspace,
		wk,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}