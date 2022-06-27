package keeper_test

import (
	"github.com/public-awesome/stargaze/v6/x/claim/types"
)

func (s *KeeperTestSuite) TestExportGenesis() {
	app, ctx := s.app, s.ctx
	app.ClaimKeeper.InitGenesis(ctx, *types.DefaultGenesis())
	// app.ClaimKeeper.SetParams(ctx, types.DefaultParams())
	exported := app.ClaimKeeper.ExportGenesis(ctx)
	params := types.DefaultParams()
	params.AirdropStartTime = ctx.BlockTime()
	s.Require().Equal(params.AllowedClaimers, exported.Params.AllowedClaimers)
}
