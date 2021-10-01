package keeper_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/app"
	"github.com/public-awesome/stargaze/testutil/simapp"
	"github.com/public-awesome/stargaze/x/claim/types"
	"github.com/stretchr/testify/suite"
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
	suite.app.ClaimKeeper.CreateModuleAccount(suite.ctx, sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000000)))
}

func (s *KeeperTestSuite) TestKeeper() {
	app, ctx := s.app, s.ctx
	moduleAddress := app.AccountKeeper.GetModuleAddress(types.ModuleName)
	balance := app.BankKeeper.GetBalance(ctx, moduleAddress, sdk.DefaultBondDenom)
	s.Require().Equal(fmt.Sprintf("10000000%s", sdk.DefaultBondDenom), balance.String())

}
func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
