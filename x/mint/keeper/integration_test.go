package keeper_test

import (
	"encoding/json"
	"time"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cometbft/cometbft/libs/log"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stargazeapp "github.com/public-awesome/stargaze/v10/app"
	"github.com/public-awesome/stargaze/v10/x/mint/types"
)

// DefaultConsensusParams defines the default CometBFT consensus params used in
// SimApp testing.
var DefaultConsensusParams = &cmtproto.ConsensusParams{
	Block: &cmtproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &cmtproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &cmtproto.ValidatorParams{
		PubKeyTypes: []string{
			cmttypes.ABCIPubKeyTypeEd25519,
		},
	},
}

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*stargazeapp.App, sdk.Context) {
	app := setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, cmtproto.Header{})
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
				ConsensusParams: DefaultConsensusParams,
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
		stargazeapp.EmptyAppOptions{},
		nil,
		wasm.DisableAllProposals,
	)

	originalApp := app

	if withGenesis {
		return originalApp, stargazeapp.NewDefaultGenesisState(encCdc.Codec)
	}

	return originalApp, stargazeapp.GenesisState{}
}
