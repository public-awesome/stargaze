package app

import (
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// next upgrade name
const upgradeName = "v2"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *App) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(upgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		const (
			HumanCoinUnit = "stars"
			BaseCoinUnit  = "ustars"
			StarsExponent = 6
		)
		denomMetadata := banktypes.Metadata{
			Description: "The native token of Stargaze",
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    BaseCoinUnit,
					Exponent: 0,
					Aliases:  []string{"microstars"},
				},
				{
					Denom:    HumanCoinUnit,
					Exponent: StarsExponent,
					Aliases:  nil,
				},
			},
			Base:    BaseCoinUnit,
			Display: HumanCoinUnit,
			Name:    "Stargaze STARS",
			Symbol:  "STARS",
		}

		app.BankKeeper.SetDenomMetaData(ctx, denomMetadata)
		return app.mm.RunMigrations(ctx, cfg, vm)
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{authz.ModuleName, wasm.ModuleName},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
