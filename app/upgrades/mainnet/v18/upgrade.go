package v18

import (
	"context"
	"time"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/public-awesome/stargaze/v17/app/keepers"
	"github.com/public-awesome/stargaze/v17/app/upgrades"
	pausertypes "github.com/public-awesome/stargaze/v17/x/pauser/types"
)

// next upgrade name
const UpgradeName = "v18"

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	CreateUpgradeHandler: func(mm *module.Manager, cfg module.Configurator, keepers keepers.StargazeKeepers) upgradetypes.UpgradeHandler {
		return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			startTime := time.Now()
			wctx := sdk.UnwrapSDKContext(ctx)
			wctx.Logger().Info("upgrade started", "upgrade_name", UpgradeName)
			migrations, err := mm.RunMigrations(ctx, cfg, fromVM)
			if err != nil {
				return nil, err
			}

			// Set initial pauser module params with authorized address
			pauserParams := pausertypes.Params{
				PrivilegedAddresses: []string{
					"stars1s8qx0zvz8yd6e4x0mqmqf7fr9vvfn622wtp3g3",
				},
			}
			if err := keepers.PauserKeeper.SetParams(wctx, pauserParams); err != nil {
				return nil, err
			}

			wctx.Logger().Info("upgrade completed", "duration_ms", time.Since(startTime).Milliseconds())
			return migrations, nil
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{"pauser"},
	},
}
