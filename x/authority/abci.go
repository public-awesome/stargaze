package authority

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govV1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/public-awesome/stargaze/v13/x/authority/keeper"
	"github.com/public-awesome/stargaze/v13/x/authority/types"
)

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	logger := keeper.Logger(ctx)

	keeper.IterateActiveProposalsQueue(ctx, func(proposal govV1.Proposal) bool {
		var logMsg, tagValue string

		var (
			idx    int
			events sdk.Events
			msg    sdk.Msg
		)

		cacheCtx, writeCache := ctx.CacheContext()
		messages, err := proposal.GetMsgs()
		if err == nil {
			for idx, msg = range messages {
				handler := keeper.Router().Handler(msg)

				var res *sdk.Result
				res, err = handler(cacheCtx, msg)
				if err != nil {
					break
				}

				events = append(events, res.GetEvents()...)
			}
		}

		if err == nil {
			proposal.Status = govV1.StatusPassed
			tagValue = govtypes.AttributeValueProposalPassed
			logMsg = "passed"

			// write state to the underlying multi-store
			writeCache()

			// propagate the msg events to the current context
			ctx.EventManager().EmitEvents(events)
		} else {
			proposal.Status = govV1.StatusFailed
			tagValue = govtypes.AttributeValueProposalFailed
			logMsg = fmt.Sprintf("passed, but msg %d (%s) failed on execution: %s", idx, sdk.MsgTypeURL(msg), err)
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				govtypes.EventTypeActiveProposal,
				sdk.NewAttribute(govtypes.AttributeKeyProposalID, fmt.Sprintf("%d", proposal.Id)),
				sdk.NewAttribute(govtypes.AttributeKeyProposalResult, tagValue),
			),
		)

		logger.Info(
			"proposal tallied",
			"proposal", proposal.Id,
			"title", proposal.GetTitle(),
			"result", logMsg,
		)

		return false
	})
}
