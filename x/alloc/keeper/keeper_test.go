package keeper_test

import (
	"testing"
	"time"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/public-awesome/stargaze/v13/app"
	"github.com/public-awesome/stargaze/v13/testutil/simapp"
	"github.com/public-awesome/stargaze/v13/x/alloc/keeper"
	"github.com/public-awesome/stargaze/v13/x/alloc/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx sdk.Context
	app *app.App
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = simapp.New(suite.T())
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stargaze-1", Time: time.Now().UTC()})
	suite.app.AllocKeeper.SetParams(suite.ctx, types.DefaultParams())
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

	allocKeeper := suite.app.AllocKeeper

	params := suite.app.AllocKeeper.GetParams(suite.ctx)

	params.DistributionProportions.NftIncentives = sdk.ZeroDec()

	suite.app.AllocKeeper.SetParams(suite.ctx, params)

	err := allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestModuleAccountAddress() {
	acc, err := sdk.GetFromBech32("stars1mnyrspq208uv5m2krdctan2dkyht0szje9s43h", "stars")
	suite.Require().NoError(err)
	suite.Require().Equal(authtypes.NewModuleAddress(types.SupplementPoolName).Bytes(), acc)
}

func (suite *KeeperTestSuite) TestDistribution() {
	suite.SetupTest()

	denom := suite.app.StakingKeeper.BondDenom(suite.ctx)
	allocKeeper := suite.app.AllocKeeper
	params := suite.app.AllocKeeper.GetParams(suite.ctx)
	devRewardsReceiver := sdk.AccAddress([]byte("addr1---------------"))
	nftIncentives := sdk.AccAddress([]byte("addr2---------------"))
	params.SupplementAmount = sdk.NewCoins(sdk.NewInt64Coin(denom, 10_000_000))
	params.DistributionProportions.NftIncentives = sdk.NewDecWithPrec(45, 2)
	params.DistributionProportions.DeveloperRewards = sdk.NewDecWithPrec(15, 2)
	params.WeightedDeveloperRewardsReceivers = []types.WeightedAddress{
		{
			Address: devRewardsReceiver.String(),
			Weight:  sdk.NewDec(1),
		},
	}
	params.WeightedIncentivesRewardsReceivers = []types.WeightedAddress{
		{
			Address: nftIncentives.String(),
			Weight:  sdk.NewDec(1),
		},
	}
	suite.app.AllocKeeper.SetParams(suite.ctx, params)

	feePool := suite.app.DistrKeeper.GetFeePool(suite.ctx)
	feeCollector := suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	suite.Equal(
		"0",
		suite.app.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())
	suite.Equal(
		sdk.NewDec(0),
		feePool.CommunityPool.AmountOf(denom))

	mintCoin := sdk.NewCoin(denom, sdk.NewInt(100_000))
	mintCoins := sdk.Coins{mintCoin}
	feeCollectorAccount := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, authtypes.FeeCollectorName)
	suite.Require().NotNil(feeCollectorAccount)

	suite.Require().NoError(FundModuleAccount(suite.app.BankKeeper, suite.ctx, feeCollectorAccount.GetName(), mintCoins))

	feeCollector = suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	suite.Equal(
		mintCoin.Amount.String(),
		suite.app.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())

	suite.Equal(
		sdk.NewDec(0),
		feePool.CommunityPool.AmountOf(denom))

	err := allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)

	feeCollector = suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	modulePortion := params.DistributionProportions.NftIncentives.
		Add(params.DistributionProportions.DeveloperRewards).
		Add(params.DistributionProportions.CommunityPool)

	// remaining going to next module should be 100% - 60%  - 5% community pooll = 35%
	suite.Equal(
		sdk.NewDecFromInt(mintCoin.Amount).Mul(sdk.NewDecWithPrec(100, 2).Sub(modulePortion)).RoundInt().String(),
		suite.app.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())

	// assigned dev reward receiver should get the allocation
	suite.Equal(
		sdk.NewDecFromInt(mintCoin.Amount).Mul(params.DistributionProportions.DeveloperRewards).TruncateInt(),
		suite.app.BankKeeper.GetBalance(suite.ctx, devRewardsReceiver, denom).Amount)

	// assigned incentive address should receive the allocation
	suite.Equal(
		sdk.NewDecFromInt(mintCoin.Amount).Mul(params.DistributionProportions.NftIncentives).TruncateInt(),
		suite.app.BankKeeper.GetBalance(suite.ctx, nftIncentives, denom).Amount)

	// community pool should get 5%
	feePool = suite.app.DistrKeeper.GetFeePool(suite.ctx)
	suite.Equal(
		sdk.NewDecFromInt(sdk.NewInt(5_000)).String(),
		feePool.CommunityPool.AmountOf(denom).String(),
	)
}

func (suite *KeeperTestSuite) TestFairburnPool() {
	suite.SetupTest()

	pub1 := secp256k1.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pub1.Address())

	// set params
	denom := suite.app.StakingKeeper.BondDenom(suite.ctx)
	allocKeeper := suite.app.AllocKeeper
	params := suite.app.AllocKeeper.GetParams(suite.ctx)
	devRewardsReceiver := sdk.AccAddress([]byte("addr1---------------"))
	params.DistributionProportions.NftIncentives = sdk.NewDecWithPrec(45, 2)
	params.DistributionProportions.DeveloperRewards = sdk.NewDecWithPrec(15, 2)
	params.WeightedDeveloperRewardsReceivers = []types.WeightedAddress{
		{
			Address: devRewardsReceiver.String(),
			Weight:  sdk.NewDec(1),
		},
	}
	allocKeeper.SetParams(suite.ctx, params)
	fundAmount := sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(100_000_000)))

	fairBurnPool := suite.app.AccountKeeper.GetModuleAddress(types.FairburnPoolName)
	feeCollector := suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)

	// should be 0
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).IsZero())
	err := allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)

	// should be 0
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).IsZero())

	msgServer := keeper.NewMsgServerImpl(allocKeeper)

	// fundAccount
	err = FundAccount(suite.app.BankKeeper, suite.ctx, addr1, fundAmount)
	suite.NoError(err)
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
	_, err = msgServer.FundFairburnPool(sdk.WrapSDKContext(suite.ctx), types.NewMsgFundFairburnPool(addr1, fundAmount))
	suite.NoError(err)

	// should have funds now
	suite.Require().Equal(fundAmount.String(), suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).String())
	// still 0
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).IsZero())

	err = allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)

	// fee collector should have funds now
	suite.Require().Equal(fundAmount.String(), suite.app.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).String())
	// fairburn pool should be 0
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
}

func (suite *KeeperTestSuite) TestDistributionWithSupplement() {
	suite.SetupTest()

	denom := suite.app.StakingKeeper.BondDenom(suite.ctx)
	allocKeeper := suite.app.AllocKeeper
	params := suite.app.AllocKeeper.GetParams(suite.ctx)
	devRewardsReceiver := sdk.AccAddress([]byte("addr1---------------"))
	nftIncentives := sdk.AccAddress([]byte("addr2---------------"))

	supplementAmount := sdk.NewInt64Coin(denom, 10_000)
	params.SupplementAmount = sdk.NewCoins(supplementAmount)
	params.DistributionProportions.NftIncentives = sdk.NewDecWithPrec(20, 2)
	params.DistributionProportions.DeveloperRewards = sdk.NewDecWithPrec(15, 2)
	params.WeightedDeveloperRewardsReceivers = []types.WeightedAddress{
		{
			Address: devRewardsReceiver.String(),
			Weight:  sdk.NewDec(1),
		},
	}
	params.WeightedIncentivesRewardsReceivers = []types.WeightedAddress{
		{
			Address: nftIncentives.String(),
			Weight:  sdk.NewDec(1),
		},
	}
	suite.app.AllocKeeper.SetParams(suite.ctx, params)
	suite.Require().False(suite.app.AllocKeeper.GetParams(suite.ctx).SupplementAmount.IsZero())

	feePool := suite.app.DistrKeeper.GetFeePool(suite.ctx)
	feeCollector := suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	suite.Equal(
		"0",
		suite.app.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())
	suite.Equal(
		sdk.NewDec(0),
		feePool.CommunityPool.AmountOf(denom))

	mintCoin := sdk.NewCoin(denom, sdk.NewInt(100_000))
	mintCoins := sdk.Coins{mintCoin}
	feeCollectorAccount := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, authtypes.FeeCollectorName)
	suite.Require().NotNil(feeCollectorAccount)

	supplementAccount := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, types.SupplementPoolName)
	suite.Require().NotNil(supplementAccount)

	suite.Require().NoError(FundModuleAccount(suite.app.BankKeeper, suite.ctx, feeCollectorAccount.GetName(), mintCoins))
	suite.Require().NoError(FundModuleAccount(suite.app.BankKeeper, suite.ctx, supplementAccount.GetName(), mintCoins))

	feeCollector = suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	suite.Equal(
		mintCoin.Amount.String(),
		suite.app.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())
	supplementAddress := supplementAccount.GetAddress()
	suite.Equal(
		mintCoin.Amount.String(),
		suite.app.BankKeeper.GetAllBalances(suite.ctx, supplementAddress).AmountOf(denom).String())

	suite.Equal(
		sdk.NewDec(0),
		feePool.CommunityPool.AmountOf(denom))

	err := allocKeeper.DistributeInflation(suite.ctx)
	suite.Require().NoError(err)

	feeCollector = suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	modulePortion := params.DistributionProportions.NftIncentives.
		Add(params.DistributionProportions.DeveloperRewards).
		Add(params.DistributionProportions.CommunityPool)
	suite.Equal(sdk.NewDecWithPrec(40, 2), modulePortion)

	totalAmount := mintCoin.Add(supplementAmount)
	// remaining going to next module should be (100_000 + 10_000)  - 40% = 66
	suite.Equal(
		sdk.NewDecFromInt(totalAmount.Amount).Mul(sdk.NewDecWithPrec(100, 2).Sub(modulePortion)).TruncateInt().String(),
		suite.app.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String(),
	)

	// assigned dev reward receiver should get the allocation
	suite.Equal(
		sdk.NewDecFromInt(totalAmount.Amount).Mul(params.DistributionProportions.DeveloperRewards).TruncateInt(),
		suite.app.BankKeeper.GetBalance(suite.ctx, devRewardsReceiver, denom).Amount)

	// assigned incentive address should receive the allocation
	suite.Equal(
		sdk.NewDecFromInt(totalAmount.Amount).Mul(params.DistributionProportions.NftIncentives).TruncateInt(),
		suite.app.BankKeeper.GetBalance(suite.ctx, nftIncentives, denom).Amount)

	// community pool should get 5%
	feePool = suite.app.DistrKeeper.GetFeePool(suite.ctx)
	suite.Equal(
		sdk.NewDecFromInt(totalAmount.Amount).Mul(params.DistributionProportions.CommunityPool).TruncateInt().String(),
		feePool.CommunityPool.AmountOf(denom).TruncateInt().String(),
	)

	// should have been decresead by supplement amount
	suite.Equal(
		"90000",
		suite.app.BankKeeper.GetAllBalances(suite.ctx, supplementAddress).AmountOf(denom).String())
}
