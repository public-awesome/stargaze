package v14

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/public-awesome/stargaze/v14/app/keepers"
	"github.com/public-awesome/stargaze/v14/app/upgrades"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
)

const UpgradeName = "v14-slinky"

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	CreateUpgradeHandler: func(mm *module.Manager, cfg module.Configurator, keepers keepers.StargazeKeepers) upgradetypes.UpgradeHandler {
		return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			migrations, err := mm.RunMigrations(ctx, cfg, fromVM)
			if err != nil {
				return nil, err
			}

			// upgrade consensus params to enable vote extensions
			consensusParams, err := keepers.ConsensusParamsKeeper.Params(ctx, nil)
			if err != nil {
				return nil, err
			}
			wctx := sdk.UnwrapSDKContext(ctx)
			consensusParams.Params.Abci = &cmttypes.ABCIParams{
				VoteExtensionsEnableHeight: wctx.BlockHeight() + int64(10),
			}

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
			marketmaptypes.ModuleName,
			oracletypes.ModuleName,
		},
	},
}
