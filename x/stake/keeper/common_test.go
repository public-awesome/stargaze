package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/rocket-protocol/stakebird/x/stake/keeper"
	"github.com/rocket-protocol/stakebird/x/stake/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// TODO
type StakeApp struct {
	*simapp.SimApp
	StakeKeeper keeper.Keeper
}

// createTestInput Returns a simapp with custom StakingKeeper
func createTestInput() (*codec.Codec, *StakeApp, sdk.Context) {
	app := &StakeApp{
		SimApp:      simapp.Setup(false),
		StakeKeeper: keeper.Keeper{},
	}
	ctx := app.BaseApp.NewContext(false, abci.Header{})

	appCodec := std.NewAppCodec(codec.New())

	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		app.GetKey(stakingtypes.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	// TODO: need to add store key to sim app?
	// SimApp doesn't have stake keeper module, so store is not loaded, no store key, etc.

	app.StakeKeeper = keeper.NewKeeper(
		appCodec,
		app.GetKey(types.StoreKey),
		app.StakingKeeper,
		nil)

	return codec.New(), app, ctx
}
