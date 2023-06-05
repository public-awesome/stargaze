package app

import (
	"fmt"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	crontypes "github.com/public-awesome/stargaze/v10/x/cron/types"
	globalfeetypes "github.com/public-awesome/stargaze/v10/x/globalfee/types"
)

// next upgrade name
const upgradeName = "v10"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *App) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(upgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// run migrations before modifying state
		migrations, err := app.mm.RunMigrations(ctx, cfg, fromVM)
		if err != nil {
			return nil, err
		}
		// set global fee params to 1ustars
		stakeDenom := app.StakingKeeper.BondDenom(ctx)
		globalfeeParams := app.GlobalFeeKeeper.GetParams(ctx)
		globalfeeParams.MinimumGasPrices = sdk.NewDecCoins(sdk.NewDecCoinFromDec(stakeDenom, sdk.OneDec()))

		globalfeeParams.PrivilegedAddresses = []string{
			// Multisig that is in charge of managing Launchpad Factories until a DAO is created
			// https://www.mintscan.io/stargaze/account/stars1zmaulflshft579x37acrdad72r57vnn5zsn0ee
			"stars1zmaulflshft579x37acrdad72r57vnn5zsn0ee",
		}
		app.GlobalFeeKeeper.SetParams(ctx, globalfeeParams)
		return migrations, nil
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{crontypes.ModuleName, globalfeetypes.ModuleName},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
