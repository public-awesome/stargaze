package keepers

import (
	packetforwardkeeper "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/keeper"
	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8"
	ibchookskeeper "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/keeper"
	ibcwasmkeeper "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/keeper"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	allocmodulekeeper "github.com/public-awesome/stargaze/v14/x/alloc/keeper"
	cronmodulekeeper "github.com/public-awesome/stargaze/v14/x/cron/keeper"
	globalfeemodulekeeper "github.com/public-awesome/stargaze/v14/x/globalfee/keeper"
	mintkeeper "github.com/public-awesome/stargaze/v14/x/mint/keeper"
	tokenfactorykeeper "github.com/public-awesome/stargaze/v14/x/tokenfactory/keeper"
)

type StargazeKeepers struct {
	// Cosmos-SDK Keepers

	// IBC Keepers
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCWasmKeeper       ibcwasmkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper

	// IBC Apps/Middlewares Keepers
	IBCHooksKeeper      ibchookskeeper.Keeper
	Ics20WasmHooks      *ibchooks.WasmHooks
	HooksICS4Wrapper    ibchooks.ICS4Middleware
	PacketForwardKeeper *packetforwardkeeper.Keeper

	// Stargaze Keepers
	AllocKeeper        allocmodulekeeper.Keeper
	CronKeeper         cronmodulekeeper.Keeper
	GlobalFeeKeeper    globalfeemodulekeeper.Keeper
	MintKeeper         mintkeeper.Keeper
	TokenFactoryKeeper tokenfactorykeeper.Keeper
}
