package alloc

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v18/x/alloc/keeper"
	"github.com/public-awesome/stargaze/v18/x/alloc/types"
)

// BeginBlocker to distribute specific rewards on every begin block
func BeginBlocker(goCtx context.Context, k keeper.Keeper) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	if err := k.DistributeInflation(ctx); err != nil {
		panic(fmt.Sprintf("Error distribute inflation: %s", err.Error()))
	}
}
