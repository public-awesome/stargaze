package v18patch_test

import (
	"errors"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/suite"

	stargazeapp "github.com/public-awesome/stargaze/v18/app"
	"github.com/public-awesome/stargaze/v18/app/keepers"
	"github.com/public-awesome/stargaze/v18/app/upgrades"
	v18patch "github.com/public-awesome/stargaze/v18/app/upgrades/mainnet/v18_patch"
	"github.com/public-awesome/stargaze/v18/testutil/simapp"
	minttypes "github.com/public-awesome/stargaze/v18/x/mint/types"
)

const (
	// testForkHeight is the height we register the test fork at. The
	// production v18-patch fork's UpgradeHeight points to mainnet; the
	// dispatch tests swap in a small known height instead.
	testForkHeight = int64(5)

	// foreignChainID is any chain ID that is not v18patch.ChainID; used to
	// assert the ChainID guard skips the fork on the wrong network.
	foreignChainID = "foreign-1"
)

type V18PatchForkTestSuite struct {
	suite.Suite

	App *stargazeapp.App
}

func TestV18PatchForkTestSuite(t *testing.T) {
	suite.Run(t, new(V18PatchForkTestSuite))
}

func (s *V18PatchForkTestSuite) SetupTest() {
	s.App = simapp.New(s.T())
}

// breakMinDepositRatio reproduces the pre-fork mainnet state where
// MinDepositRatio was accidentally left empty. The SDK's
// gov v1 Params.ValidateBasic does NOT inspect MinDepositRatio, so the broken
// state passed every standard check at genesis and stayed live for a long
// time. The actual failure happens at deposit time in keeper.AddDeposit,
// which parses the field via sdkmath.LegacyNewDecFromStr.
func (s *V18PatchForkTestSuite) breakMinDepositRatio(ctx sdk.Context) {
	p, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)
	p.MinDepositRatio = ""
	s.Require().NoError(s.App.Keepers.GovKeeper.Params.Set(ctx, p))
}

// assertV18PatchParams reads gov params and asserts every field matches what
// the v18-patch fork is supposed to write.
func (s *V18PatchForkTestSuite) assertV18PatchParams(ctx sdk.Context) {
	p, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)

	s.Require().Equal(
		sdk.NewCoins(sdk.NewInt64Coin("ustars", 500_000_000_000)),
		sdk.Coins(p.MinDeposit),
	)
	s.Require().Equal(
		sdk.NewCoins(sdk.NewInt64Coin("ustars", 1_000_000_000_000)),
		sdk.Coins(p.ExpeditedMinDeposit),
	)

	s.Require().NotNil(p.MaxDepositPeriod)
	s.Require().Equal(336*time.Hour, *p.MaxDepositPeriod)
	s.Require().NotNil(p.VotingPeriod)
	s.Require().Equal(72*time.Hour, *p.VotingPeriod)
	s.Require().NotNil(p.ExpeditedVotingPeriod)
	s.Require().Equal(24*time.Hour, *p.ExpeditedVotingPeriod)

	s.Require().Equal(sdkmath.LegacyMustNewDecFromStr("0.2").String(), p.Quorum)
	s.Require().Equal(sdkmath.LegacyMustNewDecFromStr("0.5").String(), p.Threshold)
	s.Require().Equal(sdkmath.LegacyMustNewDecFromStr("0.334").String(), p.VetoThreshold)
	s.Require().Equal(sdkmath.LegacyMustNewDecFromStr("0.2").String(), p.MinInitialDepositRatio)
	s.Require().Equal(sdkmath.LegacyMustNewDecFromStr("0.5").String(), p.ProposalCancelRatio)
	s.Require().Equal(sdkmath.LegacyMustNewDecFromStr("0.667").String(), p.ExpeditedThreshold)
	s.Require().Equal(sdkmath.LegacyMustNewDecFromStr("0.01").String(), p.MinDepositRatio)

	s.Require().True(p.BurnVoteQuorum)
	s.Require().True(p.BurnProposalDepositPrevote)
	s.Require().True(p.BurnVoteVeto)

	// Same validation MsgUpdateParams runs. NOTE: ValidateBasic does not
	// touch MinDepositRatio, so this only proves the OTHER params are
	// well-formed. The MinDepositRatio repair is exercised end-to-end by
	// TestForkFixesProposalDepositPath via keeper.AddDeposit.
	s.Require().NoError(p.ValidateBasic())
}

// withTestForks replaces stargazeapp.Forks for the lifetime of the current
// test and restores it via t.Cleanup, so tests can exercise BeginBlockForks
// without needing the production placeholder UpgradeHeight to be set.
func (s *V18PatchForkTestSuite) withTestForks(forks []upgrades.Fork) {
	orig := stargazeapp.Forks
	stargazeapp.Forks = forks
	s.T().Cleanup(func() { stargazeapp.Forks = orig })
}

// TestRunForkLogic_UpdatesAllParams exercises the migration directly: it
// puts the gov params into the broken state observed on mainnet (empty
// MinDepositRatio), runs RunForkLogic, and verifies every post-fork value.
func (s *V18PatchForkTestSuite) TestRunForkLogic_UpdatesAllParams() {
	ctx := s.App.BaseApp.NewContext(false).WithChainID(v18patch.ChainID).WithBlockHeight(1)
	s.breakMinDepositRatio(ctx)

	s.Require().NoError(v18patch.RunForkLogic(ctx, s.App.Keepers))

	s.assertV18PatchParams(ctx)
}

// TestBeginBlockForks_AppliesV18Patch exercises the full dispatch path: it
// registers v18patch.Fork at testForkHeight, advances to that height on the
// v18-patch target chain, calls BeginBlockForks, and asserts the cache
// context's writes persisted.
func (s *V18PatchForkTestSuite) TestBeginBlockForks_AppliesV18Patch() {
	f := v18patch.Fork
	f.UpgradeHeight = testForkHeight
	s.withTestForks([]upgrades.Fork{f})

	ctx := s.App.BaseApp.NewContext(false).WithChainID(v18patch.ChainID).WithBlockHeight(testForkHeight)
	s.breakMinDepositRatio(ctx)

	stargazeapp.BeginBlockForks(ctx, s.App)

	s.assertV18PatchParams(ctx)
}

// TestBeginBlockForks_SkipsOnWrongChainID verifies the ChainID guard: the
// fork must not fire if the running chain isn't the one it targets, even at
// the matching height.
func (s *V18PatchForkTestSuite) TestBeginBlockForks_SkipsOnWrongChainID() {
	f := v18patch.Fork
	f.UpgradeHeight = testForkHeight
	s.withTestForks([]upgrades.Fork{f})

	ctx := s.App.BaseApp.NewContext(false).WithChainID(foreignChainID).WithBlockHeight(testForkHeight)
	s.breakMinDepositRatio(ctx)

	stargazeapp.BeginBlockForks(ctx, s.App)

	p, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)
	s.Require().Equal("", p.MinDepositRatio, "fork must not run on a chain it does not target")
}

// TestBeginBlockForks_SkipsAtWrongHeight verifies the height guard: even on
// the target chain, the fork only fires at the exact UpgradeHeight.
func (s *V18PatchForkTestSuite) TestBeginBlockForks_SkipsAtWrongHeight() {
	f := v18patch.Fork
	f.UpgradeHeight = testForkHeight
	s.withTestForks([]upgrades.Fork{f})

	ctx := s.App.BaseApp.NewContext(false).WithChainID(v18patch.ChainID).WithBlockHeight(testForkHeight + 1)
	s.breakMinDepositRatio(ctx)

	stargazeapp.BeginBlockForks(ctx, s.App)

	p, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)
	s.Require().Equal("", p.MinDepositRatio, "fork must not run at a non-matching block height")
}

// TestBeginBlockForks_DiscardsStateOnError verifies the cache-context wrapper:
// a fork that mutates state and then returns an error must not commit those
// writes, and the chain must keep running (no panic).
func (s *V18PatchForkTestSuite) TestBeginBlockForks_DiscardsStateOnError() {
	const sentinel = "0.500000000000000000" // distinguishable from the fork's mutation
	mutated := sdkmath.LegacyMustNewDecFromStr("0.999").String()

	failingFork := upgrades.Fork{
		UpgradeName:   "test-fail",
		ChainID:       v18patch.ChainID,
		UpgradeHeight: testForkHeight,
		BeginForkLogic: func(ctx sdk.Context, k keepers.StargazeKeepers) error {
			// Mutate state inside the cache context, then return an error.
			// The wrapper should discard the mutation.
			p, err := k.GovKeeper.Params.Get(ctx)
			if err != nil {
				return err
			}
			p.MinDepositRatio = mutated
			if err := k.GovKeeper.Params.Set(ctx, p); err != nil {
				return err
			}
			return errors.New("intentional failure for state-discard test")
		},
	}
	s.withTestForks([]upgrades.Fork{failingFork})

	ctx := s.App.BaseApp.NewContext(false).WithChainID(v18patch.ChainID).WithBlockHeight(testForkHeight)

	pre, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)
	pre.MinDepositRatio = sentinel
	s.Require().NoError(s.App.Keepers.GovKeeper.Params.Set(ctx, pre))

	// Failure is recovered into a logged error by ApplyFuncIfNoError; the
	// chain must not halt.
	s.Require().NotPanics(func() { stargazeapp.BeginBlockForks(ctx, s.App) })

	after, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)
	s.Require().Equal(sentinel, after.MinDepositRatio,
		"failed fork's writes must be discarded by the cache context")
}

// TestBeginBlockForks_WildcardChainIDFiresAnywhere verifies that a Fork with
// ChainID set to the wildcard (upgrades.ChainIDAny) fires on every chain at
// UpgradeHeight, including ones unrelated to the production target.
func (s *V18PatchForkTestSuite) TestBeginBlockForks_WildcardChainIDFiresAnywhere() {
	fired := 0
	wildcardFork := upgrades.Fork{
		UpgradeName:   "test-wildcard",
		ChainID:       upgrades.ChainIDAny,
		UpgradeHeight: testForkHeight,
		BeginForkLogic: func(_ sdk.Context, _ keepers.StargazeKeepers) error {
			fired++
			return nil
		},
	}
	s.withTestForks([]upgrades.Fork{wildcardFork})

	for _, chainID := range []string{v18patch.ChainID, "elgafar-1", "totally-unrelated-3", ""} {
		ctx := s.App.BaseApp.NewContext(false).WithChainID(chainID).WithBlockHeight(testForkHeight)
		stargazeapp.BeginBlockForks(ctx, s.App)
	}
	s.Require().Equal(4, fired, "ChainIDAny should fire on every chain ID, including unset")
}

// TestBeginBlockForks_TestnetForkOnlyFiresOnTestnet verifies the exact-match
// semantics for a non-mainnet target: a fork pinned to a testnet chain ID
// fires on testnet and skips on mainnet / unrelated chains.
func (s *V18PatchForkTestSuite) TestBeginBlockForks_TestnetForkOnlyFiresOnTestnet() {
	const testnetChainID = "elgafar-1"
	fired := 0
	testnetFork := upgrades.Fork{
		UpgradeName:   "test-testnet",
		ChainID:       testnetChainID,
		UpgradeHeight: testForkHeight,
		BeginForkLogic: func(_ sdk.Context, _ keepers.StargazeKeepers) error {
			fired++
			return nil
		},
	}
	s.withTestForks([]upgrades.Fork{testnetFork})

	// Matches target testnet
	ctx := s.App.BaseApp.NewContext(false).WithChainID(testnetChainID).WithBlockHeight(testForkHeight)
	stargazeapp.BeginBlockForks(ctx, s.App)
	s.Require().Equal(1, fired, "testnet fork should fire on the testnet chain ID")

	// Same height on mainnet must not fire
	ctx = s.App.BaseApp.NewContext(false).WithChainID(v18patch.ChainID).WithBlockHeight(testForkHeight)
	stargazeapp.BeginBlockForks(ctx, s.App)
	s.Require().Equal(1, fired, "testnet fork must not fire on mainnet")

	// Unrelated chain must not fire
	ctx = s.App.BaseApp.NewContext(false).WithChainID("randomchain-3").WithBlockHeight(testForkHeight)
	stargazeapp.BeginBlockForks(ctx, s.App)
	s.Require().Equal(1, fired, "testnet fork must not fire on an unrelated chain")
}

// TestForkFixesProposalDepositPath proves the user-visible bug is actually
// repaired. Before the fork, submitting a gov proposal with the full
// MinDeposit fails because keeper.AddDeposit can't parse an empty
// MinDepositRatio. After the fork rewrites the field, the same submission
// succeeds and reaches voting period.
func (s *V18PatchForkTestSuite) TestForkFixesProposalDepositPath() {
	ctx := s.App.BaseApp.NewContext(false).WithChainID(v18patch.ChainID).WithBlockHeight(1)

	// Reproduce the pre-fork mainnet state: ustars-denominated MinDeposit
	// (so the deposit denom matches what the fork writes) plus the empty
	// MinDepositRatio that broke deposits.
	p, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)
	p.MinDeposit = sdk.NewCoins(sdk.NewInt64Coin("ustars", 500_000_000_000))
	p.ExpeditedMinDeposit = sdk.NewCoins(sdk.NewInt64Coin("ustars", 1_000_000_000_000))
	p.MinDepositRatio = ""
	s.Require().NoError(s.App.Keepers.GovKeeper.Params.Set(ctx, p))

	// Fund a proposer with the full MinDeposit so the deposit amount is not
	// itself the reason a submission fails.
	proposer := sdk.AccAddress([]byte("v18patch-deposit-tst"))
	deposit := sdk.NewCoins(sdk.NewInt64Coin("ustars", 500_000_000_000))
	s.Require().NoError(s.App.Keepers.BankKeeper.MintCoins(ctx, minttypes.ModuleName, deposit))
	s.Require().NoError(s.App.Keepers.BankKeeper.SendCoinsFromModuleToAccount(
		ctx, minttypes.ModuleName, proposer, deposit,
	))

	msgServer := govkeeper.NewMsgServerImpl(&s.App.Keepers.GovKeeper)
	buildMsg := func() *govv1types.MsgSubmitProposal {
		// gov v1 requires either inner Messages or non-empty Metadata; we
		// use a non-empty metadata string. The SDK only enforces a
		// title/summary cross-check when the metadata parses as JSON, so a
		// plain string is fine.
		msg, err := govv1types.NewMsgSubmitProposal(
			nil,
			deposit,
			proposer.String(),
			"v18patch-deposit-test-metadata",
			"test title",
			"test summary",
			false,
		)
		s.Require().NoError(err)
		return msg
	}

	// Pre-fork: submission must fail. Run inside a cache context so the
	// partial proposal record created by Keeper.SubmitProposal (before
	// AddDeposit returns the parse error) doesn't leak into the next
	// attempt, matching BaseApp's per-msg rollback semantics.
	cacheCtx, _ := ctx.CacheContext()
	_, err = msgServer.SubmitProposal(cacheCtx, buildMsg())
	s.Require().Error(err,
		"submitting a proposal with full deposit must fail while MinDepositRatio is empty")
	// Confirm the failure is the LegacyDec parse on the empty ratio (so the
	// test isn't passing because of some unrelated validation error like
	// metadata/denom/amount).
	s.Require().ErrorContains(err, "decimal string cannot be empty",
		"pre-fork failure must originate from parsing the empty MinDepositRatio")

	// Apply the fork.
	s.Require().NoError(v18patch.RunForkLogic(ctx, s.App.Keepers))

	// Post-fork: the same submission succeeds and the proposal enters
	// voting period because the deposit meets MinDeposit.
	resp, err := msgServer.SubmitProposal(ctx, buildMsg())
	s.Require().NoError(err,
		"submitting a proposal with full deposit must succeed after the fork rewrites MinDepositRatio")
	s.Require().NotZero(resp.ProposalId)

	// Confirm the proposal actually advanced to voting period (full deposit
	// >= MinDeposit), not just that the call returned without error.
	proposal, err := s.App.Keepers.GovKeeper.Proposals.Get(ctx, resp.ProposalId)
	s.Require().NoError(err)
	s.Require().Equal(govv1types.StatusVotingPeriod, proposal.Status,
		"proposal must enter voting period when initial deposit covers MinDeposit")
}

// TestBeginBlockForks_ChainIDMatchIsExact verifies the matcher uses exact
// string equality (no prefix, suffix, or case-folding tolerance). Near-misses
// must not fire the fork.
func (s *V18PatchForkTestSuite) TestBeginBlockForks_ChainIDMatchIsExact() {
	fired := 0
	targetFork := upgrades.Fork{
		UpgradeName:   "test-exact",
		ChainID:       v18patch.ChainID, // "stargaze-1"
		UpgradeHeight: testForkHeight,
		BeginForkLogic: func(_ sdk.Context, _ keepers.StargazeKeepers) error {
			fired++
			return nil
		},
	}
	s.withTestForks([]upgrades.Fork{targetFork})

	nearMisses := []string{
		"stargaze",     // prefix of target
		"stargaze-2",   // different suffix
		"stargaze-10",  // numeric continuation
		"STARGAZE-1",   // uppercase
		"Stargaze-1",   // mixed case
		"stargaze-1 ",  // trailing whitespace
		" stargaze-1",  // leading whitespace
		"x-stargaze-1", // prefixed
		"stargaze-1-x", // suffixed
	}
	for _, chainID := range nearMisses {
		ctx := s.App.BaseApp.NewContext(false).WithChainID(chainID).WithBlockHeight(testForkHeight)
		stargazeapp.BeginBlockForks(ctx, s.App)
	}
	s.Require().Equal(0, fired, "near-miss chain IDs must not trigger the fork")

	// The exact match fires the fork
	ctx := s.App.BaseApp.NewContext(false).WithChainID(v18patch.ChainID).WithBlockHeight(testForkHeight)
	stargazeapp.BeginBlockForks(ctx, s.App)
	s.Require().Equal(1, fired, "exact chain ID match should fire the fork")
}
