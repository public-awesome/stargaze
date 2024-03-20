<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

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
    - [Msg](#publicawesome.stargaze.mint.v1beta1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



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


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="publicawesome.stargaze.mint.v1beta1.Msg"></a>

### Msg
Msg defines the mint Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |

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

