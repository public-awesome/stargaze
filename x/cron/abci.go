package cron

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v8/x/cron/contract"
	"github.com/public-awesome/stargaze/v8/x/cron/keeper"
	"github.com/public-awesome/stargaze/v8/x/cron/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper, w types.WasmKeeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	sudoMsg := contract.SudoMsg{EndBlock: &struct{}{}}
	msgBz, err := json.Marshal(sudoMsg)
	if err != nil {
		panic(err)
	}
	k.IteratePrivileged(ctx, abciContractCallback(ctx, k, w, msgBz))
	return nil
}

// returns safe method to send the message via sudo to the privileged contract
func abciContractCallback(parentCtx sdk.Context, k keeper.Keeper, w types.WasmKeeper, msgBz []byte) func(contractAddr sdk.AccAddress) bool {
	logger := keeper.ModuleLogger(parentCtx)
	return func(contractAddr sdk.AccAddress) bool {
		// any panic will crash the node, so we are better taking care of them here
		defer RecoverToLog(logger, contractAddr)()

		logger.Debug("privileged contract callback", "type", "end_blocker", "msg", string(msgBz))
		ctx, commit := parentCtx.CacheContext()

		if _, err := w.Sudo(ctx, contractAddr, msgBz); err != nil {
			logger.Error(
				"abci callback to privileged contract failed",
				"type", "end_blocker",
				"cause", err,
				"contract-address", contractAddr,
			)
			return false // return without commit
		}
		commit()
		return false
	}
}

// RecoverToLog catches panic and logs cause to error
func RecoverToLog(logger log.Logger, contractAddr sdk.AccAddress) func() {
	return func() {
		if r := recover(); r != nil {
			var cause string
			switch rType := r.(type) {
			case sdk.ErrorOutOfGas:
				cause = fmt.Sprintf("out of gas in location: %v", rType.Descriptor)
			default:
				cause = fmt.Sprintf("%s", r)
			}
			logger.
				Error("panic executing callback",
					"cause", cause,
					"contract-address", contractAddr.String(),
					"stacktrace", string(debug.Stack()),
				)
		}
	}
}
