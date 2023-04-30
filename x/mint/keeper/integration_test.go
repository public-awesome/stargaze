package keeper_test

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/simapp"
	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stargazeapp "github.com/public-awesome/stargaze/v10/app"
	"github.com/public-awesome/stargaze/v10/x/mint/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*stargazeapp.App, sdk.Context) {
	app := setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}

func setup(isCheckTx bool) *stargazeapp.App {
	app, genesisState := genApp(!isCheckTx, 5)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: tmproto.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func genApp(withGenesis bool, invCheckPeriod uint) (*stargazeapp.App, stargazeapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := stargazeapp.MakeTestEncodingConfig()
	app := stargazeapp.NewStargazeApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		simapp.DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		simapp.EmptyAppOptions{},
		nil,
		wasm.DisableAllProposals,
	)

	originalApp := app

	if withGenesis {
		return originalApp, stargazeapp.NewDefaultGenesisState(encCdc.Marshaler)
	}

	return originalApp, stargazeapp.GenesisState{}
}
