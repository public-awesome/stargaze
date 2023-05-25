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
  
    - [Msg](#osmosis.tokenfactory.v1beta1.Msg)
  
- [stargaze/alloc/v1beta1/params.proto](#stargaze/alloc/v1beta1/params.proto)
    - [DistributionProportions](#publicawesome.stargaze.alloc.v1beta1.DistributionProportions)
    - [Params](#publicawesome.stargaze.alloc.v1beta1.Params)
    - [WeightedAddress](#publicawesome.stargaze.alloc.v1beta1.WeightedAddress)
  
- [stargaze/alloc/v1beta1/genesis.proto](#stargaze/alloc/v1beta1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.alloc.v1beta1.GenesisState)
  
- [stargaze/alloc/v1beta1/query.proto](#stargaze/alloc/v1beta1/query.proto)
    - [QueryParamsRequest](#publicawesome.stargaze.alloc.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#publicawesome.stargaze.alloc.v1beta1.QueryParamsResponse)
  
    - [Query](#publicawesome.stargaze.alloc.v1beta1.Query)
  
- [stargaze/alloc/v1beta1/tx.proto](#stargaze/alloc/v1beta1/tx.proto)
    - [MsgCreateVestingAccount](#publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccount)
    - [MsgCreateVestingAccountResponse](#publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccountResponse)
    - [MsgFundFairburnPool](#publicawesome.stargaze.alloc.v1beta1.MsgFundFairburnPool)
    - [MsgFundFairburnPoolResponse](#publicawesome.stargaze.alloc.v1beta1.MsgFundFairburnPoolResponse)
  
    - [Msg](#publicawesome.stargaze.alloc.v1beta1.Msg)
  
- [stargaze/claim/v1beta1/claim_record.proto](#stargaze/claim/v1beta1/claim_record.proto)
    - [ClaimRecord](#publicawesome.stargaze.claim.v1beta1.ClaimRecord)
  
    - [Action](#publicawesome.stargaze.claim.v1beta1.Action)
  
- [stargaze/claim/v1beta1/params.proto](#stargaze/claim/v1beta1/params.proto)
    - [ClaimAuthorization](#publicawesome.stargaze.claim.v1beta1.ClaimAuthorization)
    - [Params](#publicawesome.stargaze.claim.v1beta1.Params)
  
- [stargaze/claim/v1beta1/genesis.proto](#stargaze/claim/v1beta1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.claim.v1beta1.GenesisState)
  
- [stargaze/claim/v1beta1/query.proto](#stargaze/claim/v1beta1/query.proto)
    - [QueryClaimRecordRequest](#publicawesome.stargaze.claim.v1beta1.QueryClaimRecordRequest)
    - [QueryClaimRecordResponse](#publicawesome.stargaze.claim.v1beta1.QueryClaimRecordResponse)
    - [QueryClaimableForActionRequest](#publicawesome.stargaze.claim.v1beta1.QueryClaimableForActionRequest)
    - [QueryClaimableForActionResponse](#publicawesome.stargaze.claim.v1beta1.QueryClaimableForActionResponse)
    - [QueryModuleAccountBalanceRequest](#publicawesome.stargaze.claim.v1beta1.QueryModuleAccountBalanceRequest)
    - [QueryModuleAccountBalanceResponse](#publicawesome.stargaze.claim.v1beta1.QueryModuleAccountBalanceResponse)
    - [QueryParamsRequest](#publicawesome.stargaze.claim.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#publicawesome.stargaze.claim.v1beta1.QueryParamsResponse)
    - [QueryTotalClaimableRequest](#publicawesome.stargaze.claim.v1beta1.QueryTotalClaimableRequest)
    - [QueryTotalClaimableResponse](#publicawesome.stargaze.claim.v1beta1.QueryTotalClaimableResponse)
  
    - [Query](#publicawesome.stargaze.claim.v1beta1.Query)
  
- [stargaze/claim/v1beta1/tx.proto](#stargaze/claim/v1beta1/tx.proto)
    - [MsgClaimFor](#publicawesome.stargaze.claim.v1beta1.MsgClaimFor)
    - [MsgClaimForResponse](#publicawesome.stargaze.claim.v1beta1.MsgClaimForResponse)
    - [MsgInitialClaim](#publicawesome.stargaze.claim.v1beta1.MsgInitialClaim)
    - [MsgInitialClaimResponse](#publicawesome.stargaze.claim.v1beta1.MsgInitialClaimResponse)
  
    - [Msg](#publicawesome.stargaze.claim.v1beta1.Msg)
  
- [stargaze/cron/v1/genesis.proto](#stargaze/cron/v1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.cron.v1.GenesisState)
  
- [stargaze/cron/v1/proposal.proto](#stargaze/cron/v1/proposal.proto)
    - [DemotePrivilegedContractProposal](#publicawesome.stargaze.cron.v1.DemotePrivilegedContractProposal)
    - [PromoteToPrivilegedContractProposal](#publicawesome.stargaze.cron.v1.PromoteToPrivilegedContractProposal)
  
- [stargaze/cron/v1/query.proto](#stargaze/cron/v1/query.proto)
    - [QueryListPrivilegedRequest](#publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest)
    - [QueryListPrivilegedResponse](#publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse)
  
    - [Query](#publicawesome.stargaze.cron.v1.Query)
  
- [stargaze/globalfee/v1/globalfee.proto](#stargaze/globalfee/v1/globalfee.proto)
    - [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization)
    - [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization)
    - [Params](#publicawesome.stargaze.globalfee.v1.Params)
  
- [stargaze/globalfee/v1/genesis.proto](#stargaze/globalfee/v1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.globalfee.v1.GenesisState)
  
- [stargaze/globalfee/v1/proposal.proto](#stargaze/globalfee/v1/proposal.proto)
    - [RemoveCodeAuthorizationProposal](#publicawesome.stargaze.globalfee.v1.RemoveCodeAuthorizationProposal)
    - [RemoveContractAuthorizationProposal](#publicawesome.stargaze.globalfee.v1.RemoveContractAuthorizationProposal)
    - [SetCodeAuthorizationProposal](#publicawesome.stargaze.globalfee.v1.SetCodeAuthorizationProposal)
    - [SetContractAuthorizationProposal](#publicawesome.stargaze.globalfee.v1.SetContractAuthorizationProposal)
  
- [stargaze/globalfee/v1/query.proto](#stargaze/globalfee/v1/query.proto)
    - [QueryAuthorizationsRequest](#publicawesome.stargaze.globalfee.v1.QueryAuthorizationsRequest)
    - [QueryAuthorizationsResponse](#publicawesome.stargaze.globalfee.v1.QueryAuthorizationsResponse)
    - [QueryCodeAuthorizationRequest](#publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationRequest)
    - [QueryCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationResponse)
    - [QueryContractAuthorizationRequest](#publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationRequest)
    - [QueryContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationResponse)
    - [QueryParamsRequest](#publicawesome.stargaze.globalfee.v1.QueryParamsRequest)
    - [QueryParamsResponse](#publicawesome.stargaze.globalfee.v1.QueryParamsResponse)
  
    - [Query](#publicawesome.stargaze.globalfee.v1.Query)
  
- [stargaze/globalfee/v1/tx.proto](#stargaze/globalfee/v1/tx.proto)
    - [MsgRemoveCodeAuthorization](#publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorization)
    - [MsgRemoveCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorizationResponse)
    - [MsgRemoveContractAuthorization](#publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorization)
    - [MsgRemoveContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorizationResponse)
    - [MsgSetCodeAuthorization](#publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorization)
    - [MsgSetCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorizationResponse)
    - [MsgSetContractAuthorization](#publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorization)
    - [MsgSetContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorizationResponse)
  
    - [Msg](#publicawesome.stargaze.globalfee.v1.Msg)
  
- [stargaze/mint/v1beta1/mint.proto](#stargaze/mint/v1beta1/mint.proto)
    - [Minter](#stargaze.mint.v1beta1.Minter)
    - [Params](#stargaze.mint.v1beta1.Params)
  
- [stargaze/mint/v1beta1/genesis.proto](#stargaze/mint/v1beta1/genesis.proto)
    - [GenesisState](#stargaze.mint.v1beta1.GenesisState)
  
- [stargaze/mint/v1beta1/query.proto](#stargaze/mint/v1beta1/query.proto)
    - [QueryAnnualProvisionsRequest](#stargaze.mint.v1beta1.QueryAnnualProvisionsRequest)
    - [QueryAnnualProvisionsResponse](#stargaze.mint.v1beta1.QueryAnnualProvisionsResponse)
    - [QueryParamsRequest](#stargaze.mint.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#stargaze.mint.v1beta1.QueryParamsResponse)
  
    - [Query](#stargaze.mint.v1beta1.Query)
  
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





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="osmosis.tokenfactory.v1beta1.Msg"></a>

### Msg
Msg defines the tokefactory module's gRPC message service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateDenom` | [MsgCreateDenom](#osmosis.tokenfactory.v1beta1.MsgCreateDenom) | [MsgCreateDenomResponse](#osmosis.tokenfactory.v1beta1.MsgCreateDenomResponse) |  | |
| `Mint` | [MsgMint](#osmosis.tokenfactory.v1beta1.MsgMint) | [MsgMintResponse](#osmosis.tokenfactory.v1beta1.MsgMintResponse) |  | |
| `Burn` | [MsgBurn](#osmosis.tokenfactory.v1beta1.MsgBurn) | [MsgBurnResponse](#osmosis.tokenfactory.v1beta1.MsgBurnResponse) |  | |
| `ChangeAdmin` | [MsgChangeAdmin](#osmosis.tokenfactory.v1beta1.MsgChangeAdmin) | [MsgChangeAdminResponse](#osmosis.tokenfactory.v1beta1.MsgChangeAdminResponse) |  | |
| `SetDenomMetadata` | [MsgSetDenomMetadata](#osmosis.tokenfactory.v1beta1.MsgSetDenomMetadata) | [MsgSetDenomMetadataResponse](#osmosis.tokenfactory.v1beta1.MsgSetDenomMetadataResponse) |  | |

 <!-- end services -->



<a name="stargaze/alloc/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/alloc/v1beta1/params.proto



<a name="publicawesome.stargaze.alloc.v1beta1.DistributionProportions"></a>

### DistributionProportions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_incentives` | [string](#string) |  |  |
| `developer_rewards` | [string](#string) |  |  |






<a name="publicawesome.stargaze.alloc.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `distribution_proportions` | [DistributionProportions](#publicawesome.stargaze.alloc.v1beta1.DistributionProportions) |  | distribution_proportions defines the proportion of the minted denom |
| `weighted_developer_rewards_receivers` | [WeightedAddress](#publicawesome.stargaze.alloc.v1beta1.WeightedAddress) | repeated | address to receive developer rewards |






<a name="publicawesome.stargaze.alloc.v1beta1.WeightedAddress"></a>

### WeightedAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stargaze/alloc/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/alloc/v1beta1/genesis.proto



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



<a name="stargaze/alloc/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/alloc/v1beta1/query.proto



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



<a name="stargaze/alloc/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/alloc/v1beta1/tx.proto



<a name="publicawesome.stargaze.alloc.v1beta1.MsgCreateVestingAccount"></a>

### MsgCreateVestingAccount
MsgCreateVestingAccount defines a message that enables creating a vesting
account.


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

 <!-- end services -->



<a name="stargaze/claim/v1beta1/claim_record.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/claim/v1beta1/claim_record.proto



<a name="publicawesome.stargaze.claim.v1beta1.ClaimRecord"></a>

### ClaimRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address of claim user |
| `initial_claimable_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | total initial claimable amount for the user |
| `action_completed` | [bool](#bool) | repeated | true if action is completed index of bool in array refers to action enum # |





 <!-- end messages -->


<a name="publicawesome.stargaze.claim.v1beta1.Action"></a>

### Action


| Name | Number | Description |
| ---- | ------ | ----------- |
| ActionInitialClaim | 0 |  |
| ActionBidNFT | 1 |  |
| ActionMintNFT | 2 |  |
| ActionVote | 3 |  |
| ActionDelegateStake | 4 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stargaze/claim/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/claim/v1beta1/params.proto



<a name="publicawesome.stargaze.claim.v1beta1.ClaimAuthorization"></a>

### ClaimAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `action` | [Action](#publicawesome.stargaze.claim.v1beta1.Action) |  |  |






<a name="publicawesome.stargaze.claim.v1beta1.Params"></a>

### Params
Params defines the claim module's parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `airdrop_enabled` | [bool](#bool) |  |  |
| `airdrop_start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `duration_until_decay` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `duration_of_decay` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `claim_denom` | [string](#string) |  | denom of claimable asset |
| `allowed_claimers` | [ClaimAuthorization](#publicawesome.stargaze.claim.v1beta1.ClaimAuthorization) | repeated | list of contracts and their allowed claim actions |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stargaze/claim/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/claim/v1beta1/genesis.proto



<a name="publicawesome.stargaze.claim.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the claim module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module_account_balance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | this line is used by starport scaffolding # genesis/proto/state balance of the claim module's account |
| `params` | [Params](#publicawesome.stargaze.claim.v1beta1.Params) |  | params defines all the parameters of the module. |
| `claim_records` | [ClaimRecord](#publicawesome.stargaze.claim.v1beta1.ClaimRecord) | repeated | list of claim records, one for every airdrop recipient |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stargaze/claim/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/claim/v1beta1/query.proto



<a name="publicawesome.stargaze.claim.v1beta1.QueryClaimRecordRequest"></a>

### QueryClaimRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="publicawesome.stargaze.claim.v1beta1.QueryClaimRecordResponse"></a>

### QueryClaimRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `claim_record` | [ClaimRecord](#publicawesome.stargaze.claim.v1beta1.ClaimRecord) |  |  |






<a name="publicawesome.stargaze.claim.v1beta1.QueryClaimableForActionRequest"></a>

### QueryClaimableForActionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `action` | [Action](#publicawesome.stargaze.claim.v1beta1.Action) |  |  |






<a name="publicawesome.stargaze.claim.v1beta1.QueryClaimableForActionResponse"></a>

### QueryClaimableForActionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="publicawesome.stargaze.claim.v1beta1.QueryModuleAccountBalanceRequest"></a>

### QueryModuleAccountBalanceRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="publicawesome.stargaze.claim.v1beta1.QueryModuleAccountBalanceResponse"></a>

### QueryModuleAccountBalanceResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `moduleAccountBalance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | params defines the parameters of the module. |






<a name="publicawesome.stargaze.claim.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="publicawesome.stargaze.claim.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.claim.v1beta1.Params) |  | params defines the parameters of the module. |






<a name="publicawesome.stargaze.claim.v1beta1.QueryTotalClaimableRequest"></a>

### QueryTotalClaimableRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="publicawesome.stargaze.claim.v1beta1.QueryTotalClaimableResponse"></a>

### QueryTotalClaimableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.claim.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

this line is used by starport scaffolding # 3

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ModuleAccountBalance` | [QueryModuleAccountBalanceRequest](#publicawesome.stargaze.claim.v1beta1.QueryModuleAccountBalanceRequest) | [QueryModuleAccountBalanceResponse](#publicawesome.stargaze.claim.v1beta1.QueryModuleAccountBalanceResponse) | this line is used by starport scaffolding # 2 | GET|/stargaze/claim/v1beta1/module_account_balance|
| `Params` | [QueryParamsRequest](#publicawesome.stargaze.claim.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#publicawesome.stargaze.claim.v1beta1.QueryParamsResponse) |  | GET|/stargaze/claim/v1beta1/params|
| `ClaimRecord` | [QueryClaimRecordRequest](#publicawesome.stargaze.claim.v1beta1.QueryClaimRecordRequest) | [QueryClaimRecordResponse](#publicawesome.stargaze.claim.v1beta1.QueryClaimRecordResponse) |  | GET|/stargaze/claim/v1beta1/claim_record/{address}|
| `ClaimableForAction` | [QueryClaimableForActionRequest](#publicawesome.stargaze.claim.v1beta1.QueryClaimableForActionRequest) | [QueryClaimableForActionResponse](#publicawesome.stargaze.claim.v1beta1.QueryClaimableForActionResponse) |  | GET|/stargaze/claim/v1beta1/claimable_for_action/{address}/{action}|
| `TotalClaimable` | [QueryTotalClaimableRequest](#publicawesome.stargaze.claim.v1beta1.QueryTotalClaimableRequest) | [QueryTotalClaimableResponse](#publicawesome.stargaze.claim.v1beta1.QueryTotalClaimableResponse) |  | GET|/stargaze/claim/v1beta1/total_claimable/{address}|

 <!-- end services -->



<a name="stargaze/claim/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/claim/v1beta1/tx.proto



<a name="publicawesome.stargaze.claim.v1beta1.MsgClaimFor"></a>

### MsgClaimFor



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `action` | [Action](#publicawesome.stargaze.claim.v1beta1.Action) |  |  |






<a name="publicawesome.stargaze.claim.v1beta1.MsgClaimForResponse"></a>

### MsgClaimForResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `claimed_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | total initial claimable amount for the user |






<a name="publicawesome.stargaze.claim.v1beta1.MsgInitialClaim"></a>

### MsgInitialClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |






<a name="publicawesome.stargaze.claim.v1beta1.MsgInitialClaimResponse"></a>

### MsgInitialClaimResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `claimed_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | total initial claimable amount for the user |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.claim.v1beta1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `InitialClaim` | [MsgInitialClaim](#publicawesome.stargaze.claim.v1beta1.MsgInitialClaim) | [MsgInitialClaimResponse](#publicawesome.stargaze.claim.v1beta1.MsgInitialClaimResponse) |  | |
| `ClaimFor` | [MsgClaimFor](#publicawesome.stargaze.claim.v1beta1.MsgClaimFor) | [MsgClaimForResponse](#publicawesome.stargaze.claim.v1beta1.MsgClaimForResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



<a name="stargaze/cron/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/cron/v1/genesis.proto



<a name="publicawesome.stargaze.cron.v1.GenesisState"></a>

### GenesisState
GenesisState defines the cron module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `privileged_contract_addresses` | [string](#string) | repeated | List of all the contracts that have been given the privilege status via governance. They can set up hooks to abci.EndBlocker |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stargaze/cron/v1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/cron/v1/proposal.proto



<a name="publicawesome.stargaze.cron.v1.DemotePrivilegedContractProposal"></a>

### DemotePrivilegedContractProposal
DemotePrivilegedContractProposal gov proposal content type to remove
"privileges" from a contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `contract` | [string](#string) |  | Contract is the bech32 address of the smart contract |






<a name="publicawesome.stargaze.cron.v1.PromoteToPrivilegedContractProposal"></a>

### PromoteToPrivilegedContractProposal
PromoteToPrivilegedContractProposal gov proposal content type to add
"privileges" to a contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `contract` | [string](#string) |  | Contract is the bech32 address of the smart contract |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stargaze/cron/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/cron/v1/query.proto



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





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.cron.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ListPrivileged` | [QueryListPrivilegedRequest](#publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest) | [QueryListPrivilegedResponse](#publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse) | ListPrivileged queries the contracts which have the priviledge status | GET|/stargaze/cron/v1/list-privileged|

 <!-- end services -->



<a name="stargaze/globalfee/v1/globalfee.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/globalfee/v1/globalfee.proto



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



<a name="stargaze/globalfee/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/globalfee/v1/genesis.proto



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



<a name="stargaze/globalfee/v1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/globalfee/v1/proposal.proto



<a name="publicawesome.stargaze.globalfee.v1.RemoveCodeAuthorizationProposal"></a>

### RemoveCodeAuthorizationProposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `code_id` | [uint64](#uint64) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.RemoveContractAuthorizationProposal"></a>

### RemoveContractAuthorizationProposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.SetCodeAuthorizationProposal"></a>

### SetCodeAuthorizationProposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `code_authorization` | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.SetContractAuthorizationProposal"></a>

### SetContractAuthorizationProposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `contract_authorization` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stargaze/globalfee/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/globalfee/v1/query.proto



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



<a name="stargaze/globalfee/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/globalfee/v1/tx.proto



<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorization"></a>

### MsgRemoveCodeAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `code_id` | [uint64](#uint64) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorizationResponse"></a>

### MsgRemoveCodeAuthorizationResponse







<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorization"></a>

### MsgRemoveContractAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorizationResponse"></a>

### MsgRemoveContractAuthorizationResponse







<a name="publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorization"></a>

### MsgSetCodeAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `code_authorization` | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorizationResponse"></a>

### MsgSetCodeAuthorizationResponse







<a name="publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorization"></a>

### MsgSetContractAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `contract_authorization` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) |  |  |






<a name="publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorizationResponse"></a>

### MsgSetContractAuthorizationResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.globalfee.v1.Msg"></a>

### Msg
Msg defines the alloc Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetCodeAuthorization` | [MsgSetCodeAuthorization](#publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorization) | [MsgSetCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorizationResponse) |  | |
| `RemoveCodeAuthorization` | [MsgRemoveCodeAuthorization](#publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorization) | [MsgRemoveCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorizationResponse) |  | |
| `SetContractAuthorization` | [MsgSetContractAuthorization](#publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorization) | [MsgSetContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorizationResponse) |  | |
| `RemoveContractAuthorization` | [MsgRemoveContractAuthorization](#publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorization) | [MsgRemoveContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorizationResponse) |  | |

 <!-- end services -->



<a name="stargaze/mint/v1beta1/mint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/mint/v1beta1/mint.proto



<a name="stargaze.mint.v1beta1.Minter"></a>

### Minter
Minter represents the minting state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `annual_provisions` | [string](#string) |  | current annual expected provisions |






<a name="stargaze.mint.v1beta1.Params"></a>

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



<a name="stargaze/mint/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/mint/v1beta1/genesis.proto



<a name="stargaze.mint.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the mint module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minter` | [Minter](#stargaze.mint.v1beta1.Minter) |  | minter is a space for holding current inflation information. |
| `params` | [Params](#stargaze.mint.v1beta1.Params) |  | params defines all the paramaters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stargaze/mint/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/mint/v1beta1/query.proto



<a name="stargaze.mint.v1beta1.QueryAnnualProvisionsRequest"></a>

### QueryAnnualProvisionsRequest
QueryAnnualProvisionsRequest is the request type for the
Query/AnnualProvisions RPC method.






<a name="stargaze.mint.v1beta1.QueryAnnualProvisionsResponse"></a>

### QueryAnnualProvisionsResponse
QueryAnnualProvisionsResponse is the response type for the
Query/AnnualProvisions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `annual_provisions` | [bytes](#bytes) |  | annual_provisions is the current minting annual provisions value. |






<a name="stargaze.mint.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="stargaze.mint.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#stargaze.mint.v1beta1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stargaze.mint.v1beta1.Query"></a>

### Query
Query provides defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#stargaze.mint.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#stargaze.mint.v1beta1.QueryParamsResponse) | Params returns the total set of minting parameters. | GET|/stargaze/mint/v1beta1/params|
| `AnnualProvisions` | [QueryAnnualProvisionsRequest](#stargaze.mint.v1beta1.QueryAnnualProvisionsRequest) | [QueryAnnualProvisionsResponse](#stargaze.mint.v1beta1.QueryAnnualProvisionsResponse) | AnnualProvisions current minting annual provisions value. | GET|/stargaze/mint/v1beta1/annual_provisions|

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

