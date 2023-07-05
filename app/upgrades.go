package app

import (
	"fmt"
	"time"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	allocmoduletypes "github.com/public-awesome/stargaze/v11/x/alloc/types"
	ibchooks "github.com/public-awesome/stargaze/v11/x/ibchooks/types"
	tokenfactorytypes "github.com/public-awesome/stargaze/v11/x/tokenfactory/types"
)

// next upgrade name
const upgradeName = "v11"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *App) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(upgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// run migrations before modifying state
		migrations, err := app.mm.RunMigrations(ctx, cfg, fromVM)
		if err != nil {
			return nil, err
		}

		// Following param changes reflect what was approved by prop 165 and combined in a single upgrade for Prop 1-3
		// https://www.mintscan.io/stargaze/proposals/165
		// 1- Reduce Emissions to 711k daily  this is done by adjusting mint params
		// 2- Introduce a new Supplement Amount and Redirect Funds for the next 6 months
		//    requiring future proposals to refill the module account
		// 3- Stop funding community pool and nft incentive allocation
		//    at the same time this code allows re-enabling incentive allocation
		//    through param change proposals only

		// change mint params to include the new supplement amount
		// and store it back to the keeper
		mintParams := app.MintKeeper.GetParams(ctx)
		// 259.7M to the upgrade happening on the 11th of July 2023
		mintParams.InitialAnnualProvisions = sdk.NewDec(259_700_000_000_000)
		// reset to 2023-01-01 so there is a thirdening happening on next January 1
		mintParams.StartTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		// update blocks per year using  5.9 s avg block time
		mintParams.BlocksPerYear = 5345036
		app.MintKeeper.SetParams(ctx, mintParams)
		denom := app.MintKeeper.GetParams(ctx).MintDenom

		// token factory params
		params := app.TokenFactoryKeeper.GetParams(ctx)
		params.DenomCreationFee = sdk.NewCoins(sdk.NewInt64Coin(denom, 10_000_000_000)) // 10k STARS

		app.TokenFactoryKeeper.SetParams(ctx, params)

		// set community tax to 0 since the allocation module will now take care of it
		// making an accurate allocation of the inflation
		distributionParams := app.DistrKeeper.GetParams(ctx)
		distributionParams.CommunityTax = sdk.ZeroDec()
		app.DistrKeeper.SetParams(ctx, distributionParams)

		// change alloc params to set nft incentives to 0% until incentives are live
		allocParams := app.AllocKeeper.GetParams(ctx)

		// distribution proportions
		proportions := allocParams.DistributionProportions
		proportions.NftIncentives = sdk.ZeroDec()            // nft incentives to 0%
		proportions.CommunityPool = sdk.NewDecWithPrec(5, 2) // 5% community pool

		allocParams.DistributionProportions = proportions
		// supplement amount from the specific module account
		// set to 100k STARS daily ~= 6.82 STARS per block using same 5.9s avg block time
		allocParams.SupplementAmount = sdk.NewCoins(sdk.NewInt64Coin(denom, 6_828_704)) // 6.9 STARS per block
		app.AllocKeeper.SetParams(ctx, allocParams)

		// check if the account was previously created if that's the case reset it
		supplementPoolAddress := authtypes.NewModuleAddress(allocmoduletypes.SupplementPoolName)
		if app.AccountKeeper.HasAccount(ctx, supplementPoolAddress) {
			account := app.AccountKeeper.GetAccount(ctx, supplementPoolAddress)
			app.AccountKeeper.RemoveAccount(ctx, account)
		}
		// create module account
		supplmentPoolAccount := app.AccountKeeper.GetModuleAccount(ctx, allocmoduletypes.SupplementPoolName)

		fundAmount := sdk.NewInt64Coin(denom, 18_300_000_000_000) // 18M STARS
		// if there is not enough founds skip
		if app.DistrKeeper.GetFeePoolCommunityCoins(ctx).AmountOf(denom).LT(sdk.NewDecCoinFromCoin(fundAmount).Amount) {
			return migrations, nil
		}
		err = app.DistrKeeper.DistributeFromFeePool(ctx, sdk.NewCoins(fundAmount), supplmentPoolAccount.GetAddress())
		if err != nil {
			return nil, err
		}

		return migrations, nil
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{tokenfactorytypes.ModuleName, ibchooks.StoreKey},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
