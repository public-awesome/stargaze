package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/std"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// createTestInput Returns a simapp with custom StakingKeeper
// to avoid messing with the hooks.
func createTestInput() (*codec.Codec, *simapp.SimApp, sdk.Context) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, abci.Header{})

	appCodec := std.NewAppCodec(codec.New())

	app.StakingKeeper = keeper.NewKeeper(
		appCodec,
		app.GetKey(staking.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(staking.ModuleName),
	)

	return codec.New(), app, ctx
}
