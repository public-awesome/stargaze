package v19_test

import (
	"errors"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	stargazeapp "github.com/public-awesome/stargaze/v18/app"
	"github.com/public-awesome/stargaze/v18/app/keepers"
	"github.com/public-awesome/stargaze/v18/app/upgrades"
	v19 "github.com/public-awesome/stargaze/v18/app/upgrades/mainnet/v19"
	"github.com/public-awesome/stargaze/v18/testutil/simapp"
)

const (
	// testForkHeight is the height we register the test fork at. The
	// production v19 fork uses a placeholder UpgradeHeight of 0, so the
	// dispatch tests swap in a non-zero height.
	testForkHeight = int64(5)

	// foreignChainID is any chain ID that is not v19.ChainID; used to assert
	// the ChainID guard skips the fork on the wrong network.
	foreignChainID = "foreign-1"
)

type V19ForkTestSuite struct {
	suite.Suite

	App *stargazeapp.App
}

func TestV19ForkTestSuite(t *testing.T) {
	suite.Run(t, new(V19ForkTestSuite))
}

func (s *V19ForkTestSuite) SetupTest() {
	s.App = simapp.New(s.T())
}

// breakMinDepositRatio reproduces the pre-fork mainnet state where
// MinDepositRatio was accidentally left empty. ValidateBasic rejects this
// state, but the keeper's collection setter does not, so the chain ran with
// it for a long time.
func (s *V19ForkTestSuite) breakMinDepositRatio(ctx sdk.Context) {
	p, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)
	p.MinDepositRatio = ""
	s.Require().NoError(s.App.Keepers.GovKeeper.Params.Set(ctx, p))
}

// assertV19Params reads gov params and asserts every field matches what the
// v19 fork is supposed to write.
func (s *V19ForkTestSuite) assertV19Params(ctx sdk.Context) {
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

	// Same validation MsgUpdateParams runs. Before the fork, this returned an
	// error because MinDepositRatio was empty.
	s.Require().NoError(p.ValidateBasic())
}

// withTestForks replaces stargazeapp.Forks for the lifetime of the current
// test and restores it via t.Cleanup, so tests can exercise BeginBlockForks
// without needing the production placeholder UpgradeHeight to be set.
func (s *V19ForkTestSuite) withTestForks(forks []upgrades.Fork) {
	orig := stargazeapp.Forks
	stargazeapp.Forks = forks
	s.T().Cleanup(func() { stargazeapp.Forks = orig })
}

// TestRunForkLogic_UpdatesAllParams exercises the migration directly: it
// puts the gov params into the broken state observed on mainnet (empty
// MinDepositRatio), runs RunForkLogic, and verifies every post-fork value.
func (s *V19ForkTestSuite) TestRunForkLogic_UpdatesAllParams() {
	ctx := s.App.BaseApp.NewContext(false).WithChainID(v19.ChainID).WithBlockHeight(1)
	s.breakMinDepositRatio(ctx)

	s.Require().NoError(v19.RunForkLogic(ctx, s.App.Keepers))

	s.assertV19Params(ctx)
}

// TestBeginBlockForks_AppliesV19 exercises the full dispatch path: it
// registers v19.Fork at testForkHeight, advances to that height on the v19
// target chain, calls BeginBlockForks, and asserts the cache context's writes
// persisted.
func (s *V19ForkTestSuite) TestBeginBlockForks_AppliesV19() {
	f := v19.Fork
	f.UpgradeHeight = testForkHeight
	s.withTestForks([]upgrades.Fork{f})

	ctx := s.App.BaseApp.NewContext(false).WithChainID(v19.ChainID).WithBlockHeight(testForkHeight)
	s.breakMinDepositRatio(ctx)

	stargazeapp.BeginBlockForks(ctx, s.App)

	s.assertV19Params(ctx)
}

// TestBeginBlockForks_SkipsOnWrongChainID verifies the ChainID guard: the
// fork must not fire if the running chain isn't the one it targets, even at
// the matching height.
func (s *V19ForkTestSuite) TestBeginBlockForks_SkipsOnWrongChainID() {
	f := v19.Fork
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
func (s *V19ForkTestSuite) TestBeginBlockForks_SkipsAtWrongHeight() {
	f := v19.Fork
	f.UpgradeHeight = testForkHeight
	s.withTestForks([]upgrades.Fork{f})

	ctx := s.App.BaseApp.NewContext(false).WithChainID(v19.ChainID).WithBlockHeight(testForkHeight + 1)
	s.breakMinDepositRatio(ctx)

	stargazeapp.BeginBlockForks(ctx, s.App)

	p, err := s.App.Keepers.GovKeeper.Params.Get(ctx)
	s.Require().NoError(err)
	s.Require().Equal("", p.MinDepositRatio, "fork must not run at a non-matching block height")
}

// TestBeginBlockForks_DiscardsStateOnError verifies the cache-context wrapper:
// a fork that mutates state and then returns an error must not commit those
// writes, and the chain must keep running (no panic).
func (s *V19ForkTestSuite) TestBeginBlockForks_DiscardsStateOnError() {
	const sentinel = "0.500000000000000000" // distinguishable from the fork's mutation
	mutated := sdkmath.LegacyMustNewDecFromStr("0.999").String()

	failingFork := upgrades.Fork{
		UpgradeName:   "test-fail",
		ChainID:       v19.ChainID,
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

	ctx := s.App.BaseApp.NewContext(false).WithChainID(v19.ChainID).WithBlockHeight(testForkHeight)

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
func (s *V19ForkTestSuite) TestBeginBlockForks_WildcardChainIDFiresAnywhere() {
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

	for _, chainID := range []string{v19.ChainID, "elgafar-1", "totally-unrelated-3", ""} {
		ctx := s.App.BaseApp.NewContext(false).WithChainID(chainID).WithBlockHeight(testForkHeight)
		stargazeapp.BeginBlockForks(ctx, s.App)
	}
	s.Require().Equal(4, fired, "ChainIDAny should fire on every chain ID, including unset")
}

// TestBeginBlockForks_TestnetForkOnlyFiresOnTestnet verifies the exact-match
// semantics for a non-mainnet target: a fork pinned to a testnet chain ID
// fires on testnet and skips on mainnet / unrelated chains.
func (s *V19ForkTestSuite) TestBeginBlockForks_TestnetForkOnlyFiresOnTestnet() {
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
	ctx = s.App.BaseApp.NewContext(false).WithChainID(v19.ChainID).WithBlockHeight(testForkHeight)
	stargazeapp.BeginBlockForks(ctx, s.App)
	s.Require().Equal(1, fired, "testnet fork must not fire on mainnet")

	// Unrelated chain must not fire
	ctx = s.App.BaseApp.NewContext(false).WithChainID("randomchain-3").WithBlockHeight(testForkHeight)
	stargazeapp.BeginBlockForks(ctx, s.App)
	s.Require().Equal(1, fired, "testnet fork must not fire on an unrelated chain")
}

// TestBeginBlockForks_ChainIDMatchIsExact verifies the matcher uses exact
// string equality (no prefix, suffix, or case-folding tolerance). Near-misses
// must not fire the fork.
func (s *V19ForkTestSuite) TestBeginBlockForks_ChainIDMatchIsExact() {
	fired := 0
	targetFork := upgrades.Fork{
		UpgradeName:   "test-exact",
		ChainID:       v19.ChainID, // "stargaze-1"
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
	ctx := s.App.BaseApp.NewContext(false).WithChainID(v19.ChainID).WithBlockHeight(testForkHeight)
	stargazeapp.BeginBlockForks(ctx, s.App)
	s.Require().Equal(1, fired, "exact chain ID match should fire the fork")
}
