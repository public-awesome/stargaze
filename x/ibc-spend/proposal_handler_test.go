package ibcspend_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	ibcspend "github.com/public-awesome/stargaze/x/ibc-spend"
	ibctesting "github.com/public-awesome/stargaze/x/ibc-spend/testing"
	"github.com/public-awesome/stargaze/x/ibc-spend/types"
)

var (
	delPk1   = ed25519.GenPrivKey().PubKey()
	delAddr1 = sdk.AccAddress(delPk1.Address())

	amount = sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1)))
)

func testProposal(
	recipient sdk.AccAddress,
	amount sdk.Coins,
	sourceChannel string) *types.CommunityPoolIBCSpendProposal {

	return types.NewCommunityPoolIBCSpendProposal("Test", "description", recipient.String(), amount, sourceChannel, 100)
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
	app := suite.chainA.App
	ctx := suite.chainA.GetContext()
	// recipient := delAddr1

	// setup IBC
	_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
	channelA, _ := suite.coordinator.CreateTransferChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

	// create the distribution module account
	macc1 := app.DistrKeeper.GetDistributionAccount(ctx)
	fmt.Printf("dist module account %v\n", macc1.String())
	// create the ibcspend module account
	macc2 := app.IBCSpendKeeper.GetIBCSpendAccount(ctx)
	fmt.Printf("ibcspend module account %v\n", macc2.String())

	// have to fund distribution keeper module account since it
	// includes the community pool
	balances := app.BankKeeper.GetAllBalances(ctx, macc1.GetAddress())
	err := app.BankKeeper.SetBalances(ctx, macc1.GetAddress(), balances.Add(amount...))
	suite.Require().NoError(err)

	feePool := app.DistrKeeper.GetFeePool(ctx)
	feePool.CommunityPool = sdk.NewDecCoinsFromCoins(amount...)
	app.DistrKeeper.SetFeePool(ctx, feePool)

	recipient := suite.chainB.SenderAccount.GetAddress()
	tp := testProposal(recipient, amount, channelA.ID)
	hdlr := ibcspend.NewCommunityPoolIBCSpendProposalHandler(app.IBCSpendKeeper)
	suite.Require().NoError(hdlr(ctx, tp))

	// check if recipient on chain B has balance of amount
	balances = suite.chainB.App.BankKeeper.GetAllBalances(suite.chainB.GetContext(), recipient)
	// TODO: check if denom trace matches what's expected
	suite.Require().Equal(amount, balances)

	// TODO: send coin back to chain A
	// check if denom changed back to ustarx
}

// func (suite *TransferTestSuite) TestProposalHandlerFailed_InsufficientCoinsInCommunityPool() {
// app := simapp.Setup(false)
// ctx := app.BaseApp.NewContext(false, tmproto.Header{})

// recipient := delAddr1

// account := app.AccountKeeper.NewAccountWithAddress(ctx, recipient)
// app.AccountKeeper.SetAccount(ctx, account)
// suite.Require().True(app.BankKeeper.GetAllBalances(ctx, account.GetAddress()).IsZero())

// tp := testProposal(recipient, amount, "sourceChannel")
// hdlr := ibcspend.NewCommunityPoolIBCSpendProposalHandler(app.IBCSpendKeeper)
// suite.Require().Error(hdlr(ctx, tp))

// balances := app.BankKeeper.GetAllBalances(ctx, recipient)
// suite.Require().True(balances.IsZero())
// }

// func (suite *TransferTestSuite) TestProposalHandlerFailed_IBCNotSetup() {
// app := simapp.Setup(false)
// ctx := app.BaseApp.NewContext(false, tmproto.Header{})

// recipient := delAddr1

// add coins to the module account
// macc := app.DistrKeeper.GetDistributionAccount(ctx)
// balances := app.BankKeeper.GetAllBalances(ctx, macc.GetAddress())
// err := app.BankKeeper.SetBalances(ctx, macc.GetAddress(), balances.Add(amount...))
// suite.Require().NoError(err)

// app.AccountKeeper.SetModuleAccount(ctx, macc)

// account := app.AccountKeeper.NewAccountWithAddress(ctx, recipient)
// app.AccountKeeper.SetAccount(ctx, account)
// suite.Require().True(app.BankKeeper.GetAllBalances(ctx, account.GetAddress()).IsZero())

// feePool := app.DistrKeeper.GetFeePool(ctx)
// feePool.CommunityPool = sdk.NewDecCoinsFromCoins(amount...)
// app.DistrKeeper.SetFeePool(ctx, feePool)

// tp := testProposal(recipient, amount, "sourceChannel")
// hdlr := ibcspend.NewCommunityPoolIBCSpendProposalHandler(app.IBCSpendKeeper)
// suite.Require().Error(hdlr(ctx, tp))
// }

func TestTransferTestSuite(t *testing.T) {
	suite.Run(t, new(TransferTestSuite))
}
