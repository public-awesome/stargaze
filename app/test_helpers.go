package app

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func setup(withGenesis bool, invCheckPeriod uint) (cosmoscmd.App, GenesisState) {
	db := dbm.NewMemDB()
	encoding := cosmoscmd.MakeEncodingConfig(ModuleBasics)

	app := NewStargazeApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, invCheckPeriod, encoding, simapp.EmptyAppOptions{})
	if withGenesis {
		return app, NewDefaultGenesisState(encoding.Marshaler)
	}
	return app, GenesisState{}
}

func Setup(isCheckTx bool) cosmoscmd.App {
	app, _ := setup(!isCheckTx, 5)

	return app
}

func MakeTestEncodingConfig() simappparams.EncodingConfig {
	encodingConfig := simappparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
