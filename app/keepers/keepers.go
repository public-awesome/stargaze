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

	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"

	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"

	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	marketmapkeeper "github.com/skip-mev/slinky/x/marketmap/keeper"
	oraclekeeper "github.com/skip-mev/slinky/x/oracle/keeper"
)

type StargazeKeepers struct {
	// Cosmos-SDK Keepers
	AccountKeeper         authkeeper.AccountKeeper
	AuthzKeeper           authzkeeper.Keeper
	BankKeeper            bankkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	GovKeeper             govkeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper

	// Wasm Keepers
	WasmKeeper     wasmkeeper.Keeper
	ContractKeeper *wasmkeeper.PermissionedKeeper

	// IBC Keepers
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	IBCWasmKeeper       ibcwasmkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper

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

	// Slinky
	OracleKeeper    *oraclekeeper.Keeper
	MarketMapKeeper *marketmapkeeper.Keeper
}
