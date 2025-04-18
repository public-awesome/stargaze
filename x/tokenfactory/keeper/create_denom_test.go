package keeper_test

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestMsgCreateDenom() {
	var (
		tokenFactoryKeeper = suite.App.Keepers.TokenFactoryKeeper
		bankKeeper         = suite.App.Keepers.BankKeeper
		denomCreationFee   = tokenFactoryKeeper.GetParams(suite.Ctx).DenomCreationFee
	)

	// Get balance of acc 0 before creating a denom
	preCreateBalance := bankKeeper.GetBalance(suite.Ctx, suite.TestAccs[0], denomCreationFee[0].Denom)

	// Creating a denom should work
	res, err := suite.msgServer.CreateDenom(suite.Ctx, types.NewMsgCreateDenom(suite.TestAccs[0].String(), "testy"))
	suite.Require().NoError(err)
	suite.Require().NotEmpty(res.GetNewTokenDenom())

	// Make sure that the admin is set correctly
	queryRes, err := suite.queryClient.DenomAuthorityMetadata(suite.Ctx.Context(), &types.QueryDenomAuthorityMetadataRequest{
		Denom: res.GetNewTokenDenom(),
	})
	suite.Require().NoError(err)
	suite.Require().Equal(suite.TestAccs[0].String(), queryRes.AuthorityMetadata.Admin)

	// Make sure that creation fee was deducted
	postCreateBalance := bankKeeper.GetBalance(suite.Ctx, suite.TestAccs[0], tokenFactoryKeeper.GetParams(suite.Ctx).DenomCreationFee[0].Denom)
	suite.Require().True(preCreateBalance.Sub(postCreateBalance).IsEqual(denomCreationFee[0]))

	// Make sure that a second version of the same denom can't be recreated
	_, err = suite.msgServer.CreateDenom(suite.Ctx, types.NewMsgCreateDenom(suite.TestAccs[0].String(), "testy"))
	suite.Require().Error(err)

	// Creating a second denom should work
	res, err = suite.msgServer.CreateDenom(suite.Ctx, types.NewMsgCreateDenom(suite.TestAccs[0].String(), "minty"))
	suite.Require().NoError(err)
	suite.Require().NotEmpty(res.GetNewTokenDenom())

	// Try querying all the denoms created by suite.TestAccs[0]
	queryRes2, err := suite.queryClient.DenomsFromCreator(suite.Ctx.Context(), &types.QueryDenomsFromCreatorRequest{
		Creator: suite.TestAccs[0].String(),
	})
	suite.Require().NoError(err)
	suite.Require().Len(queryRes2.Denoms, 2)

	// Make sure that a second account can create a denom with the same subdenom
	res, err = suite.msgServer.CreateDenom(suite.Ctx, types.NewMsgCreateDenom(suite.TestAccs[1].String(), "testy"))
	suite.Require().NoError(err)
	suite.Require().NotEmpty(res.GetNewTokenDenom())

	// Make sure that an address with a "/" in it can't create denoms
	_, err = suite.msgServer.CreateDenom(suite.Ctx, types.NewMsgCreateDenom("stargaze.stars/creator", "testy"))
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestCreateDenom() {
	var (
		primaryDenom            = types.DefaultParams().DenomCreationFee[0].Denom
		defaultDenomCreationFee = types.Params{DenomCreationFee: sdk.NewCoins(sdk.NewCoin(primaryDenom, math.NewInt(50000000)))}
		twoDenomCreationFee     = types.Params{DenomCreationFee: sdk.NewCoins(sdk.NewCoin(primaryDenom, math.NewInt(50000000)), sdk.NewCoin("utest", math.NewInt(50000000)))}
		nilCreationFee          = types.Params{DenomCreationFee: nil}
		largeCreationFee        = types.Params{DenomCreationFee: sdk.NewCoins(sdk.NewCoin(primaryDenom, math.NewInt(5000000000)))}
	)

	for _, tc := range []struct {
		desc             string
		denomCreationFee types.Params
		setup            func()
		subdenom         string
		valid            bool
	}{
		{
			desc:             "subdenom too long",
			denomCreationFee: defaultDenomCreationFee,
			subdenom:         "assadsadsadasdasdsadsadsadsadsadsadsklkadaskkkdasdasedskhanhassyeunganassfnlksdflksafjlkasd",
			valid:            false,
		},
		{
			desc:             "subdenom and creator pair already exists",
			denomCreationFee: defaultDenomCreationFee,
			setup: func() {
				_, err := suite.msgServer.CreateDenom(suite.Ctx, types.NewMsgCreateDenom(suite.TestAccs[0].String(), "testy"))
				suite.Require().NoError(err)
			},
			subdenom: "testy",
			valid:    false,
		},
		{
			desc:             "success case: defaultDenomCreationFee",
			denomCreationFee: defaultDenomCreationFee,
			subdenom:         "evmos",
			valid:            true,
		},
		{
			desc:             "success case: twoDenomCreationFee",
			denomCreationFee: twoDenomCreationFee,
			subdenom:         "catcoin",
			valid:            true,
		},
		{
			desc:             "success case: nilCreationFee",
			denomCreationFee: nilCreationFee,
			subdenom:         "bzzzcoin",
			valid:            true,
		},
		{
			desc:             "account doesn't have enough to pay for denom creation fee",
			denomCreationFee: largeCreationFee,
			subdenom:         "tooexpensive",
			valid:            false,
		},
		{
			desc:             "subdenom having invalid characters",
			denomCreationFee: defaultDenomCreationFee,
			subdenom:         "bit/***///&&&/coin",
			valid:            false,
		},
	} {
		suite.SetupTest()
		suite.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			if tc.setup != nil {
				tc.setup()
			}
			tokenFactoryKeeper := suite.App.Keepers.TokenFactoryKeeper
			bankKeeper := suite.App.Keepers.BankKeeper
			// Set denom creation fee in params
			err := tokenFactoryKeeper.SetParams(suite.Ctx, tc.denomCreationFee)
			suite.Require().NoError(err)
			denomCreationFee := tokenFactoryKeeper.GetParams(suite.Ctx).DenomCreationFee
			suite.Require().Equal(tc.denomCreationFee.DenomCreationFee, denomCreationFee)

			// note balance, create a tokenfactory denom, then note balance again
			preCreateBalance := bankKeeper.GetAllBalances(suite.Ctx, suite.TestAccs[0])
			res, err := suite.msgServer.CreateDenom(suite.Ctx, types.NewMsgCreateDenom(suite.TestAccs[0].String(), tc.subdenom))
			postCreateBalance := bankKeeper.GetAllBalances(suite.Ctx, suite.TestAccs[0])
			if tc.valid {
				suite.Require().NoError(err)
				suite.Require().True(preCreateBalance.Sub(postCreateBalance...).Equal(denomCreationFee))

				// Make sure that the admin is set correctly
				queryRes, err := suite.queryClient.DenomAuthorityMetadata(suite.Ctx.Context(), &types.QueryDenomAuthorityMetadataRequest{
					Denom: res.GetNewTokenDenom(),
				})

				suite.Require().NoError(err)
				suite.Require().Equal(suite.TestAccs[0].String(), queryRes.AuthorityMetadata.Admin)

			} else {
				suite.Require().Error(err)
				// Ensure we don't charge if we expect an error
				suite.Require().True(preCreateBalance.Equal(postCreateBalance))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGasConsume() {
	// It's hard to estimate exactly how much gas will be consumed when creating a
	// denom, because besides consuming the gas specified by the params, the keeper
	// also does a bunch of other things that consume gas.
	//
	// Rather, we test whether the gas consumed is within a range. Specifically,
	// the range [gasConsume, gasConsume + offset]. If the actual gas consumption
	// falls within the range for all test cases, we consider the test passed.
	//
	// In experience, the total amount of gas consumed should consume be ~30k more
	// than the set amount.
	const offset = 50000

	for _, tc := range []struct {
		desc       string
		gasConsume uint64
	}{
		{
			desc:       "gas consume zero",
			gasConsume: 0,
		},
		{
			desc:       "gas consume 1,000,000",
			gasConsume: 1_000_000,
		},
		{
			desc:       "gas consume 10,000,000",
			gasConsume: 10_000_000,
		},
		{
			desc:       "gas consume 25,000,000",
			gasConsume: 25_000_000,
		},
		{
			desc:       "gas consume 50,000,000",
			gasConsume: 50_000_000,
		},
		{
			desc:       "gas consume 200,000,000",
			gasConsume: 200_000_000,
		},
	} {
		suite.SetupTest()
		suite.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			// set params with the gas consume amount
			err := suite.App.Keepers.TokenFactoryKeeper.SetParams(suite.Ctx, types.NewParams(nil, tc.gasConsume))
			suite.Require().NoError(err)

			// amount of gas consumed prior to the denom creation
			gasConsumedBefore := suite.Ctx.GasMeter().GasConsumed()

			// create a denom
			_, err = suite.msgServer.CreateDenom(suite.Ctx, types.NewMsgCreateDenom(suite.TestAccs[0].String(), "larry"))
			suite.Require().NoError(err)

			// amount of gas consumed after the denom creation
			gasConsumedAfter := suite.Ctx.GasMeter().GasConsumed()

			// the amount of gas consumed must be within the range
			gasConsumed := gasConsumedAfter - gasConsumedBefore
			suite.Require().Greater(gasConsumed, tc.gasConsume)
			suite.Require().Less(gasConsumed, tc.gasConsume+offset)
		})
	}
}
