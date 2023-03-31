package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
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

	k := keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		subspace,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	params := types.Params{PrivilegedAddress: []string{}}
	k.SetParams(ctx, params)

	return k, ctx
}
