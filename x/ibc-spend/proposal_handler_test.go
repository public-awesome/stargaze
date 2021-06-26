package ibcspend_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	ibctesting "github.com/cosmos/cosmos-sdk/x/ibc/testing"
	"github.com/public-awesome/stargaze/simapp"
	ibcspend "github.com/public-awesome/stargaze/x/ibc-spend"
	"github.com/public-awesome/stargaze/x/ibc-spend/types"
)

var (
	delPk1   = ed25519.GenPrivKey().PubKey()
	delAddr1 = sdk.AccAddress(delPk1.Address())

	amount = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1)))
)

func testProposal(
	recipient sdk.AccAddress,
	amount sdk.Coins,
	sourceChannel string) *types.CommunityPoolIBCSpendProposal {

	return types.NewCommunityPoolIBCSpendProposal("Test", "description", "recipient", amount, sourceChannel, 0)
}

type TransferTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
}

func (suite *TransferTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))
}

func (suite *TransferTestSuite) TestProposalHandlerPassed() {
	// app := simapp.Setup(false)
	// TODO: does not have IBCSpendKeeper in this SimApp...
	// have to copy ibc testing framework into module..
	app := suite.chainA.App
	// ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx := suite.chainA.GetContext()
	recipient := delAddr1

	// setup IBC
	_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
	// channelA, _ := suite.coordinator.CreateTransferChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
	suite.coordinator.CreateTransferChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

	// add coins to the module account
	macc := app.DistrKeeper.GetDistributionAccount(ctx)
	balances := app.BankKeeper.GetAllBalances(ctx, macc.GetAddress())
	err := app.BankKeeper.SetBalances(ctx, macc.GetAddress(), balances.Add(amount...))
	suite.Require().NoError(err)

	app.AccountKeeper.SetModuleAccount(ctx, macc)

	account := app.AccountKeeper.NewAccountWithAddress(ctx, recipient)
	app.AccountKeeper.SetAccount(ctx, account)
	suite.Require().True(app.BankKeeper.GetAllBalances(ctx, account.GetAddress()).IsZero())

	feePool := app.DistrKeeper.GetFeePool(ctx)
	feePool.CommunityPool = sdk.NewDecCoinsFromCoins(amount...)
	app.DistrKeeper.SetFeePool(ctx, feePool)

	// tp := testProposal(recipient, amount, channelA.ID)
	// hdlr := ibcspend.NewCommunityPoolIBCSpendProposalHandler(app.IBCSpendKeeper)
	// suite.Require().NoError(hdlr(ctx, tp))

	// balances = app.BankKeeper.GetAllBalances(ctx, recipient)
	// suite.Require().Equal(balances, amount)
}

func (suite *TransferTestSuite) TestProposalHandlerFailed_InsufficientCoinsInCommunityPool() {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	recipient := delAddr1

	account := app.AccountKeeper.NewAccountWithAddress(ctx, recipient)
	app.AccountKeeper.SetAccount(ctx, account)
	suite.Require().True(app.BankKeeper.GetAllBalances(ctx, account.GetAddress()).IsZero())

	tp := testProposal(recipient, amount, "sourceChannel")
	hdlr := ibcspend.NewCommunityPoolIBCSpendProposalHandler(app.IBCSpendKeeper)
	suite.Require().Error(hdlr(ctx, tp))

	balances := app.BankKeeper.GetAllBalances(ctx, recipient)
	suite.Require().True(balances.IsZero())
}

func (suite *TransferTestSuite) TestProposalHandlerFailed_IBCNotSetup() {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	recipient := delAddr1

	// add coins to the module account
	macc := app.DistrKeeper.GetDistributionAccount(ctx)
	balances := app.BankKeeper.GetAllBalances(ctx, macc.GetAddress())
	err := app.BankKeeper.SetBalances(ctx, macc.GetAddress(), balances.Add(amount...))
	suite.Require().NoError(err)

	app.AccountKeeper.SetModuleAccount(ctx, macc)

	account := app.AccountKeeper.NewAccountWithAddress(ctx, recipient)
	app.AccountKeeper.SetAccount(ctx, account)
	suite.Require().True(app.BankKeeper.GetAllBalances(ctx, account.GetAddress()).IsZero())

	feePool := app.DistrKeeper.GetFeePool(ctx)
	feePool.CommunityPool = sdk.NewDecCoinsFromCoins(amount...)
	app.DistrKeeper.SetFeePool(ctx, feePool)

	tp := testProposal(recipient, amount, "sourceChannel")
	hdlr := ibcspend.NewCommunityPoolIBCSpendProposalHandler(app.IBCSpendKeeper)
	suite.Require().Error(hdlr(ctx, tp))
}

func TestTransferTestSuite(t *testing.T) {
	suite.Run(t, new(TransferTestSuite))
}
