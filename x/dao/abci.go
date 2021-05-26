package dao

import (
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/dao/keeper"
	"github.com/public-awesome/stargaze/x/dao/types"
)

// EndBlocker called every block, update validator set
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// funder, err := sdk.AccAddressFromBech32(k.GetParams(ctx).Funder)
	// if err != nil {
	// 	panic(err)
	// }
	// availableFunds := bk.GetAllBalances(ctx, funder)

	// if !daoFund.Empty() {
	// 	err = dk.FundCommunityPool(ctx, daoFund, funder)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	return []abci.ValidatorUpdate{}
}
