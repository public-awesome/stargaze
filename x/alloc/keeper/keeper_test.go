package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/public-awesome/stargaze/v6/app"
	"github.com/public-awesome/stargaze/v6/testutil/simapp"
	"github.com/public-awesome/stargaze/v6/x/alloc/keeper"
	"github.com/public-awesome/stargaze/v6/x/alloc/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx sdk.Context
	app *app.App
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = simapp.New(suite.T().TempDir())
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

func (suite *KeeperTestSuite) TestDistribution() {
	suite.SetupTest()

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

	allocKeeper.DistributeInflation(suite.ctx)

	feeCollector = suite.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	modulePortion := params.DistributionProportions.NftIncentives.
		Add(params.DistributionProportions.DeveloperRewards) // 60%

	// remaining going to next module should be 100% - 60% = 40%
	suite.Equal(
		mintCoin.Amount.ToDec().Mul(sdk.NewDecWithPrec(100, 2).Sub(modulePortion)).RoundInt().String(),
		suite.app.BankKeeper.GetAllBalances(suite.ctx, feeCollector).AmountOf(denom).String())

	suite.Equal(
		mintCoin.Amount.ToDec().Mul(params.DistributionProportions.DeveloperRewards).TruncateInt(),
		suite.app.BankKeeper.GetBalance(suite.ctx, devRewardsReceiver, denom).Amount)

	// since the NFT incentives are not setup yet, funds go into the communtiy pool
	feePool = suite.app.DistrKeeper.GetFeePool(suite.ctx)
	suite.Equal(
		mintCoin.Amount.ToDec().Mul(params.DistributionProportions.NftIncentives),
		feePool.CommunityPool.AmountOf(denom))
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
	allocKeeper.DistributeInflation(suite.ctx)

	// should be 0
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).IsZero())

	msgServer := keeper.NewMsgServerImpl(allocKeeper)

	// fundAccount
	FundAccount(suite.app.BankKeeper, suite.ctx, addr1, fundAmount)
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
	_, err := msgServer.FundFairburnPool(sdk.WrapSDKContext(suite.ctx), types.NewMsgFundFairburnPool(addr1, fundAmount))
	suite.NoError(err)

	// should have funds now
	suite.Require().Equal(fundAmount.String(), suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).String())
	// still 0
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).IsZero())

	allocKeeper.DistributeInflation(suite.ctx)

	// fee collector should have funds now
	suite.Require().Equal(fundAmount.String(), suite.app.BankKeeper.GetBalance(suite.ctx, feeCollector, denom).String())
	// fairburn pool should be 0
	suite.Require().True(suite.app.BankKeeper.GetBalance(suite.ctx, fairBurnPool, denom).IsZero())
}
