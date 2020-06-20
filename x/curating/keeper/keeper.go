package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stakebird/x/stake/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the x/stake store
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.Marshaler
	stakingKeeper types.StakingKeeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	paramstore    paramtypes.Subspace
}

// NewKeeper creates a x/stake keeper
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey, accountKeeper types.AccountKeeper,
	stakingKeeper types.StakingKeeper, bankKeeper types.BankKeeper, ps paramtypes.Subspace) Keeper {

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(ParamKeyTable())
	}

	// ensure reward pool module account is set
	if addr := accountKeeper.GetModuleAddress(types.RewardPoolName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.RewardPoolName))
	}

	keeper := Keeper{
		storeKey:      key,
		cdc:           cdc,
		stakingKeeper: stakingKeeper,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		paramstore:    ps,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}
