package keeper_test

import (
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rocket-protocol/stakebird/x/stake/testdata"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func createTestInput() (*codec.Codec, *testdata.SimApp, sdk.Context) {
	db := dbm.NewMemDB()
	logger := log.NewTMJSONLogger(log.NewSyncWriter(os.Stdout))

	opts := []func(*baseapp.BaseApp){baseapp.SetPruning(store.PruneNothing)}
	app := testdata.NewSimApp(logger, db, nil, true, 0, map[int64]bool{}, simapp.DefaultNodeHome, opts...)

	genesisState := testdata.ModuleBasics.DefaultGenesis(app.Codec())
	stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
	if err != nil {
		panic(err)
	}

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: time.Now()}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := app.NewContext(false, header)

	return codec.New(), app, ctx
}
