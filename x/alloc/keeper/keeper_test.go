package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/core/header"
	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/public-awesome/stargaze/v16/app"
	"github.com/public-awesome/stargaze/v16/testutil/simapp"
	"github.com/public-awesome/stargaze/v16/x/alloc/keeper"
	"github.com/public-awesome/stargaze/v16/x/alloc/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx sdk.Context
	app *app.App
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = simapp.New(suite.T())
	suite.ctx = suite.app.BaseApp.NewContext(false).WithHeaderInfo(header.Info{
		Height:  1,
		Time:    time.Now().UTC(),
		ChainID: "stargaze-1",
	})
	err := suite.app.Keepers.AllocKeeper.SetParams(suite.ctx, types.DefaultParams())
	suite.Require().NoError(err)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

func FundModuleAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, recipientMod, amounts)
}

func (suite *KeeperTestSuite) TestZeroAllocation() {
	suite.SetupTest()

	allocKeeper := suite.app.Keepers.AllocKeeper

	params, err := suite.app.Keepers.AllocKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)

	params.DistributionProportions.NftIncentives = math.LegacyZeroDec()

	err = suite.app.Keepers.AllocKeeper.SetParams(suite.ctx, params)
	suite.Require().NoError(err)

	err = allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestModuleAccountAddress() {
	acc, err := sdk.GetFromBech32("stars1mnyrspq208uv5m2krdctan2dkyht0szje9s43h", "stars")
	suite.Require().NoError(err)
	suite.Require().Equal(authtypes.NewModuleAddress(types.SupplementPoolName).Bytes(), acc)
}

func (suite *KeeperTestSuite) TestDistribution() {
	suite.SetupTest()

	denom, err := suite.app.Keepers.StakingKeeper.BondDenom(suite.ctx)
	suite.Require().NoError(err)
	allocKeeper := suite.app.Keepers.AllocKeeper
	params, err := suite.app.Keepers.AllocKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)
	devRewardsReceiver := sdk.AccAddress([]byte("addr1---------------"))
	nftIncentives := sdk.AccAddress([]byte("addr2---------------"))
	params.SupplementAmount = sdk.NewCoins(sdk.NewInt64Coin(denom, 10_000_000))
	params.DistributionProportions.NftIncentives = math.LegacyNewDecWithPrec(45, 2)
	params.DistributionProportions.DeveloperRewards = math.LegacyNewDecWithPrec(15, 2)
	params.WeightedDeveloperRewardsReceivers = []types.WeightedAddress{
		{
			Address: devRewardsReceiver.String(),
			Weight:  math.LegacyNewDec(1),
		},
	}
	params.WeightedIncentivesRewardsReceivers = []types.WeightedAddress{
		{
			Address: nftIncentives.String(),
			Weight:  math.LegacyNewDec(1),
		},
	}

	err = suite.app.Keepers.AllocKeeper.SetParams(suite.ctx, params)
	suite.Require().NoError(err)

	feePool, err := suite.app.Keepers.DistrKeeper.FeePool.Get(suite.ctx)
	suite.Require().NoError(err)
	feeCollector := suite.app.Keepers.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	suite.Equal(
		"0",
		suite.app.Keepers.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())
	suite.Equal(
		math.LegacyNewDec(0),
		feePool.CommunityPool.AmountOf(denom))

	mintCoin := sdk.NewCoin(denom, math.NewInt(100_000))
	mintCoins := sdk.Coins{mintCoin}
	feeCollectorAccount := suite.app.Keepers.AccountKeeper.GetModuleAccount(suite.ctx, authtypes.FeeCollectorName)
	suite.Require().NotNil(feeCollectorAccount)

	suite.Require().NoError(FundModuleAccount(suite.app.Keepers.BankKeeper, suite.ctx, feeCollectorAccount.GetName(), mintCoins))

	feeCollector = suite.app.Keepers.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	suite.Equal(
		mintCoin.Amount.String(),
		suite.app.Keepers.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())

	suite.Equal(
		math.LegacyNewDec(0),
		feePool.CommunityPool.AmountOf(denom))

	err = allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)

	feeCollector = suite.app.Keepers.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	modulePortion := params.DistributionProportions.NftIncentives.
		Add(params.DistributionProportions.DeveloperRewards).
		Add(params.DistributionProportions.CommunityPool)

	// remaining going to next module should be 100% - 60%  - 5% community pooll = 35%
	suite.Equal(
		math.LegacyNewDecFromInt(mintCoin.Amount).Mul(math.LegacyNewDecWithPrec(100, 2).Sub(modulePortion)).RoundInt().String(),
		suite.app.Keepers.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())

	// assigned dev reward receiver should get the allocation
	suite.Equal(
		math.LegacyNewDecFromInt(mintCoin.Amount).Mul(params.DistributionProportions.DeveloperRewards).TruncateInt(),
		suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, devRewardsReceiver, denom).Amount)

	// assigned incentive address should receive the allocation
	suite.Equal(
		math.LegacyNewDecFromInt(mintCoin.Amount).Mul(params.DistributionProportions.NftIncentives).TruncateInt(),
		suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, nftIncentives, denom).Amount)

	// community pool should get 5%
	feePool, err = suite.app.Keepers.DistrKeeper.FeePool.Get(suite.ctx)
	suite.Require().NoError(err)
	suite.Equal(
		math.LegacyNewDecFromInt(math.NewInt(5_000)).String(),
		feePool.CommunityPool.AmountOf(denom).String(),
	)
}

func (suite *KeeperTestSuite) TestFairburnPool() {
	suite.SetupTest()

	pub1 := secp256k1.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pub1.Address())

	// set params
	denom, err := suite.app.Keepers.StakingKeeper.BondDenom(suite.ctx)
	suite.Require().NoError(err)
	allocKeeper := suite.app.Keepers.AllocKeeper
	params, err := suite.app.Keepers.AllocKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)
	devRewardsReceiver := sdk.AccAddress([]byte("addr1---------------"))
	params.DistributionProportions.NftIncentives = math.LegacyNewDecWithPrec(45, 2)
	params.DistributionProportions.DeveloperRewards = math.LegacyNewDecWithPrec(15, 2)
	params.WeightedDeveloperRewardsReceivers = []types.WeightedAddress{
		{
			Address: devRewardsReceiver.String(),
			Weight:  math.LegacyNewDec(1),
		},
	}
	err = allocKeeper.SetParams(suite.ctx, params)
	suite.Require().NoError(err)
	fundAmount := sdk.NewCoins(sdk.NewCoin(denom, math.NewInt(100_000_000)))

	fairBurnPool := suite.app.Keepers.AccountKeeper.GetModuleAddress(types.FairburnPoolName)
	feeCollector := suite.app.Keepers.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)

	// should be 0
	suite.Require().True(suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
	suite.Require().True(suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).IsZero())
	err = allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)

	// should be 0
	suite.Require().True(suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
	suite.Require().True(suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).IsZero())

	msgServer := keeper.NewMsgServerImpl(allocKeeper)

	// fundAccount
	err = FundAccount(suite.app.Keepers.BankKeeper, suite.ctx, addr1, fundAmount)
	suite.NoError(err)
	suite.Require().True(suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
	_, err = msgServer.FundFairburnPool(suite.ctx, types.NewMsgFundFairburnPool(addr1, fundAmount))
	suite.NoError(err)

	// should have funds now
	suite.Require().Equal(fundAmount.String(), suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).String())
	// still 0
	suite.Require().True(suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).IsZero())

	err = allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)

	// fee collector should have funds now
	suite.Require().Equal(fundAmount.String(), suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).String())
	// fairburn pool should be 0
	suite.Require().True(suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
}

func (suite *KeeperTestSuite) TestDistributionWithSupplement() {
	suite.SetupTest()

	denom, err := suite.app.Keepers.StakingKeeper.BondDenom(suite.ctx)
	suite.Require().NoError(err)
	allocKeeper := suite.app.Keepers.AllocKeeper
	params, err := suite.app.Keepers.AllocKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)
	devRewardsReceiver := sdk.AccAddress([]byte("addr1---------------"))
	nftIncentives := sdk.AccAddress([]byte("addr2---------------"))

	supplementAmount := sdk.NewInt64Coin(denom, 10_000)
	params.SupplementAmount = sdk.NewCoins(supplementAmount)
	params.DistributionProportions.NftIncentives = math.LegacyNewDecWithPrec(20, 2)
	params.DistributionProportions.DeveloperRewards = math.LegacyNewDecWithPrec(15, 2)
	params.WeightedDeveloperRewardsReceivers = []types.WeightedAddress{
		{
			Address: devRewardsReceiver.String(),
			Weight:  math.LegacyNewDec(1),
		},
	}
	params.WeightedIncentivesRewardsReceivers = []types.WeightedAddress{
		{
			Address: nftIncentives.String(),
			Weight:  math.LegacyNewDec(1),
		},
	}
	err = suite.app.Keepers.AllocKeeper.SetParams(suite.ctx, params)
	suite.Require().NoError(err)
	params, err = suite.app.Keepers.AllocKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)
	suite.Require().False(params.SupplementAmount.IsZero())

	feePool, err := suite.app.Keepers.DistrKeeper.FeePool.Get(suite.ctx)
	suite.Require().NoError(err)
	feeCollector := suite.app.Keepers.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	suite.Equal(
		"0",
		suite.app.Keepers.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())
	suite.Equal(
		math.LegacyNewDec(0),
		feePool.CommunityPool.AmountOf(denom))

	mintCoin := sdk.NewCoin(denom, math.NewInt(100_000))
	mintCoins := sdk.Coins{mintCoin}
	feeCollectorAccount := suite.app.Keepers.AccountKeeper.GetModuleAccount(suite.ctx, authtypes.FeeCollectorName)
	suite.Require().NotNil(feeCollectorAccount)

	supplementAccount := suite.app.Keepers.AccountKeeper.GetModuleAccount(suite.ctx, types.SupplementPoolName)
	suite.Require().NotNil(supplementAccount)

	suite.Require().NoError(FundModuleAccount(suite.app.Keepers.BankKeeper, suite.ctx, feeCollectorAccount.GetName(), mintCoins))
	suite.Require().NoError(FundModuleAccount(suite.app.Keepers.BankKeeper, suite.ctx, supplementAccount.GetName(), mintCoins))

	feeCollector = suite.app.Keepers.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	suite.Equal(
		mintCoin.Amount.String(),
		suite.app.Keepers.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())
	supplementAddress := supplementAccount.GetAddress()
	suite.Equal(
		mintCoin.Amount.String(),
		suite.app.Keepers.BankKeeper.GetAllBalances(suite.ctx, supplementAddress).AmountOf(denom).String())

	suite.Equal(
		math.LegacyNewDec(0),
		feePool.CommunityPool.AmountOf(denom))

	err = allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)

	feeCollector = suite.app.Keepers.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	modulePortion := params.DistributionProportions.NftIncentives.
		Add(params.DistributionProportions.DeveloperRewards).
		Add(params.DistributionProportions.CommunityPool)
	suite.Equal(math.LegacyNewDecWithPrec(40, 2), modulePortion)

	totalAmount := mintCoin.Add(supplementAmount)
	// remaining going to next module should be (100_000 + 10_000)  - 40% = 66
	suite.Equal(
		math.LegacyNewDecFromInt(totalAmount.Amount).Mul(math.LegacyNewDecWithPrec(100, 2).Sub(modulePortion)).TruncateInt().String(),
		suite.app.Keepers.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String(),
	)

	// assigned dev reward receiver should get the allocation
	suite.Equal(
		math.LegacyNewDecFromInt(totalAmount.Amount).Mul(params.DistributionProportions.DeveloperRewards).TruncateInt(),
		suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, devRewardsReceiver, denom).Amount)

	// assigned incentive address should receive the allocation
	suite.Equal(
		math.LegacyNewDecFromInt(totalAmount.Amount).Mul(params.DistributionProportions.NftIncentives).TruncateInt(),
		suite.app.Keepers.BankKeeper.GetBalance(suite.ctx, nftIncentives, denom).Amount)

	// community pool should get 5%
	feePool, err = suite.app.Keepers.DistrKeeper.FeePool.Get(suite.ctx)
	suite.Require().NoError(err)
	suite.Equal(
		math.LegacyNewDecFromInt(totalAmount.Amount).Mul(params.DistributionProportions.CommunityPool).TruncateInt().String(),
		feePool.CommunityPool.AmountOf(denom).TruncateInt().String(),
	)

	// should have been decresead by supplement amount
	suite.Equal(
		"90000",
		suite.app.Keepers.BankKeeper.GetAllBalances(suite.ctx, supplementAddress).AmountOf(denom).String())
}
