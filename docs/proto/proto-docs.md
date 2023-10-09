<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [publicawesome/stargaze/cron/v1/genesis.proto](#publicawesome/stargaze/cron/v1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.cron.v1.GenesisState)
  
- [publicawesome/stargaze/cron/v1/proposal.proto](#publicawesome/stargaze/cron/v1/proposal.proto)
    - [DemotePrivilegedContractProposal](#publicawesome.stargaze.cron.v1.DemotePrivilegedContractProposal)
    - [PromoteToPrivilegedContractProposal](#publicawesome.stargaze.cron.v1.PromoteToPrivilegedContractProposal)
  
- [publicawesome/stargaze/cron/v1/query.proto](#publicawesome/stargaze/cron/v1/query.proto)
    - [QueryListPrivilegedRequest](#publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest)
    - [QueryListPrivilegedResponse](#publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse)
  
    - [Query](#publicawesome.stargaze.cron.v1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="publicawesome/stargaze/cron/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/cron/v1/genesis.proto



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



<a name="publicawesome/stargaze/cron/v1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/cron/v1/proposal.proto



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

