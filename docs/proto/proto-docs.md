<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

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
  
- [Scalar Value Types](#scalar-value-types)



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

