package upgrades

import (
	"fmt"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/public-awesome/stargaze/v18/app/keepers"
)

type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// CreateUpgradeHandler defines the function that creates an upgrade handler
	CreateUpgradeHandler func(*module.Manager, module.Configurator, keepers.StargazeKeepers) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades storetypes.StoreUpgrades
}

// ChainIDAny is the wildcard ChainID that lets a Fork fire on any chain.
// Forks must set ChainID explicitly; this constant exists for the rare case
// (e.g. an e2e test or local devnet) where matching every chain is desired.
const ChainIDAny = "*"

// Fork defines a hard fork that runs custom state transition code at a specific
// block height, without going through the x/upgrade module. Useful when
// governance is unavailable. Any logic in BeginForkLogic must be deterministic
// and identical across all validators.
//
// BeginForkLogic is invoked inside a cache context: if it returns an error or
// panics, state changes are discarded and the chain keeps running.
type Fork struct {
	// Fork name, e.g. `v19`.
	UpgradeName string

	// ChainID restricts the fork to a specific chain (e.g. "stargaze-1" for
	// mainnet, "elgafar-1" for testnet). Required; empty values cause the
	// binary to fail to start. Use ChainIDAny ("*") to match any chain.
	ChainID string

	// Block height at which BeginForkLogic runs.
	UpgradeHeight int64

	// State transition code run once at UpgradeHeight in BeginBlock.
	BeginForkLogic func(ctx sdk.Context, keepers keepers.StargazeKeepers) error
}

// Validate ensures the Fork is wired up correctly. Called at app init time
// against every registered fork so misconfiguration prevents the binary from
// starting rather than silently misfiring (or not firing) on chain.
func (f Fork) Validate() error {
	if f.UpgradeName == "" {
		return fmt.Errorf("fork: UpgradeName is required")
	}
	if f.ChainID == "" {
		return fmt.Errorf("fork %q: ChainID is required (use %q for any chain)", f.UpgradeName, ChainIDAny)
	}
	if f.BeginForkLogic == nil {
		return fmt.Errorf("fork %q: BeginForkLogic is required", f.UpgradeName)
	}
	return nil
}
