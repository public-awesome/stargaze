package keeper_test

import (
	"fmt"
	"testing"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/public-awesome/stargaze/v6/app"
	"github.com/public-awesome/stargaze/v6/testutil/simapp"
	"github.com/public-awesome/stargaze/v6/x/claim/keeper"
	"github.com/public-awesome/stargaze/v6/x/claim/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx     sdk.Context
	app     *app.App
	msgSrvr types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = simapp.New(suite.T().TempDir())
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{Height: 2, ChainID: "stargaze-1", Time: time.Now().UTC()})
	suite.app.ClaimKeeper.CreateModuleAccount(suite.ctx, sdk.NewCoin(types.DefaultClaimDenom, sdk.NewInt(10000000)))
	startTime := time.Now()

	suite.msgSrvr = keeper.NewMsgServerImpl(suite.app.ClaimKeeper)
	suite.app.ClaimKeeper.SetParams(suite.ctx, types.Params{
		AirdropEnabled:     true,
		AirdropStartTime:   startTime,
		DurationUntilDecay: types.DefaultDurationUntilDecay,
		DurationOfDecay:    types.DefaultDurationOfDecay,
		ClaimDenom:         types.DefaultClaimDenom,
	})
	suite.ctx = suite.ctx.WithBlockTime(startTime)
}

func (s *KeeperTestSuite) TestModuleAccountCreated() {
	app, ctx := s.app, s.ctx
	moduleAddress := app.AccountKeeper.GetModuleAddress(types.ModuleName)
	balance := app.BankKeeper.GetBalance(ctx, moduleAddress, types.DefaultClaimDenom)
	s.Require().Equal(fmt.Sprintf("10000000%s", types.DefaultClaimDenom), balance.String())
}

func (s *KeeperTestSuite) TestClaimFor() {
	pub1 := secp256k1.GenPrivKey().PubKey()
	pub2 := secp256k1.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pub1.Address())
	addr2 := sdk.AccAddress(pub2.Address())
	contractAddress1 := wasmkeeper.BuildContractAddress(1, 1)
	contractAddress2 := wasmkeeper.BuildContractAddress(1, 2)
	claimRecords := []types.ClaimRecord{
		{
			Address:                addr1.String(),
			InitialClaimableAmount: sdk.NewCoins(sdk.NewInt64Coin(types.DefaultClaimDenom, 2000)),
			ActionCompleted:        []bool{false, false, false, false, false},
		},
		{
			Address:                addr2.String(),
			InitialClaimableAmount: sdk.NewCoins(sdk.NewInt64Coin(types.DefaultClaimDenom, 2000)),
			ActionCompleted:        []bool{false, false, false, false, false},
		},
	}

	s.app.AccountKeeper.SetAccount(s.ctx, authtypes.NewBaseAccount(addr1, nil, 0, 0))
	s.app.AccountKeeper.SetAccount(s.ctx, authtypes.NewBaseAccount(addr2, nil, 0, 0))

	s.app.ClaimKeeper.SetParams(s.ctx, types.Params{
		AirdropEnabled:     false,
		AirdropStartTime:   time.Now().Add(time.Hour * -1),
		ClaimDenom:         types.DefaultClaimDenom,
		DurationUntilDecay: time.Hour,
		DurationOfDecay:    time.Hour * 4,
		AllowedClaimers:    make([]types.ClaimAuthorization, 0),
	})
	err := s.app.ClaimKeeper.SetClaimRecords(s.ctx, claimRecords)
	s.Require().NoError(err)
	msgClaimFor := types.NewMsgClaimFor(contractAddress1.String(), addr1.String(), types.ActionBidNFT)
	ctx := sdk.WrapSDKContext(s.ctx)
	_, err = s.msgSrvr.ClaimFor(ctx, msgClaimFor)
	s.Error(err)
	s.Contains(err.Error(), "airdrop not enabled")
	s.app.ClaimKeeper.SetParams(s.ctx, types.Params{
		AirdropEnabled:     true,
		AirdropStartTime:   time.Now().Add(time.Hour * -1),
		ClaimDenom:         types.DefaultClaimDenom,
		DurationUntilDecay: time.Hour,
		DurationOfDecay:    time.Hour * 4,
		AllowedClaimers: []types.ClaimAuthorization{
			{
				ContractAddress: contractAddress1.String(),
				Action:          types.ActionBidNFT,
			},
			{
				ContractAddress: contractAddress2.String(),
				Action:          types.ActionMintNFT,
			},
		},
	})

	// unauthorized
	msgClaimFor = types.NewMsgClaimFor(wasmkeeper.BuildContractAddress(2, 1).String(), addr1.String(), types.ActionMintNFT)
	_, err = s.msgSrvr.ClaimFor(ctx, msgClaimFor)
	s.Require().Error(err)
	s.Contains(err.Error(), "address is not allowed to claim")

	// unauthorized to claim another action

	msgClaimFor = types.NewMsgClaimFor(contractAddress1.String(), addr1.String(), types.ActionMintNFT)
	_, err = s.msgSrvr.ClaimFor(ctx, msgClaimFor)
	s.Require().Error(err)
	s.Contains(err.Error(), "address is not allowed to claim")

	// claim
	msgClaimFor = types.NewMsgClaimFor(contractAddress1.String(), addr1.String(), types.ActionBidNFT)
	_, err = s.msgSrvr.ClaimFor(ctx, msgClaimFor)
	s.Require().NoError(err)

	claimedCoins := s.app.BankKeeper.GetAllBalances(s.ctx, addr1)
	s.Require().Equal(claimedCoins.AmountOf(types.DefaultClaimDenom), claimRecords[0].InitialClaimableAmount.AmountOf(types.DefaultClaimDenom).Quo(sdk.NewInt(5)))

	record, err := s.app.ClaimKeeper.GetClaimRecord(s.ctx, addr1)
	s.Require().NoError(err)
	s.Require().True(record.ActionCompleted[1])
	s.Require().Equal([]bool{false, true, false, false, false}, record.ActionCompleted)

	// claim 2
	msgClaimFor = types.NewMsgClaimFor(contractAddress2.String(), addr1.String(), types.ActionMintNFT)
	_, err = s.msgSrvr.ClaimFor(ctx, msgClaimFor)
	s.Require().NoError(err)

	claimedCoins = s.app.BankKeeper.GetAllBalances(s.ctx, addr1)
	s.Require().Equal(
		claimedCoins.AmountOf(types.DefaultClaimDenom).String(),
		claimRecords[0].InitialClaimableAmount.AmountOf(types.DefaultClaimDenom).Quo(sdk.NewInt(5)).Mul(sdk.NewInt(2)).String(), // 2 actions claimed
	)

	record, err = s.app.ClaimKeeper.GetClaimRecord(s.ctx, addr1)
	s.Require().NoError(err)
	s.Require().True(record.ActionCompleted[1])
	s.Require().True(record.ActionCompleted[2])
	s.Require().Equal([]bool{false, true, true, false, false}, record.ActionCompleted)

	// claim second address
	msgClaimFor = types.NewMsgClaimFor(contractAddress2.String(), addr2.String(), types.ActionMintNFT)
	_, err = s.msgSrvr.ClaimFor(ctx, msgClaimFor)
	s.Require().NoError(err)

	claimedCoins = s.app.BankKeeper.GetAllBalances(s.ctx, addr2)
	s.Require().Equal(
		claimedCoins.AmountOf(types.DefaultClaimDenom).String(),
		claimRecords[0].InitialClaimableAmount.AmountOf(types.DefaultClaimDenom).Quo(sdk.NewInt(5)).String(), // 1 action claimed
	)

	record, err = s.app.ClaimKeeper.GetClaimRecord(s.ctx, addr2)
	s.Require().NoError(err)
	s.Require().False(record.ActionCompleted[1])
	s.Require().True(record.ActionCompleted[2])
	s.Require().Equal([]bool{false, false, true, false, false}, record.ActionCompleted)
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
