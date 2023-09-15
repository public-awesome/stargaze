package keeper_test

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	simapp "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	stargazeapp "github.com/public-awesome/stargaze/v12/app"
	"github.com/public-awesome/stargaze/v12/x/mint/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool, homePath string) (*stargazeapp.App, sdk.Context) {
	app := setup(isCheckTx, homePath)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}

func setup(isCheckTx bool, homePath string) *stargazeapp.App {
	app, genesisState := genApp(!isCheckTx, 5, homePath)
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
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func genApp(withGenesis bool, invCheckPeriod uint, homePath string) (*stargazeapp.App, stargazeapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := stargazeapp.MakeEncodingConfig()
	app := stargazeapp.NewStargazeApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		homePath,
		invCheckPeriod,
		simapp.EmptyAppOptions{},
		nil,
		wasm.DisableAllProposals,
	)

	if withGenesis {
		return app, stargazeapp.NewDefaultGenesisState(encCdc.Codec)
	}

	return app, stargazeapp.GenesisState{}
}
