package v15

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	wasmlctypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	"github.com/public-awesome/stargaze/v15/app/keepers"
	"github.com/public-awesome/stargaze/v15/app/upgrades"
)

// next upgrade name
const UpgradeName = "v15"

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

			// upgrade consensus params to enable vote extensions
			consensusParams, err := keepers.ConsensusParamsKeeper.Params(ctx, nil)
			if err != nil {
				return nil, err
			}

			consensusParams.Params.Abci = &cmttypes.ABCIParams{}

			blockParams := consensusParams.Params.Block
			blockParams.MaxBytes = 4_190_208 // 4MB
			blockParams.MaxGas = 200_000_000 // 200M
			_, err = keepers.ConsensusParamsKeeper.UpdateParams(ctx, &consensustypes.MsgUpdateParams{
				Authority: keepers.ConsensusParamsKeeper.GetAuthority(),
				Block:     consensusParams.Params.Block,
				Evidence:  consensusParams.Params.Evidence,
				Validator: consensusParams.Params.Validator,
				Abci:      consensusParams.Params.Abci,
			})
			if err != nil {
				return nil, err
			}

			return migrations, nil
		}
	},
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			wasmlctypes.ModuleName,
		},
	},
}
