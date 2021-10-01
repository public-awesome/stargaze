package keeper_test

import (
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/app"
	"github.com/public-awesome/stargaze/testutil/simapp"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx sdk.Context
	app *app.App
}

func (suite *KeeperTestSuite) SetupTest() {

	simapp.New(os.TempDir())
}
