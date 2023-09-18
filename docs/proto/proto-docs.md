<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

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
  
    - [Msg](#publicawesome.stargaze.alloc.v1beta1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="publicawesome/stargaze/alloc/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/alloc/v1beta1/params.proto



<a name="publicawesome.stargaze.alloc.v1beta1.DistributionProportions"></a>

### DistributionProportions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft_incentives` | [string](#string) |  |  |
| `developer_rewards` | [string](#string) |  |  |
| `community_pool` | [string](#string) |  |  |






<a name="publicawesome.stargaze.alloc.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `distribution_proportions` | [DistributionProportions](#publicawesome.stargaze.alloc.v1beta1.DistributionProportions) |  | distribution_proportions defines the proportion of the minted denom |
| `weighted_developer_rewards_receivers` | [WeightedAddress](#publicawesome.stargaze.alloc.v1beta1.WeightedAddress) | repeated | addresses to receive developer rewards |
| `weighted_incentives_rewards_receivers` | [WeightedAddress](#publicawesome.stargaze.alloc.v1beta1.WeightedAddress) | repeated | addresses to receive incentive rewards |
| `supplement_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | SupplementAmount is the amount to be supplemented from the pool on top of newly minted coins. |






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

