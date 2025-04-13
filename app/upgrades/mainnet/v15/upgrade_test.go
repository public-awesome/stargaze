package v15_test

import (
	"fmt"
	"testing"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	stargazeapp "github.com/public-awesome/stargaze/v16/app"
	"github.com/public-awesome/stargaze/v16/testutil/simapp"
	"github.com/stretchr/testify/suite"
)

const (
	dummyUpgradeHeight = 5
)

type UpgradeTestSuite struct {
	suite.Suite

	App *stargazeapp.App
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (suite *UpgradeTestSuite) SetupTest() {
	suite.App = simapp.New(suite.T())
}

func (suite *UpgradeTestSuite) TestUpgrade() {
	testCases := []struct {
		name        string
		preUpgrade  func()
		postUpgrade func()
	}{
		{
			"Ensure any state transitions are handled correctly during the upgrade process",
			func() {
				// test conditions setup
			},
			func() {
				// post upgrade conditions check
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			tc.preUpgrade()

			ctx := suite.App.BaseApp.NewContext(false).WithBlockHeight(dummyUpgradeHeight - 1)
			plan := upgradetypes.Plan{Name: "v14", Height: dummyUpgradeHeight}
			upgradekeeper := suite.App.Keepers.UpgradeKeeper
			err := upgradekeeper.ScheduleUpgrade(ctx, plan)
			suite.Require().NoError(err)
			_, err = upgradekeeper.GetUpgradePlan(ctx)
			suite.Require().NoError(err)
			ctx = ctx.WithBlockHeight(dummyUpgradeHeight)
			_, err = suite.App.BeginBlocker(ctx)
			suite.Require().NoError(err)

			tc.postUpgrade()
		})
	}
}
