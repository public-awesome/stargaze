<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [publicawesome/stargaze/authority/v1/authority.proto](#publicawesome/stargaze/authority/v1/authority.proto)
    - [Authorization](#publicawesome.stargaze.authority.v1.Authorization)
    - [Params](#publicawesome.stargaze.authority.v1.Params)
  
- [publicawesome/stargaze/authority/v1/genesis.proto](#publicawesome/stargaze/authority/v1/genesis.proto)
    - [GenesisState](#publicawesome.stargaze.authority.v1.GenesisState)
  
- [publicawesome/stargaze/authority/v1/query.proto](#publicawesome/stargaze/authority/v1/query.proto)
    - [QueryParamsRequest](#publicawesome.stargaze.authority.v1.QueryParamsRequest)
    - [QueryParamsResponse](#publicawesome.stargaze.authority.v1.QueryParamsResponse)
  
    - [Query](#publicawesome.stargaze.authority.v1.Query)
  
- [publicawesome/stargaze/authority/v1/tx.proto](#publicawesome/stargaze/authority/v1/tx.proto)
    - [MsgExecuteProposal](#publicawesome.stargaze.authority.v1.MsgExecuteProposal)
    - [MsgExecuteProposalResponse](#publicawesome.stargaze.authority.v1.MsgExecuteProposalResponse)
  
    - [Msg](#publicawesome.stargaze.authority.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="publicawesome/stargaze/authority/v1/authority.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/authority/v1/authority.proto



<a name="publicawesome.stargaze.authority.v1.Authorization"></a>

### Authorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msgTypeUrl` | [string](#string) |  | The type url of a proposal sdk.Msg |
| `addresses` | [string](#string) | repeated | The addresses who are allowed to execute that proposal |






<a name="publicawesome.stargaze.authority.v1.Params"></a>

### Params
Params holds parameters for the authority module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authorizations` | [Authorization](#publicawesome.stargaze.authority.v1.Authorization) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/authority/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/authority/v1/genesis.proto



<a name="publicawesome.stargaze.authority.v1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.authority.v1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="publicawesome/stargaze/authority/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/authority/v1/query.proto



<a name="publicawesome.stargaze.authority.v1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="publicawesome.stargaze.authority.v1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.authority.v1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.authority.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#publicawesome.stargaze.authority.v1.QueryParamsRequest) | [QueryParamsResponse](#publicawesome.stargaze.authority.v1.QueryParamsResponse) |  | GET|/stargaze/authority/v1/params|

 <!-- end services -->



<a name="publicawesome/stargaze/authority/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/authority/v1/tx.proto



<a name="publicawesome.stargaze.authority.v1.MsgExecuteProposal"></a>

### MsgExecuteProposal
MsgExecuteProposal defines an sdk.Msg type that supports submitting arbitrary
proposal Content.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  |  |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="publicawesome.stargaze.authority.v1.MsgExecuteProposalResponse"></a>

### MsgExecuteProposalResponse
MsgExecuteProposalResponse defines the Msg/ExecuteProposal response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.authority.v1.Msg"></a>

### Msg
Msg defines the authority Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ExecuteProposal` | [MsgExecuteProposal](#publicawesome.stargaze.authority.v1.MsgExecuteProposal) | [MsgExecuteProposalResponse](#publicawesome.stargaze.authority.v1.MsgExecuteProposalResponse) |  | |

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

