package app

import (
	"fmt"
	"time"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibchooks "github.com/public-awesome/stargaze/v11/x/ibchooks/types"
	tokenfactorytypes "github.com/public-awesome/stargaze/v11/x/tokenfactory/types"
)

// next upgrade name
const upgradeName = "v11"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *App) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(upgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// run migrations before modifying state
		migrations, err := app.mm.RunMigrations(ctx, cfg, fromVM)
		if err != nil {
			return nil, err
		}
		params := app.TokenFactoryKeeper.GetParams(ctx)
		params.DenomCreationFee = nil
		params.DenomCreationGasConsume = 50_000_000 // 50STARS at 1ustars
		app.TokenFactoryKeeper.SetParams(ctx, params)

		// change mint params to include the new supplement amount
		mintParams := app.MintKeeper.GetParams(ctx)
		mintParams.InitialAnnualProvisions = sdk.NewDec(267_000_000_000_000) // 267M
		mintParams.StartTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)   // 2023-01-01
		mintParams.SupplementAmount = sdk.NewCoins()                         // set amount
		app.MintKeeper.SetParams(ctx, mintParams)

		// set community tax to 0 since the allocation params now will take care of it
		distributionParams := app.DistrKeeper.GetParams(ctx)
		distributionParams.CommunityTax = sdk.ZeroDec()
		app.DistrKeeper.SetParams(ctx, distributionParams)

		// change alloc params to set nft incentives to 0% until incentives are live
		allocParams := app.AllocKeeper.GetParams(ctx)

		// distribution proportions
		proportions := allocParams.DistributionProportions
		proportions.NftIncentives = sdk.ZeroDec() // nft incentives to 0%
		allocParams.DistributionProportions = proportions
		app.AllocKeeper.SetParams(ctx, allocParams)

		return migrations, nil
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{tokenfactorytypes.ModuleName, ibchooks.ModuleName},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
