package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v8/testutil/sample"
)

func (s *KeeperTestSuite) Test_SetPrivileged() {
	app, ctx := s.app, s.ctx
	acc1 := sample.AccAddress()

	app.CronKeeper.SetPrivileged(ctx, acc1)

	s.Require().True(app.CronKeeper.IsPrivileged(ctx, acc1), "expected %s to be privileged", acc1)
}

func (s *KeeperTestSuite) Test_UnsetPrivileged() {
	app, ctx := s.app, s.ctx
	acc1 := sample.AccAddress()

	app.CronKeeper.SetPrivileged(ctx, acc1)
	app.CronKeeper.UnsetPrivileged(ctx, acc1)

	s.Require().False(app.CronKeeper.IsPrivileged(ctx, acc1), "expected %s to not be privileged", acc1)
}

func (s *KeeperTestSuite) Test_IteratePrivileged() {
	app, ctx := s.app, s.ctx
	acc1 := sample.AccAddress()
	acc2 := sample.AccAddress()
	acc3 := sample.AccAddress()
	// Setting three contracts as privileged.
	app.CronKeeper.SetPrivileged(ctx, acc1)
	app.CronKeeper.SetPrivileged(ctx, acc2)
	app.CronKeeper.SetPrivileged(ctx, acc3)
	// Removing one contract as privileged
	app.CronKeeper.UnsetPrivileged(ctx, acc2)

	var contracts []sdk.AccAddress
	app.CronKeeper.IteratePrivileged(ctx, func(addr sdk.AccAddress) bool {
		contracts = append(contracts, addr)
		return false
	})
	s.Require().Len(contracts, 2, "expected 2 privileged contracts, got %d", len(contracts))
}
