package v14

import (
	"context"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/public-awesome/stargaze/v14/app/keepers"
	"github.com/public-awesome/stargaze/v14/app/upgrades"
	"github.com/public-awesome/stargaze/v14/internal/oracle/markets"
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

			// set marketmap params
			mmParams := marketmaptypes.DefaultParams()
			// TODO: stargaze foundation or another address?
			mmParams.Admin = authtypes.NewModuleAddress(govtypes.ModuleName).String()
			mmParams.MarketAuthorities = []string{"stars1ua63s43u2p4v38pxhcxmps0tj2gudyw2828x65"}
			if err := mmParams.ValidateBasic(); err != nil {
				return nil, err
			}

			if err := keepers.MarketMapKeeper.SetParams(wctx, mmParams); err != nil {
				return nil, err
			}

			// add markets
			m, err := markets.Slice()
			if err != nil {
				return nil, err
			}

			// iterates over slice and not map
			for _, market := range m {
				// create market
				err = keepers.MarketMapKeeper.CreateMarket(wctx, market)
				if err != nil {
					return nil, err
				}

				// invoke hooks
				err = keepers.MarketMapKeeper.Hooks().AfterMarketCreated(wctx, market)
				if err != nil {
					return nil, err
				}
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
