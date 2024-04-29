<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [osmosis/tokenfactory/v1beta1/tokenfactory.proto](#osmosis/tokenfactory/v1beta1/tokenfactory.proto)
    - [DenomAuthorityMetadata](#osmosis.tokenfactory.v1beta1.DenomAuthorityMetadata)
    - [Params](#osmosis.tokenfactory.v1beta1.Params)
  
- [osmosis/tokenfactory/v1beta1/genesis.proto](#osmosis/tokenfactory/v1beta1/genesis.proto)
    - [GenesisDenom](#osmosis.tokenfactory.v1beta1.GenesisDenom)
    - [GenesisState](#osmosis.tokenfactory.v1beta1.GenesisState)
  
- [osmosis/tokenfactory/v1beta1/query.proto](#osmosis/tokenfactory/v1beta1/query.proto)
    - [QueryDenomAuthorityMetadataRequest](#osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataRequest)
    - [QueryDenomAuthorityMetadataResponse](#osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataResponse)
    - [QueryDenomsFromCreatorRequest](#osmosis.tokenfactory.v1beta1.QueryDenomsFromCreatorRequest)
    - [QueryDenomsFromCreatorResponse](#osmosis.tokenfactory.v1beta1.QueryDenomsFromCreatorResponse)
    - [QueryParamsRequest](#osmosis.tokenfactory.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#osmosis.tokenfactory.v1beta1.QueryParamsResponse)
  
    - [Query](#osmosis.tokenfactory.v1beta1.Query)
  
- [osmosis/tokenfactory/v1beta1/tx.proto](#osmosis/tokenfactory/v1beta1/tx.proto)
    - [MsgBurn](#osmosis.tokenfactory.v1beta1.MsgBurn)
    - [MsgBurnResponse](#osmosis.tokenfactory.v1beta1.MsgBurnResponse)
    - [MsgChangeAdmin](#osmosis.tokenfactory.v1beta1.MsgChangeAdmin)
    - [MsgChangeAdminResponse](#osmosis.tokenfactory.v1beta1.MsgChangeAdminResponse)
    - [MsgCreateDenom](#osmosis.tokenfactory.v1beta1.MsgCreateDenom)
    - [MsgCreateDenomResponse](#osmosis.tokenfactory.v1beta1.MsgCreateDenomResponse)
    - [MsgMint](#osmosis.tokenfactory.v1beta1.MsgMint)
    - [MsgMintResponse](#osmosis.tokenfactory.v1beta1.MsgMintResponse)
    - [MsgSetDenomMetadata](#osmosis.tokenfactory.v1beta1.MsgSetDenomMetadata)
    - [MsgSetDenomMetadataResponse](#osmosis.tokenfactory.v1beta1.MsgSetDenomMetadataResponse)
    - [MsgUpdateParams](#osmosis.tokenfactory.v1beta1.MsgUpdateParams)
    - [MsgUpdateParamsResponse](#osmosis.tokenfactory.v1beta1.MsgUpdateParamsResponse)
  
    - [Msg](#osmosis.tokenfactory.v1beta1.Msg)
  
- [publicawesome/stargaze/alloc/v1beta1/params.proto](#publicawesome/stargaze/alloc/v1beta1/params.proto)
    - [DistributionProportions](#publicawesome.stargaze.alloc.v1beta1.DistributionProportions)
    - [Params](#publicawesome.stargaze.alloc.v1beta1.Params)
    - [WeightedAddress](#publicawesome.stargaze.alloc.v1beta1.WeightedAddress)
  
- [publicawesome/stargaze/alloc/v1beta1/genesis.proto](#publicawesome/stargaze/alloc/v1beta1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.alloc.v1beta1.GenesisState)
  
- [publicawesome/stargaze/alloc/v1beta1/query.proto](#publicawesome/stargaze/alloc/v1beta1/query.proto)
    - [QueryParamsRequest](#publicawesome.stargaze.alloc.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#publicawesome.stargaze.alloc.v1beta1.QueryParamsResponse)
  
    - [Query](#publicawesome.stargaze.alloc.v1beta1.Query)
  
- [publicawesome/stargaze/alloc/v1beta1/tx.proto](#publicawesome/stargaze/alloc/v1beta1/tx.proto)
    - [MsgCreateVestingAccount](#publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccount)
    - [MsgCreateVestingAccountResponse](#publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccountResponse)
    - [MsgFundFairburnPool](#publicawesome.stargaze.alloc.v1beta1.MsgFundFairburnPool)
    - [MsgFundFairburnPoolResponse](#publicawesome.stargaze.alloc.v1beta1.MsgFundFairburnPoolResponse)
    - [MsgUpdateParams](#publicawesome.stargaze.alloc.v1beta1.MsgUpdateParams)
    - [MsgUpdateParamsResponse](#publicawesome.stargaze.alloc.v1beta1.MsgUpdateParamsResponse)
  
    - [Msg](#publicawesome.stargaze.alloc.v1beta1.Msg)
  
- [publicawesome/stargaze/cron/v1/cron.proto](#publicawesome/stargaze/cron/v1/cron.proto)
    - [Params](#publicawesome.stargaze.cron.v1.Params)
  
- [publicawesome/stargaze/cron/v1/genesis.proto](#publicawesome/stargaze/cron/v1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.cron.v1.GenesisState)
  
- [publicawesome/stargaze/cron/v1/proposal.proto](#publicawesome/stargaze/cron/v1/proposal.proto)
    - [DemotePrivilegedContractProposal](#publicawesome.stargaze.cron.v1.DemotePrivilegedContractProposal)
    - [PromoteToPrivilegedContractProposal](#publicawesome.stargaze.cron.v1.PromoteToPrivilegedContractProposal)
  
- [publicawesome/stargaze/cron/v1/query.proto](#publicawesome/stargaze/cron/v1/query.proto)
    - [QueryListPrivilegedRequest](#publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest)
    - [QueryListPrivilegedResponse](#publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse)
    - [QueryParamsRequest](#publicawesome.stargaze.cron.v1.QueryParamsRequest)
    - [QueryParamsResponse](#publicawesome.stargaze.cron.v1.QueryParamsResponse)
  
    - [Query](#publicawesome.stargaze.cron.v1.Query)
  
- [publicawesome/stargaze/cron/v1/tx.proto](#publicawesome/stargaze/cron/v1/tx.proto)
    - [MsgDemoteFromPrivilegedContract](#publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContract)
    - [MsgDemoteFromPrivilegedContractResponse](#publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContractResponse)
    - [MsgPromoteToPrivilegedContract](#publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContract)
    - [MsgPromoteToPrivilegedContractResponse](#publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContractResponse)
    - [MsgUpdateParams](#publicawesome.stargaze.cron.v1.MsgUpdateParams)
    - [MsgUpdateParamsResponse](#publicawesome.stargaze.cron.v1.MsgUpdateParamsResponse)
  
    - [Msg](#publicawesome.stargaze.cron.v1.Msg)
  
- [publicawesome/stargaze/globalfee/v1/globalfee.proto](#publicawesome/stargaze/globalfee/v1/globalfee.proto)
    - [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization)
    - [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization)
    - [Params](#publicawesome.stargaze.globalfee.v1.Params)
  
- [publicawesome/stargaze/globalfee/v1/genesis.proto](#publicawesome/stargaze/globalfee/v1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.globalfee.v1.GenesisState)
  
- [publicawesome/stargaze/globalfee/v1/proposal.proto](#publicawesome/stargaze/globalfee/v1/proposal.proto)
    - [RemoveCodeAuthorizationProposal](#publicawesome.stargaze.globalfee.v1.RemoveCodeAuthorizationProposal)
    - [RemoveContractAuthorizationProposal](#publicawesome.stargaze.globalfee.v1.RemoveContractAuthorizationProposal)
    - [SetCodeAuthorizationProposal](#publicawesome.stargaze.globalfee.v1.SetCodeAuthorizationProposal)
    - [SetContractAuthorizationProposal](#publicawesome.stargaze.globalfee.v1.SetContractAuthorizationProposal)
  
- [publicawesome/stargaze/globalfee/v1/query.proto](#publicawesome/stargaze/globalfee/v1/query.proto)
    - [QueryAuthorizationsRequest](#publicawesome.stargaze.globalfee.v1.QueryAuthorizationsRequest)
    - [QueryAuthorizationsResponse](#publicawesome.stargaze.globalfee.v1.QueryAuthorizationsResponse)
    - [QueryCodeAuthorizationRequest](#publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationRequest)
    - [QueryCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationResponse)
    - [QueryContractAuthorizationRequest](#publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationRequest)
    - [QueryContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationResponse)
    - [QueryParamsRequest](#publicawesome.stargaze.globalfee.v1.QueryParamsRequest)
    - [QueryParamsResponse](#publicawesome.stargaze.globalfee.v1.QueryParamsResponse)
  
    - [Query](#publicawesome.stargaze.globalfee.v1.Query)
  
- [publicawesome/stargaze/globalfee/v1/tx.proto](#publicawesome/stargaze/globalfee/v1/tx.proto)
    - [MsgRemoveCodeAuthorization](#publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorization)
    - [MsgRemoveCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorizationResponse)
    - [MsgRemoveContractAuthorization](#publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorization)
    - [MsgRemoveContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorizationResponse)
    - [MsgSetCodeAuthorization](#publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorization)
    - [MsgSetCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorizationResponse)
    - [MsgSetContractAuthorization](#publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorization)
    - [MsgSetContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorizationResponse)
    - [MsgUpdateParams](#publicawesome.stargaze.globalfee.v1.MsgUpdateParams)
    - [MsgUpdateParamsResponse](#publicawesome.stargaze.globalfee.v1.MsgUpdateParamsResponse)
  
    - [Msg](#publicawesome.stargaze.globalfee.v1.Msg)
  
- [publicawesome/stargaze/mint/v1beta1/mint.proto](#publicawesome/stargaze/mint/v1beta1/mint.proto)
    - [Minter](#publicawesome.stargaze.mint.v1beta1.Minter)
    - [Params](#publicawesome.stargaze.mint.v1beta1.Params)
  
- [publicawesome/stargaze/mint/v1beta1/genesis.proto](#publicawesome/stargaze/mint/v1beta1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.mint.v1beta1.GenesisState)
  
- [publicawesome/stargaze/mint/v1beta1/query.proto](#publicawesome/stargaze/mint/v1beta1/query.proto)
    - [QueryAnnualProvisionsRequest](#publicawesome.stargaze.mint.v1beta1.QueryAnnualProvisionsRequest)
    - [QueryAnnualProvisionsResponse](#publicawesome.stargaze.mint.v1beta1.QueryAnnualProvisionsResponse)
    - [QueryParamsRequest](#publicawesome.stargaze.mint.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#publicawesome.stargaze.mint.v1beta1.QueryParamsResponse)
  
    - [Query](#publicawesome.stargaze.mint.v1beta1.Query)
  
- [publicawesome/stargaze/mint/v1beta1/tx.proto](#publicawesome/stargaze/mint/v1beta1/tx.proto)
    - [MsgUpdateParams](#publicawesome.stargaze.mint.v1beta1.MsgUpdateParams)
    - [MsgUpdateParamsResponse](#publicawesome.stargaze.mint.v1beta1.MsgUpdateParamsResponse)
  
    - [Msg](#publicawesome.stargaze.mint.v1beta1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="osmosis/tokenfactory/v1beta1/tokenfactory.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## osmosis/tokenfactory/v1beta1/tokenfactory.proto



<a name="osmosis.tokenfactory.v1beta1.DenomAuthorityMetadata"></a>

### DenomAuthorityMetadata
DenomAuthorityMetadata specifies metadata for addresses that have specific
capabilities over a token factory denom. Right now there is only one Admin
permission, but is planned to be extended to the future.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | Can be empty for no admin, or a valid stargaze address |






<a name="osmosis.tokenfactory.v1beta1.Params"></a>

### Params
Params defines the parameters for the tokenfactory module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom_creation_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | DenomCreationFee defines the fee to be charged on the creation of a new denom. The fee is drawn from the MsgCreateDenom's sender account, and transferred to the community pool. |
| `denom_creation_gas_consume` | [uint64](#uint64) |  | DenomCreationGasConsume defines the gas cost for creating a new denom. This is intended as a spam deterrence mechanism.

See: https://github.com/CosmWasm/token-factory/issues/11 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="osmosis/tokenfactory/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## osmosis/tokenfactory/v1beta1/genesis.proto



<a name="osmosis.tokenfactory.v1beta1.GenesisDenom"></a>

### GenesisDenom
GenesisDenom defines a tokenfactory denom that is defined within genesis
state. The structure contains DenomAuthorityMetadata which defines the
denom's admin.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `authority_metadata` | [DenomAuthorityMetadata](#osmosis.tokenfactory.v1beta1.DenomAuthorityMetadata) |  |  |






<a name="osmosis.tokenfactory.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the tokenfactory module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#osmosis.tokenfactory.v1beta1.Params) |  | params defines the paramaters of the module. |
| `factory_denoms` | [GenesisDenom](#osmosis.tokenfactory.v1beta1.GenesisDenom) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="osmosis/tokenfactory/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## osmosis/tokenfactory/v1beta1/query.proto



<a name="osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataRequest"></a>

### QueryDenomAuthorityMetadataRequest
QueryDenomAuthorityMetadataRequest defines the request structure for the
DenomAuthorityMetadata gRPC query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataResponse"></a>

### QueryDenomAuthorityMetadataResponse
QueryDenomAuthorityMetadataResponse defines the response structure for the
DenomAuthorityMetadata gRPC query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority_metadata` | [DenomAuthorityMetadata](#osmosis.tokenfactory.v1beta1.DenomAuthorityMetadata) |  |  |






<a name="osmosis.tokenfactory.v1beta1.QueryDenomsFromCreatorRequest"></a>

### QueryDenomsFromCreatorRequest
QueryDenomsFromCreatorRequest defines the request structure for the
DenomsFromCreator gRPC query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |






<a name="osmosis.tokenfactory.v1beta1.QueryDenomsFromCreatorResponse"></a>

### QueryDenomsFromCreatorResponse
QueryDenomsFromCreatorRequest defines the response structure for the
DenomsFromCreator gRPC query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denoms` | [string](#string) | repeated |  |






<a name="osmosis.tokenfactory.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="osmosis.tokenfactory.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#osmosis.tokenfactory.v1beta1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="osmosis.tokenfactory.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#osmosis.tokenfactory.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#osmosis.tokenfactory.v1beta1.QueryParamsResponse) | Params defines a gRPC query method that returns the tokenfactory module's parameters. | GET|/stargaze/tokenfactory/v1/params|
| `DenomAuthorityMetadata` | [QueryDenomAuthorityMetadataRequest](#osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataRequest) | [QueryDenomAuthorityMetadataResponse](#osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataResponse) | DenomAuthorityMetadata defines a gRPC query method for fetching DenomAuthorityMetadata for a particular denom. | GET|/stargaze/tokenfactory/v1/denoms/{denom}/authority_metadata|
| `DenomsFromCreator` | [QueryDenomsFromCreatorRequest](#osmosis.tokenfactory.v1beta1.QueryDenomsFromCreatorRequest) | [QueryDenomsFromCreatorResponse](#osmosis.tokenfactory.v1beta1.QueryDenomsFromCreatorResponse) | DenomsFromCreator defines a gRPC query method for fetching all denominations created by a specific admin/creator. | GET|/stargaze/tokenfactory/v1/denoms_from_creator/{creator}|

 <!-- end services -->



<a name="osmosis/tokenfactory/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## osmosis/tokenfactory/v1beta1/tx.proto



<a name="osmosis.tokenfactory.v1beta1.MsgBurn"></a>

### MsgBurn
MsgBurn is the sdk.Msg type for allowing an admin account to burn
a token.  For now, we only support burning from the sender account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `burnFromAddress` | [string](#string) |  |  |






<a name="osmosis.tokenfactory.v1beta1.MsgBurnResponse"></a>

### MsgBurnResponse
MsgBurnResponse response from executing MsgBurn.






<a name="osmosis.tokenfactory.v1beta1.MsgChangeAdmin"></a>

### MsgChangeAdmin
MsgChangeAdmin is the sdk.Msg type for allowing an admin account to reassign
adminship of a denom to a new account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `new_admin` | [string](#string) |  |  |






<a name="osmosis.tokenfactory.v1beta1.MsgChangeAdminResponse"></a>

### MsgChangeAdminResponse
MsgChangeAdminResponse defines the response structure for an executed
MsgChangeAdmin message.






<a name="osmosis.tokenfactory.v1beta1.MsgCreateDenom"></a>

### MsgCreateDenom
MsgCreateDenom defines the message structure for the CreateDenom gRPC service
method. It allows an account to create a new denom. It requires a sender
address and a sub denomination. The (sender_address, sub_denomination) tuple
must be unique and cannot be re-used.

The resulting denom created is defined as
<factory/{creatorAddress}/{subdenom}>. The resulting denom's admin is
originally set to be the creator, but this can be changed later. The token
denom does not indicate the current admin.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `subdenom` | [string](#string) |  | subdenom can be up to 44 "alphanumeric" characters long. |






<a name="osmosis.tokenfactory.v1beta1.MsgCreateDenomResponse"></a>

### MsgCreateDenomResponse
MsgCreateDenomResponse is the return value of MsgCreateDenom
It returns the full string of the newly created denom


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `new_token_denom` | [string](#string) |  |  |






<a name="osmosis.tokenfactory.v1beta1.MsgMint"></a>

### MsgMint
MsgMint is the sdk.Msg type for allowing an admin account to mint
more of a token.  For now, we only support minting to the sender account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `mintToAddress` | [string](#string) |  |  |






<a name="osmosis.tokenfactory.v1beta1.MsgMintResponse"></a>

### MsgMintResponse
MsgMintResponse response from executing MsgMint.






<a name="osmosis.tokenfactory.v1beta1.MsgSetDenomMetadata"></a>

### MsgSetDenomMetadata
MsgSetDenomMetadata is the sdk.Msg type for allowing an admin account to set
the denom's bank metadata


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `metadata` | [cosmos.bank.v1beta1.Metadata](#cosmos.bank.v1beta1.Metadata) |  |  |






<a name="osmosis.tokenfactory.v1beta1.MsgSetDenomMetadataResponse"></a>

### MsgSetDenomMetadataResponse
MsgSetDenomMetadataResponse defines the response structure for an executed
MsgSetDenomMetadata message.






<a name="osmosis.tokenfactory.v1beta1.MsgUpdateParams"></a>

### MsgUpdateParams
MsgUpdateParams is the request type for updating module's params.

Since: v14


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `params` | [Params](#osmosis.tokenfactory.v1beta1.Params) |  | NOTE: All parameters must be supplied. |






<a name="osmosis.tokenfactory.v1beta1.MsgUpdateParamsResponse"></a>

### MsgUpdateParamsResponse
MsgUpdateParamsResponse is the response type for executing
an update.
Since: v14





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="osmosis.tokenfactory.v1beta1.Msg"></a>

### Msg
Msg defines the tokefactory module's gRPC message service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateDenom` | [MsgCreateDenom](#osmosis.tokenfactory.v1beta1.MsgCreateDenom) | [MsgCreateDenomResponse](#osmosis.tokenfactory.v1beta1.MsgCreateDenomResponse) | CreateDenom | |
| `Mint` | [MsgMint](#osmosis.tokenfactory.v1beta1.MsgMint) | [MsgMintResponse](#osmosis.tokenfactory.v1beta1.MsgMintResponse) | Mint | |
| `Burn` | [MsgBurn](#osmosis.tokenfactory.v1beta1.MsgBurn) | [MsgBurnResponse](#osmosis.tokenfactory.v1beta1.MsgBurnResponse) | Burn | |
| `ChangeAdmin` | [MsgChangeAdmin](#osmosis.tokenfactory.v1beta1.MsgChangeAdmin) | [MsgChangeAdminResponse](#osmosis.tokenfactory.v1beta1.MsgChangeAdminResponse) | ChangeAdmin | |
| `SetDenomMetadata` | [MsgSetDenomMetadata](#osmosis.tokenfactory.v1beta1.MsgSetDenomMetadata) | [MsgSetDenomMetadataResponse](#osmosis.tokenfactory.v1beta1.MsgSetDenomMetadataResponse) | SetDenomMetadata | |
| `UpdateParams` | [MsgUpdateParams](#osmosis.tokenfactory.v1beta1.MsgUpdateParams) | [MsgUpdateParamsResponse](#osmosis.tokenfactory.v1beta1.MsgUpdateParamsResponse) | UpdateParams updates the tokenfactory module's parameters. | |

 <!-- end services -->



<a name="publicawesome/stargaze/alloc/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/alloc/v1beta1/params.proto



<a name="publicawesome.stargaze.alloc.v1beta1.DistributionProportions"></a>

### DistributionProportions
DistributionProportions defines the proportion that each bucket  receives.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_incentives` | [string](#string) |  |  |
| `developer_rewards` | [string](#string) |  |  |
| `community_pool` | [string](#string) |  |  |






<a name="publicawesome.stargaze.alloc.v1beta1.Params"></a>

### Params
Params defines the parameters for the alloc module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `distribution_proportions` | [DistributionProportions](#publicawesome.stargaze.alloc.v1beta1.DistributionProportions) |  | distribution_proportions defines the proportion of the minted denom |
| `weighted_developer_rewards_receivers` | [WeightedAddress](#publicawesome.stargaze.alloc.v1beta1.WeightedAddress) | repeated | addresses to receive developer rewards |
| `weighted_incentives_rewards_receivers` | [WeightedAddress](#publicawesome.stargaze.alloc.v1beta1.WeightedAddress) | repeated | addresses to receive incentive rewards |
| `supplement_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | SupplementAmount is the amount to be supplemented from the pool on top of newly minted coins. |






<a name="publicawesome.stargaze.alloc.v1beta1.WeightedAddress"></a>

### WeightedAddress
WeightedAddress defines an address with a weight.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/alloc/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/alloc/v1beta1/genesis.proto



<a name="publicawesome.stargaze.alloc.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the alloc module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.alloc.v1beta1.Params) |  | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/alloc/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/alloc/v1beta1/query.proto



<a name="publicawesome.stargaze.alloc.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="publicawesome.stargaze.alloc.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.alloc.v1beta1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.alloc.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#publicawesome.stargaze.alloc.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#publicawesome.stargaze.alloc.v1beta1.QueryParamsResponse) | this line is used by starport scaffolding # 2 | GET|/stargaze/alloc/v1beta1/params|

 <!-- end services -->



<a name="publicawesome/stargaze/alloc/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/alloc/v1beta1/tx.proto



<a name="publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccount"></a>

### MsgCreateVestingAccount
Deprecated: Cosmos SDK's CreateVestingAccount now supports start time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `to_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `start_time` | [int64](#int64) |  |  |
| `end_time` | [int64](#int64) |  |  |
| `delayed` | [bool](#bool) |  |  |






<a name="publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccountResponse"></a>

### MsgCreateVestingAccountResponse
MsgCreateVestingAccountResponse defines the Msg/CreateVestingAccount response
type.






<a name="publicawesome.stargaze.alloc.v1beta1.MsgFundFairburnPool"></a>

### MsgFundFairburnPool
MsgFundFairburnPool allows an account to directly
fund the fee collector pool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="publicawesome.stargaze.alloc.v1beta1.MsgFundFairburnPoolResponse"></a>

### MsgFundFairburnPoolResponse
MsgFundFairburnPoolResponse defines the Msg/MsgFundFairburnPool response
type.






<a name="publicawesome.stargaze.alloc.v1beta1.MsgUpdateParams"></a>

### MsgUpdateParams
MsgUpdateParams is the request type for updating module's params.

Since: v14


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `params` | [Params](#publicawesome.stargaze.alloc.v1beta1.Params) |  | NOTE: All parameters must be supplied. |






<a name="publicawesome.stargaze.alloc.v1beta1.MsgUpdateParamsResponse"></a>

### MsgUpdateParamsResponse
MsgUpdateParamsResponse is the response type for executing
an update.
Since: v14





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.alloc.v1beta1.Msg"></a>

### Msg
Msg defines the alloc Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateVestingAccount` | [MsgCreateVestingAccount](#publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccount) | [MsgCreateVestingAccountResponse](#publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccountResponse) | CreateVestingAccount defines a method that enables creating a vesting account. | |
| `FundFairburnPool` | [MsgFundFairburnPool](#publicawesome.stargaze.alloc.v1beta1.MsgFundFairburnPool) | [MsgFundFairburnPoolResponse](#publicawesome.stargaze.alloc.v1beta1.MsgFundFairburnPoolResponse) | FundFairburnPool defines a method to allow an account to directly fund the fee collector module account. | |
| `UpdateParams` | [MsgUpdateParams](#publicawesome.stargaze.alloc.v1beta1.MsgUpdateParams) | [MsgUpdateParamsResponse](#publicawesome.stargaze.alloc.v1beta1.MsgUpdateParamsResponse) | UpdateParams updates the alloc module's parameters. | |

 <!-- end services -->



<a name="publicawesome/stargaze/cron/v1/cron.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/cron/v1/cron.proto



<a name="publicawesome.stargaze.cron.v1.Params"></a>

### Params
Params holds parameters for the cron module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin_addresses` | [string](#string) | repeated | Addresses which act as admins of the module. They can promote and demote contracts without having to go via governance. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/cron/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/cron/v1/genesis.proto



<a name="publicawesome.stargaze.cron.v1.GenesisState"></a>

### GenesisState
GenesisState defines the cron module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `privileged_contract_addresses` | [string](#string) | repeated | List of all the contracts that have been given the privilege status via governance. They can set up hooks to abci.EndBlocker |
| `params` | [Params](#publicawesome.stargaze.cron.v1.Params) |  | Module params |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/cron/v1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/cron/v1/proposal.proto



<a name="publicawesome.stargaze.cron.v1.DemotePrivilegedContractProposal"></a>

### DemotePrivilegedContractProposal
Deprecated: Do not use. To demote a contract, a
MsgDemoteFromPrivilegedContract can be invoked from the x/gov module via a v1
governance proposal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `contract` | [string](#string) |  | Contract is the bech32 address of the smart contract |






<a name="publicawesome.stargaze.cron.v1.PromoteToPrivilegedContractProposal"></a>

### PromoteToPrivilegedContractProposal
Deprecated: Do not use. To promote a contract, a
MsgPromoteToPrivilegedContract can be invoked from the x/gov module via a v1
governance proposal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `contract` | [string](#string) |  | Contract is the bech32 address of the smart contract |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/cron/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/cron/v1/query.proto



<a name="publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest"></a>

### QueryListPrivilegedRequest
QueryListPrivilegedRequest is request type for the Query/ListPrivileged RPC
method.






<a name="publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse"></a>

### QueryListPrivilegedResponse
QueryListPrivilegedResponse is response type for the Query/ListPrivileged RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_addresses` | [string](#string) | repeated | contract_addresses holds all the smart contract addresses which have privilege status. |






<a name="publicawesome.stargaze.cron.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC
method.






<a name="publicawesome.stargaze.cron.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.cron.v1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.cron.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ListPrivileged` | [QueryListPrivilegedRequest](#publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest) | [QueryListPrivilegedResponse](#publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse) | ListPrivileged queries the contracts which have the priviledge status | GET|/stargaze/cron/v1/list-privileged|
| `Params` | [QueryParamsRequest](#publicawesome.stargaze.cron.v1.QueryParamsRequest) | [QueryParamsResponse](#publicawesome.stargaze.cron.v1.QueryParamsResponse) |  | GET|/stargaze/cron/v1/params|

 <!-- end services -->



<a name="publicawesome/stargaze/cron/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/cron/v1/tx.proto



<a name="publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContract"></a>

### MsgDemoteFromPrivilegedContract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account or any whitelisted address |
| `contract` | [string](#string) |  | Contract is the bech32 address of the smart contract |






<a name="publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContractResponse"></a>

### MsgDemoteFromPrivilegedContractResponse







<a name="publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContract"></a>

### MsgPromoteToPrivilegedContract
MsgPromoteToPrivilegedContract defines the Msg/PromoteToPrivilegedContract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account or any whitelisted address |
| `contract` | [string](#string) |  | Contract is the bech32 address of the smart contract |






<a name="publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContractResponse"></a>

### MsgPromoteToPrivilegedContractResponse







<a name="publicawesome.stargaze.cron.v1.MsgUpdateParams"></a>

### MsgUpdateParams
MsgUpdateParams updates module's params through governance proposal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `params` | [Params](#publicawesome.stargaze.cron.v1.Params) |  | NOTE: All parameters must be supplied. |






<a name="publicawesome.stargaze.cron.v1.MsgUpdateParamsResponse"></a>

### MsgUpdateParamsResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.cron.v1.Msg"></a>

### Msg
Msg defines the alloc Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PromoteToPrivilegedContract` | [MsgPromoteToPrivilegedContract](#publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContract) | [MsgPromoteToPrivilegedContractResponse](#publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContractResponse) | PromoteToPrivilegedContract promotes a contract to privileged status. | |
| `DemoteFromPrivilegedContract` | [MsgDemoteFromPrivilegedContract](#publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContract) | [MsgDemoteFromPrivilegedContractResponse](#publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContractResponse) | DemoteFromPrivilegedContract demotes a contract from privileged status. | |
| `UpdateParams` | [MsgUpdateParams](#publicawesome.stargaze.cron.v1.MsgUpdateParams) | [MsgUpdateParamsResponse](#publicawesome.stargaze.cron.v1.MsgUpdateParamsResponse) | UpdateParams updates the cron module's parameters. | |

 <!-- end services -->



<a name="publicawesome/stargaze/globalfee/v1/globalfee.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/globalfee/v1/globalfee.proto



<a name="publicawesome.stargaze.globalfee.v1.CodeAuthorization"></a>

### CodeAuthorization
Configuration for code Ids which can have zero gas operations


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  | authorized code ids |
| `methods` | [string](#string) | repeated | authorized contract operation methods |






<a name="publicawesome.stargaze.globalfee.v1.ContractAuthorization"></a>

### ContractAuthorization
Configuration for contract addresses which can have zero gas operations


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | authorized contract addresses |
| `methods` | [string](#string) | repeated | authorized contract operation methods |






<a name="publicawesome.stargaze.globalfee.v1.Params"></a>

### Params
Params holds parameters for the globalfee module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `privileged_addresses` | [string](#string) | repeated | Addresses which are whitelisted to modify the gas free operations |
| `minimum_gas_prices` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | Minimum stores the minimum gas price(s) for all TX on the chain. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/globalfee/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/globalfee/v1/genesis.proto



<a name="publicawesome.stargaze.globalfee.v1.GenesisState"></a>

### GenesisState
GenesisState defines the globalfee module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.globalfee.v1.Params) |  | Module params |
| `code_authorizations` | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization) | repeated | Authorizations configured by code id |
| `contract_authorizations` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) | repeated | Authorizations configured by contract addresses |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/globalfee/v1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/globalfee/v1/proposal.proto



<a name="publicawesome.stargaze.globalfee.v1.RemoveCodeAuthorizationProposal"></a>

### RemoveCodeAuthorizationProposal
RemoveCodeAuthorizationProposal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `code_id` | [uint64](#uint64) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.RemoveContractAuthorizationProposal"></a>

### RemoveContractAuthorizationProposal
RemoveCodeAuthorizationProposal ...


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.SetCodeAuthorizationProposal"></a>

### SetCodeAuthorizationProposal
SetCodeAuthorizationProposal ...


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `code_authorization` | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.SetContractAuthorizationProposal"></a>

### SetContractAuthorizationProposal
RemoveCodeAuthorizationProposal ...


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `contract_authorization` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/globalfee/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/globalfee/v1/query.proto



<a name="publicawesome.stargaze.globalfee.v1.QueryAuthorizationsRequest"></a>

### QueryAuthorizationsRequest







<a name="publicawesome.stargaze.globalfee.v1.QueryAuthorizationsResponse"></a>

### QueryAuthorizationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_authorizations` | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization) | repeated |  |
| `contract_authorizations` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) | repeated |  |






<a name="publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationRequest"></a>

### QueryCodeAuthorizationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationResponse"></a>

### QueryCodeAuthorizationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `methods` | [string](#string) | repeated |  |






<a name="publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationRequest"></a>

### QueryContractAuthorizationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationResponse"></a>

### QueryContractAuthorizationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `methods` | [string](#string) | repeated |  |






<a name="publicawesome.stargaze.globalfee.v1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="publicawesome.stargaze.globalfee.v1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.globalfee.v1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.globalfee.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CodeAuthorization` | [QueryCodeAuthorizationRequest](#publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationRequest) | [QueryCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationResponse) |  | GET|/stargaze/globalfee/v1/code_authorization/{code_id}|
| `ContractAuthorization` | [QueryContractAuthorizationRequest](#publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationRequest) | [QueryContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationResponse) |  | GET|/stargaze/globalfee/v1/contract_authorization/{contract_address}|
| `Params` | [QueryParamsRequest](#publicawesome.stargaze.globalfee.v1.QueryParamsRequest) | [QueryParamsResponse](#publicawesome.stargaze.globalfee.v1.QueryParamsResponse) |  | GET|/stargaze/globalfee/v1/params|
| `Authorizations` | [QueryAuthorizationsRequest](#publicawesome.stargaze.globalfee.v1.QueryAuthorizationsRequest) | [QueryAuthorizationsResponse](#publicawesome.stargaze.globalfee.v1.QueryAuthorizationsResponse) |  | GET|/stargaze/globalfee/v1/authorizations|

 <!-- end services -->



<a name="publicawesome/stargaze/globalfee/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/globalfee/v1/tx.proto



<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorization"></a>

### MsgRemoveCodeAuthorization
MsgRemoveCodeAuthorization is the request for removing code authorization.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `code_id` | [uint64](#uint64) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorizationResponse"></a>

### MsgRemoveCodeAuthorizationResponse
MsgRemoveCodeAuthorizationResponse is the response for executing remove authorization.






<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorization"></a>

### MsgRemoveContractAuthorization
MsgRemoveContractAuthorization is the request for removing contract authorization.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorizationResponse"></a>

### MsgRemoveContractAuthorizationResponse
MsgRemoveContractAuthorizationResponse is the repsonse for executing a contract authorization removal.






<a name="publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorization"></a>

### MsgSetCodeAuthorization
MsgSetCodeAuthorization is the request for setting code fee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `code_authorization` | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorizationResponse"></a>

### MsgSetCodeAuthorizationResponse
MsgSetCodeAuthorizationResponse is the response for executing a set code authorization.






<a name="publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorization"></a>

### MsgSetContractAuthorization
MsgSetContractAuthorization is the request for executing set contract authorization.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `contract_authorization` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorizationResponse"></a>

### MsgSetContractAuthorizationResponse
MsgSetContractAuthorizationResponse is the response for executing contract authorization.






<a name="publicawesome.stargaze.globalfee.v1.MsgUpdateParams"></a>

### MsgUpdateParams
MsgUpdateParams is the request for updating module's params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `params` | [Params](#publicawesome.stargaze.globalfee.v1.Params) |  | NOTE: All parameters must be supplied. |






<a name="publicawesome.stargaze.globalfee.v1.MsgUpdateParamsResponse"></a>

### MsgUpdateParamsResponse
MsgUpdateParamsResponse is the response for executiong a module's params update.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.globalfee.v1.Msg"></a>

### Msg
Msg defines the alloc Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetCodeAuthorization` | [MsgSetCodeAuthorization](#publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorization) | [MsgSetCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorizationResponse) | SetCodeAuthorization will set a specific code id fee settings. | |
| `RemoveCodeAuthorization` | [MsgRemoveCodeAuthorization](#publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorization) | [MsgRemoveCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorizationResponse) | RemoveCodeAuthorization will remove code id configuration. | |
| `SetContractAuthorization` | [MsgSetContractAuthorization](#publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorization) | [MsgSetContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorizationResponse) | SetContractAuthorization will set a specific contract fee settings. | |
| `RemoveContractAuthorization` | [MsgRemoveContractAuthorization](#publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorization) | [MsgRemoveContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorizationResponse) | RemoveContractAuthorization removes specific contract fee settings. | |
| `UpdateParams` | [MsgUpdateParams](#publicawesome.stargaze.globalfee.v1.MsgUpdateParams) | [MsgUpdateParamsResponse](#publicawesome.stargaze.globalfee.v1.MsgUpdateParamsResponse) | UpdateParams will update module params, callable by governance only. | |

 <!-- end services -->



<a name="publicawesome/stargaze/mint/v1beta1/mint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/mint/v1beta1/mint.proto



<a name="publicawesome.stargaze.mint.v1beta1.Minter"></a>

### Minter
Minter represents the minting state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `annual_provisions` | [string](#string) |  | current annual expected provisions |






<a name="publicawesome.stargaze.mint.v1beta1.Params"></a>

### Params
Params holds parameters for the mint module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mint_denom` | [string](#string) |  | type of coin to mint |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | the time the chain starts |
| `initial_annual_provisions` | [string](#string) |  | initial annual provisions |
| `reduction_factor` | [string](#string) |  | factor to reduce inflation by each year |
| `blocks_per_year` | [uint64](#uint64) |  | expected blocks per year |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/mint/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/mint/v1beta1/genesis.proto



<a name="publicawesome.stargaze.mint.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the mint module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minter` | [Minter](#publicawesome.stargaze.mint.v1beta1.Minter) |  | minter is a space for holding current inflation information. |
| `params` | [Params](#publicawesome.stargaze.mint.v1beta1.Params) |  | params defines all the paramaters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/mint/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/mint/v1beta1/query.proto



<a name="publicawesome.stargaze.mint.v1beta1.QueryAnnualProvisionsRequest"></a>

### QueryAnnualProvisionsRequest
QueryAnnualProvisionsRequest is the request type for the
Query/AnnualProvisions RPC method.






<a name="publicawesome.stargaze.mint.v1beta1.QueryAnnualProvisionsResponse"></a>

### QueryAnnualProvisionsResponse
QueryAnnualProvisionsResponse is the response type for the
Query/AnnualProvisions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `annual_provisions` | [bytes](#bytes) |  | annual_provisions is the current minting annual provisions value. |






<a name="publicawesome.stargaze.mint.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="publicawesome.stargaze.mint.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.mint.v1beta1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.mint.v1beta1.Query"></a>

### Query
Query provides defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#publicawesome.stargaze.mint.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#publicawesome.stargaze.mint.v1beta1.QueryParamsResponse) | Params returns the total set of minting parameters. | GET|/stargaze/mint/v1beta1/params|
| `AnnualProvisions` | [QueryAnnualProvisionsRequest](#publicawesome.stargaze.mint.v1beta1.QueryAnnualProvisionsRequest) | [QueryAnnualProvisionsResponse](#publicawesome.stargaze.mint.v1beta1.QueryAnnualProvisionsResponse) | AnnualProvisions current minting annual provisions value. | GET|/stargaze/mint/v1beta1/annual_provisions|

 <!-- end services -->



<a name="publicawesome/stargaze/mint/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/mint/v1beta1/tx.proto



<a name="publicawesome.stargaze.mint.v1beta1.MsgUpdateParams"></a>

### MsgUpdateParams
MsgUpdateParams is the request type for updating module's params.

Since: v14


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | Authority is the address of the governance account. |
| `params` | [Params](#publicawesome.stargaze.mint.v1beta1.Params) |  | NOTE: All parameters must be supplied. |






<a name="publicawesome.stargaze.mint.v1beta1.MsgUpdateParamsResponse"></a>

### MsgUpdateParamsResponse
MsgUpdateParamsResponse is the response type for executing
an update.
Since: v14





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.mint.v1beta1.Msg"></a>

### Msg
Msg defines the mint Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `UpdateParams` | [MsgUpdateParams](#publicawesome.stargaze.mint.v1beta1.MsgUpdateParams) | [MsgUpdateParamsResponse](#publicawesome.stargaze.mint.v1beta1.MsgUpdateParamsResponse) | UpdateParams updates the mint module's parameters. | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

