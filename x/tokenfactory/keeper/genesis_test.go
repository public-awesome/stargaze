package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/public-awesome/stargaze/v14/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	genesisState := types.GenesisState{
		FactoryDenoms: []types.GenesisDenom{
			{
				Denom: "factory/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/bitcoin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f",
				},
			},
			{
				Denom: "factory/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/diff-admin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "cosmos15czt5nhlnvayqq37xun9s9yus0d6y26dx74r5p",
				},
			},
			{
				Denom: "factory/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/litecoin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f",
				},
			},
		},
	}

	suite.SetupTestForInitGenesis()
	app := suite.App

	// Test both with bank denom metadata set, and not set.
	for i, denom := range genesisState.FactoryDenoms {
		// hacky, sets bank metadata to exist if i != 0, to cover both cases.
		if i != 0 {
			app.Keepers.BankKeeper.SetDenomMetaData(suite.Ctx, banktypes.Metadata{Base: denom.GetDenom()})
		}
	}

	// check before initGenesis that the module account is nil
	tokenfactoryModuleAccount := app.Keepers.AccountKeeper.GetAccount(suite.Ctx, app.Keepers.AccountKeeper.GetModuleAddress(types.ModuleName))
	suite.Require().Nil(tokenfactoryModuleAccount)

	err := app.Keepers.TokenFactoryKeeper.SetParams(suite.Ctx, types.Params{DenomCreationFee: sdk.Coins{sdk.NewInt64Coin("utest", 100)}})
	suite.Require().NoError(err)
	app.Keepers.TokenFactoryKeeper.InitGenesis(suite.Ctx, genesisState)

	// check that the module account is now initialized
	tokenfactoryModuleAccount = app.Keepers.AccountKeeper.GetAccount(suite.Ctx, app.Keepers.AccountKeeper.GetModuleAddress(types.ModuleName))
	suite.Require().NotNil(tokenfactoryModuleAccount)

	exportedGenesis := app.Keepers.TokenFactoryKeeper.ExportGenesis(suite.Ctx)
	suite.Require().NotNil(exportedGenesis)
	suite.Require().Equal(genesisState, *exportedGenesis)
}
