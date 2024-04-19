package v14

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	wasmlctypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	authoritytypes "github.com/public-awesome/stargaze/v13/x/authority/types"
	"github.com/public-awesome/stargaze/v14/app/keepers"
	"github.com/public-awesome/stargaze/v14/app/upgrades"
)

// next upgrade name
const UpgradeName = "v14"

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	CreateUpgradeHandler: func(mm *module.Manager, cfg module.Configurator, keepers keepers.StargazeKeepers) upgradetypes.UpgradeHandler {
		return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			migrations, err := mm.RunMigrations(ctx, cfg, fromVM)
			if err != nil {
				return nil, err
			}

			wctx := sdk.UnwrapSDKContext(ctx)
			// Adding the wasm light client to allowed clients
			params := keepers.IBCKeeper.ClientKeeper.GetParams(wctx)
			params.AllowedClients = append(params.AllowedClients, wasmlctypes.Wasm)
			keepers.IBCKeeper.ClientKeeper.SetParams(wctx, params)

			app.AuthorityKeeper.SetParams(sdk.UnwrapSDKContext(ctx), authoritytypes.DefaultParams())
			return migrations, nil
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			wasmlctypes.ModuleName,
			authoritytypes.ModuleName,
		},
	},
}
