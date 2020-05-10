package keeper_test

import (
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rocket-protocol/stakebird/x/stake/testdata"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// TODO
// type StakeApp struct {
// 	*simapp.SimApp
// 	StakeKeeper keeper.Keeper
// }

// createTestInput Returns a simapp with custom StakingKeeper
// func createTestInput() (*codec.Codec, *StakeApp, sdk.Context) {
// 	app := &StakeApp{
// 		SimApp:      simapp.Setup(false),
// 		StakeKeeper: keeper.Keeper{},
// 	}
// 	ctx := app.BaseApp.NewContext(false, abci.Header{})

// 	appCodec := std.NewAppCodec(codec.New())

// 	app.StakingKeeper = stakingkeeper.NewKeeper(
// 		appCodec,
// 		app.GetKey(stakingtypes.StoreKey),
// 		app.AccountKeeper,
// 		app.BankKeeper,
// 		app.GetSubspace(stakingtypes.ModuleName),
// 	)

// 	// TODO: need to add store key to sim app?
// 	// SimApp doesn't have stake keeper module, so store is not loaded, no store key, etc.

// 	app.StakeKeeper = keeper.NewKeeper(
// 		appCodec,
// 		app.GetKey(types.StoreKey),
// 		app.StakingKeeper,
// 		nil)

// 	return codec.New(), app, ctx
// }

func createTestInput() (*codec.Codec, *testdata.SimApp, sdk.Context) {
	db := dbm.NewMemDB()
	logger := log.NewTMJSONLogger(log.NewSyncWriter(os.Stdout))

	opts := []func(*baseapp.BaseApp){baseapp.SetPruning(store.PruneNothing)}
	app := testdata.NewSimApp(logger, db, nil, true, 0, map[int64]bool{}, "home", opts...)
	ctx := app.NewContext(false, abci.Header{})
	// appCodec := app.Codec()

	return codec.New(), app, ctx
}
