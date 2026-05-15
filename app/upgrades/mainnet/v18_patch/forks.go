package v18patch

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v18/app/keepers"
)

// RunForkLogic executes the v18-patch hard fork. Governance is broken on the
// chain, so this fork rewrites the gov module params at a fixed block height
// instead of relying on a MsgUpdateParams proposal.
//
// All values mirror what `starsd q gov params --node https://rpc.stargaze-apis.com:443`
// returned at the time of the fork, so existing behavior is preserved
// verbatim except for MinDepositRatio. That field was left empty at genesis
// by accident — the SDK does not validate it as non-empty — and is set to 1%
// here so each individual deposit must clear MinDeposit * MinDepositRatio,
// rejecting dust-deposit spam.
//
// Errors are returned to the caller (BeginBlockForks), which runs this inside
// a cache context: returning an error discards any partial writes and the
// chain keeps producing blocks instead of halting.
func RunForkLogic(ctx sdk.Context, k keepers.StargazeKeepers) error {
	ctx.Logger().Info("applying v18-patch hard fork: rewriting gov params")

	params, err := k.GovKeeper.Params.Get(ctx)
	if err != nil {
		return err
	}

	maxDepositPeriod := 336 * time.Hour
	votingPeriod := 72 * time.Hour
	expeditedVotingPeriod := 24 * time.Hour

	// Total deposit required for a proposal to enter voting (500k stars).
	params.MinDeposit = sdk.NewCoins(sdk.NewInt64Coin("ustars", 500_000_000_000))
	// Window to accumulate MinDeposit before the proposal expires (14 days).
	params.MaxDepositPeriod = &maxDepositPeriod
	// Time voting is open once the proposal enters voting (3 days).
	params.VotingPeriod = &votingPeriod
	// Minimum fraction of staked supply that must vote for the result to count (20%).
	params.Quorum = sdkmath.LegacyMustNewDecFromStr("0.2").String()
	// Minimum YES share of non-abstain votes required to pass (50%).
	params.Threshold = sdkmath.LegacyMustNewDecFromStr("0.5").String()
	// NoWithVeto share that fails the proposal and burns the deposit (33.4%).
	params.VetoThreshold = sdkmath.LegacyMustNewDecFromStr("0.334").String()
	// Initial deposit must be at least this fraction of MinDeposit to register the proposal (20%).
	params.MinInitialDepositRatio = sdkmath.LegacyMustNewDecFromStr("0.2").String()
	// Fraction of the deposit burned when the proposer cancels the proposal (50%).
	params.ProposalCancelRatio = sdkmath.LegacyMustNewDecFromStr("0.5").String()
	// Voting window for expedited proposals (1 day).
	params.ExpeditedVotingPeriod = &expeditedVotingPeriod
	// YES share required to pass an expedited proposal (66.7%).
	params.ExpeditedThreshold = sdkmath.LegacyMustNewDecFromStr("0.667").String()
	// Total deposit required for an expedited proposal (1M stars).
	params.ExpeditedMinDeposit = sdk.NewCoins(sdk.NewInt64Coin("ustars", 1_000_000_000_000))
	// Burn the deposit when quorum is not reached.
	params.BurnVoteQuorum = true
	// Burn the deposit when a proposal never reaches MinDeposit and never enters voting.
	params.BurnProposalDepositPrevote = true
	// Burn the deposit when VetoThreshold is reached.
	params.BurnVoteVeto = true
	// Each deposit message must be at least this fraction of MinDeposit, blocking
	// dust spam (1% = 5k stars minimum per deposit). Previously empty; the SDK
	// does not validate this field as non-empty.
	params.MinDepositRatio = sdkmath.LegacyMustNewDecFromStr("0.01").String()

	// ValidateBasic catches malformed quorum / threshold / period / deposit
	// values before they hit the store. It does NOT cover MinDepositRatio —
	// the field this fork repairs — so that field is checked explicitly
	// after Set below.
	if err := params.ValidateBasic(); err != nil {
		return err
	}

	if err := k.GovKeeper.Params.Set(ctx, params); err != nil {
		return err
	}

	// Post-Set semantic check on MinDepositRatio. The runtime failure mode
	// for an empty / unparseable ratio lives in keeper.AddDeposit
	// (LegacyNewDecFromStr), well after Params.Set. Read back, parse, and
	// also confirm the ratio falls in (0, 1] — the fork's intent is to
	// ENABLE dust-spam protection, so 0 (disables it), negative, or > 1
	// values all defeat the purpose. Returning an error here causes the
	// cache context to drop the bad write and the chain keeps running on
	// the prior state.
	stored, err := k.GovKeeper.Params.Get(ctx)
	if err != nil {
		return err
	}
	ratio, err := sdkmath.LegacyNewDecFromStr(stored.MinDepositRatio)
	if err != nil {
		return fmt.Errorf("post-fork MinDepositRatio %q does not parse: %w", stored.MinDepositRatio, err)
	}
	if !ratio.IsPositive() || ratio.GT(sdkmath.LegacyOneDec()) {
		return fmt.Errorf("post-fork MinDepositRatio %s is out of (0, 1]", ratio)
	}
	return nil
}
