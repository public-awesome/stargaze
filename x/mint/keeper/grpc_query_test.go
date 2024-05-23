package keeper_test

import (
	gocontext "context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v14/app"
	"github.com/public-awesome/stargaze/v14/testutil/simapp"
	"github.com/public-awesome/stargaze/v14/x/mint/keeper"
	"github.com/public-awesome/stargaze/v14/x/mint/types"
)

type MintTestSuite struct {
	suite.Suite

	app         *app.App
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *MintTestSuite) SetupTest() {
	app := simapp.New(suite.T())
	ctx := app.BaseApp.NewContext(false)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, keeper.NewQueryServer(app.Keepers.MintKeeper))
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.ctx = ctx

	suite.queryClient = queryClient
}

func (suite *MintTestSuite) TestGRPCParams() {
	app, ctx, queryClient := suite.app, suite.ctx, suite.queryClient

	params, err := queryClient.Params(gocontext.Background(), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	p, err := app.Keepers.MintKeeper.GetParams(ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(params.Params, p)

	annualProvisions, err := queryClient.AnnualProvisions(gocontext.Background(), &types.QueryAnnualProvisionsRequest{})
	suite.Require().NoError(err)
	minter, err := app.Keepers.MintKeeper.GetMinter(ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(annualProvisions.AnnualProvisions, minter.AnnualProvisions)
}

func TestMintTestSuite(t *testing.T) {
	suite.Run(t, new(MintTestSuite))
}
