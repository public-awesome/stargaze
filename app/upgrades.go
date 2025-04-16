package app

import (
	"fmt"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgrades "github.com/public-awesome/stargaze/v17/app/upgrades"
	mainnetupgradesv15 "github.com/public-awesome/stargaze/v17/app/upgrades/mainnet/v15"
	mainnetupgradesv16 "github.com/public-awesome/stargaze/v17/app/upgrades/mainnet/v16"
	testnetupgradesv15b2 "github.com/public-awesome/stargaze/v17/app/upgrades/testnet/v15b2"
	testnetupgradesv15b3 "github.com/public-awesome/stargaze/v17/app/upgrades/testnet/v15b3"
)

var Upgrades = []upgrades.Upgrade{
	// mainnet upgrades
	mainnetupgradesv15.Upgrade,
	mainnetupgradesv16.Upgrade,
	// testnet upgrades
	testnetupgradesv15b2.Upgrade,
	testnetupgradesv15b3.Upgrade,
}

func (app App) RegisterUpgradeHandlers(configurator module.Configurator) {
	for _, u := range Upgrades {
		app.Keepers.UpgradeKeeper.SetUpgradeHandler(
			u.UpgradeName,
			u.CreateUpgradeHandler(app.ModuleManager, configurator, app.Keepers),
		)
	}

	upgradeInfo, err := app.Keepers.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.Keepers.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, u := range Upgrades {
		u := u
		if upgradeInfo.Name == u.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &u.StoreUpgrades))
		}
	}
}
