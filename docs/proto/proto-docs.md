<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

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
  
- [stargaze/cron/genesis.proto](#stargaze/cron/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.cron.v1.GenesisState)
  
- [stargaze/cron/proposal.proto](#stargaze/cron/proposal.proto)
    - [DemotePrivilegedContractProposal](#publicawesome.stargaze.cron.v1.DemotePrivilegedContractProposal)
    - [PromoteToPrivilegedContractProposal](#publicawesome.stargaze.cron.v1.PromoteToPrivilegedContractProposal)
  
- [stargaze/cron/query.proto](#stargaze/cron/query.proto)
    - [QueryListPrivilegedRequest](#publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest)
    - [QueryListPrivilegedResponse](#publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse)
  
    - [Query](#publicawesome.stargaze.cron.v1.Query)
  
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



<a name="stargaze/cron/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/cron/genesis.proto



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



<a name="stargaze/cron/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/cron/proposal.proto



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



<a name="stargaze/cron/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stargaze/cron/query.proto



<a name="publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest"></a>

### QueryListPrivilegedRequest
QueryListPrivilegedRequest is request type for the Query/ListPrivileged RPC method.






<a name="publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse"></a>

### QueryListPrivilegedResponse
QueryListPrivilegedResponse is response type for the Query/ListPrivileged RPC method.


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

