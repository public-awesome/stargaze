package keeper_test

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stargazeapp "github.com/public-awesome/stargaze/v6/app"
	"github.com/public-awesome/stargaze/v6/x/mint/types"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
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
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func genApp(withGenesis bool, invCheckPeriod uint) (*stargazeapp.App, stargazeapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := cosmoscmd.MakeEncodingConfig(stargazeapp.ModuleBasics)
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

	originalApp := app.(*stargazeapp.App)

	if withGenesis {
		return originalApp, stargazeapp.NewDefaultGenesisState(encCdc.Marshaler)
	}

	return originalApp, stargazeapp.GenesisState{}
}
