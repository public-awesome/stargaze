package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v8/app"
	"github.com/public-awesome/stargaze/v8/testutil/simapp"
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
}
