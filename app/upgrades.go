package app

import (
	"fmt"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	alloctypes "github.com/public-awesome/stargaze/v13/x/alloc/types"
	globalfeetypes "github.com/public-awesome/stargaze/v13/x/globalfee/types"
	minttypes "github.com/public-awesome/stargaze/v13/x/mint/types"
	tokenfactorytypes "github.com/public-awesome/stargaze/v13/x/tokenfactory/types"
)

// next upgrade name
const upgradeName = "v13-rc6-testnet"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *App) RegisterUpgradeHandlers(cfg module.Configurator) {
	// Set param key table for params module migration
	for _, subspace := range app.ParamsKeeper.GetSubspaces() {
		subspace := subspace

		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable()
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint:staticcheck
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
			// ibc types
		case ibctransfertypes.ModuleName:
			keyTable = ibctransfertypes.ParamKeyTable()
		case icahosttypes.SubModuleName:
			keyTable = icahosttypes.ParamKeyTable()
		case icacontrollertypes.SubModuleName:
			keyTable = icacontrollertypes.ParamKeyTable()
			// wasm
		case wasmtypes.ModuleName:
			keyTable = wasmtypes.ParamKeyTable() //nolint

			// stargaze modules
		case alloctypes.ModuleName:
			keyTable = alloctypes.ParamKeyTable()
		case globalfeetypes.ModuleName:
			keyTable = globalfeetypes.ParamKeyTable()
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable()
		case tokenfactorytypes.ModuleName:
			keyTable = tokenfactorytypes.ParamKeyTable()
		default:
			continue
		}

		if !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}

	baseAppLegacySS := app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())

	app.UpgradeKeeper.SetUpgradeHandler(upgradeName, func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// Migrate Tendermint consensus parameters from x/params module to a dedicated x/consensus module.
		baseapp.MigrateParams(ctx, baseAppLegacySS, &app.ConsensusParamsKeeper)

		// run migrations before modifying state
		migrations, err := app.mm.RunMigrations(ctx, cfg, fromVM)
		if err != nil {
			return nil, err
		}
		icaParams := app.ICAControllerKeeper.GetParams(ctx)
		icaParams.ControllerEnabled = true
		app.ICAControllerKeeper.SetParams(ctx, icaParams)
		return migrations, nil
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{
				icacontrollertypes.StoreKey,
			},
			Deleted: []string{},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
