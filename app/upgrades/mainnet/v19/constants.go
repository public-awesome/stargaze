package v19

import (
	"github.com/public-awesome/stargaze/v18/app/upgrades"
)

const (
	// UpgradeName is the name of the v19 hard fork.
	UpgradeName = "v19"

	// ChainID restricts this fork to the Stargaze mainnet chain. The fork will
	// not fire on testnet, local chains, or e2e tests, even if they happen to
	// reach UpgradeHeight.
	ChainID = "stargaze-1"

	// UpgradeHeight is the block height at which the v19 fork logic runs.
	// Targets ~14:00 UTC on 2026-05-25 assuming a
	// recent block time of ~2.43s
	UpgradeHeight = int64(33_362_003)
)

var Fork = upgrades.Fork{
	UpgradeName:    UpgradeName,
	ChainID:        ChainID,
	UpgradeHeight:  UpgradeHeight,
	BeginForkLogic: RunForkLogic,
}
