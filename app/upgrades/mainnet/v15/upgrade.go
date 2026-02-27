package v15

import (
	"context"
	"time"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	wasmlctypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	"github.com/public-awesome/stargaze/v18/app/keepers"
	"github.com/public-awesome/stargaze/v18/app/upgrades"
)

// next upgrade name
const UpgradeName = "v15"

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

			// Adding the wasm light client to allowed clients
			params := keepers.IBCKeeper.ClientKeeper.GetParams(wctx)
			params.AllowedClients = append(params.AllowedClients, wasmlctypes.Wasm)
			keepers.IBCKeeper.ClientKeeper.SetParams(wctx, params)

			// upgrade consensus params to enable vote extensions
			consensusParams, err := keepers.ConsensusParamsKeeper.Params(ctx, nil)
			if err != nil {
				return nil, err
			}

			consensusParams.Params.Abci = &cmttypes.ABCIParams{}

			blockParams := consensusParams.Params.Block
			blockParams.MaxBytes = 4_190_208 // 4MB
			blockParams.MaxGas = 225_000_000 // 225M
			_, err = keepers.ConsensusParamsKeeper.UpdateParams(ctx, &consensustypes.MsgUpdateParams{
				Authority: keepers.ConsensusParamsKeeper.GetAuthority(),
				Block:     blockParams,
				Evidence:  consensusParams.Params.Evidence,
				Validator: consensusParams.Params.Validator,
				Abci:      consensusParams.Params.Abci,
			})
			if err != nil {
				return nil, err
			}

			// Increase the tx size cost per byte to 15
			accountParams := keepers.AccountKeeper.GetParams(ctx)
			accountParams.TxSizeCostPerByte = 15
			err = keepers.AccountKeeper.Params.Set(ctx, accountParams)
			if err != nil {
				return nil, err
			}

			wctx.Logger().Info("upgrade completed", "duration_ms", time.Since(startTime).Milliseconds())
			return migrations, nil
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			wasmlctypes.ModuleName,
		},
	},
}
