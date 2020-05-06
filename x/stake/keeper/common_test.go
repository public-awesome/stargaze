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
	stakeKeeper keeper.Keeper
}

// createTestInput Returns a simapp with custom StakingKeeper
func createTestInput() (*codec.Codec, *StakeApp, sdk.Context) {
	app := &StakeApp{
		SimApp:      simapp.Setup(false),
		stakeKeeper: keeper.Keeper{},
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

	app.stakeKeeper = keeper.NewKeeper(appCodec, types.StoreKey, app.StakingKeeper, nil)

	return codec.New(), app, ctx
}
