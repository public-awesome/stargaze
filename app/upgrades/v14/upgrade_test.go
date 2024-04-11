package v14_test

import (
	"fmt"
	"testing"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	stargazeapp "github.com/public-awesome/stargaze/v14/app"
	"github.com/public-awesome/stargaze/v14/testutil/simapp"
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
		name         string
		pre_upgrade  func()
		post_upgrade func()
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

			tc.pre_upgrade()

			ctx := suite.App.BaseApp.NewContext(false).WithBlockHeight(dummyUpgradeHeight - 1)
			plan := upgradetypes.Plan{Name: "v14", Height: dummyUpgradeHeight}
			upgradekeeper := suite.App.UpgradeKeeper
			err := upgradekeeper.ScheduleUpgrade(ctx, plan)
			suite.Require().NoError(err)
			_, err = upgradekeeper.GetUpgradePlan(ctx)
			suite.Require().NoError(err)
			ctx = ctx.WithBlockHeight(dummyUpgradeHeight)
			suite.Require().NotPanics(func() {
				suite.App.BeginBlocker(ctx)
			})

			tc.post_upgrade()
		})
	}
}
