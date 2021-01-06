package curating

import (
	"fmt"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/keeper"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// BeginBlocker to fund reward pool on every begin block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	if err := k.InflateRewardPool(ctx); err != nil {
		panic(fmt.Sprintf("Error funding reward pool: %s", err.Error()))
	}

}

// EndBlocker called every block, update validator set
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	endTimes := make(map[time.Time]bool)
	k.IterateExpiredPosts(ctx, func(post types.Post) bool {
		postIDStr := post.String()
		k.Logger(ctx).Info(
			fmt.Sprintf("Processing vendor %d post %v", post.VendorID, post.PostID))

		qv := NewQVFData(ctx, k)

		// iterate upvoters, returning deposits, and tallying upvotes
		k.IterateUpvotes(ctx, post.VendorID, post.PostID,
			func(upvote types.Upvote) (stop bool) {
				var err error
				qv, err = qv.TallyVote(upvote.VoteAmount.Amount)
				if err != nil {
					panic(err)
				}

				return false
			})

		rewardAccount, err := sdk.AccAddressFromBech32(post.RewardAccount)
		if err != nil {
			panic(err)
		}
		creatorVotinPoolReward, err := k.RewardCreatorFromVotingPool(ctx, rewardAccount, qv.VotingPool)
		if err != nil {
			panic(err)
		}
		emitRewardEvent(ctx, types.EventTypeProtocolReward, types.EventTypeVotingPoolReturn,
			post.RewardAccount, postIDStr, creatorVotinPoolReward.String())
		creatorProtocolReward, err := k.RewardCreatorFromProtocol(ctx, rewardAccount, qv.MatchPool())
		if err != nil {
			panic(err)
		}
		emitRewardEvent(ctx, types.EventTypeProtocolReward, types.AttributeRewardTypeCreator,
			post.RewardAccount, postIDStr, creatorProtocolReward.String())

		curatorAlloc := sdk.OneDec().Sub(k.GetParams(ctx).CreatorVotingRewardAllocation)
		curatorVotingReward := curatorAlloc.MulInt(qv.VoterReward()).TruncateInt()

		curatorMatchPerVote := qv.MatchPoolPerVote()

		k.IterateUpvotes(ctx, post.VendorID, post.PostID,
			func(upvote types.Upvote) (stop bool) {
				rewardAccount, err := sdk.AccAddressFromBech32(upvote.RewardAccount)
				if err != nil {
					panic(err)
				}
				// distribute quadratic voting per capita reward from voting pool
				votingPoolReward, err := k.SendVotingReward(ctx, rewardAccount, curatorVotingReward)
				if err != nil {
					panic(err)
				}
				emitRewardEvent(ctx, types.EventTypeProtocolReward, types.EventTypeVotingPoolReturn,
					upvote.RewardAccount, postIDStr, votingPoolReward.String())

				curatingProtocolReward, err := sendMatchingReward(ctx, k, upvote.VoteAmount.Amount,
					curatorMatchPerVote, rewardAccount)
				if err != nil {
					panic(err)
				}
				emitRewardEvent(ctx, types.EventTypeProtocolReward, types.EventTypeProtocolReward,
					upvote.RewardAccount, postIDStr, curatingProtocolReward.String())

				// Remove upvote
				err = k.DeleteUpvote(ctx, post.VendorID, post.PostID, upvote)
				if err != nil {
					panic(err)
				}

				return false
			})

		endTimes[post.GetCuratingEndTime()] = true
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeCurationComplete,
				sdk.NewAttribute(types.AttributeKeyPostID, postIDStr),
			),
		})

		return false
	})

	// remove processed curationEndtime from queue
	for t := range endTimes {
		k.RemoveFromCurationQueue(ctx, t)
	}
	return []abci.ValidatorUpdate{}
}

// sends matching reward in proportion to upvote
func sendMatchingReward(ctx sdk.Context, k keeper.Keeper, upvoteAmount sdk.Int,
	matchPoolPerVote sdk.Dec, rewardAccount sdk.AccAddress) (sdk.Coin, error) {

	voteNum, err := upvoteAmount.QuoRaw(1_000_000).ToDec().ApproxSqrt()
	if err != nil {
		return sdk.Coin{}, err
	}
	matchReward := matchPoolPerVote.Mul(voteNum)

	// distribute quadratic funding reward from protocol reward pool
	reward, err := k.SendMatchingReward(ctx, rewardAccount, matchReward)
	if err != nil {
		return sdk.Coin{}, err
	}

	return reward, nil
}

func emitRewardEvent(ctx sdk.Context, evtType, evtSubType, address, postID, amount string) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			evtType,
			sdk.NewAttribute(types.AttributeKeyProtocolRewardType, evtSubType),
			sdk.NewAttribute(types.AttributeKeyRewardAccount, address),
			sdk.NewAttribute(types.AttributeKeyPostID, postID),
		),
	})
}
