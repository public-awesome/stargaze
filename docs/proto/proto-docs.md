<!-- This file is auto-generated. Please do not modify it yourself. -->

# Protobuf Documentation

<a name="top"></a>

## Table of Contents

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

- [Scalar Value Types](#scalar-value-types)

<a name="publicawesome/stargaze/globalfee/v1/globalfee.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/globalfee/v1/globalfee.proto

<a name="publicawesome.stargaze.globalfee.v1.CodeAuthorization"></a>

### CodeAuthorization

Configuration for code Ids which can have zero gas operations

| Field     | Type              | Label    | Description                           |
| --------- | ----------------- | -------- | ------------------------------------- |
| `code_id` | [uint64](#uint64) |          | authorized code ids                   |
| `methods` | [string](#string) | repeated | authorized contract operation methods |

<a name="publicawesome.stargaze.globalfee.v1.ContractAuthorization"></a>

### ContractAuthorization

Configuration for contract addresses which can have zero gas operations

| Field              | Type              | Label    | Description                           |
| ------------------ | ----------------- | -------- | ------------------------------------- |
| `contract_address` | [string](#string) |          | authorized contract addresses         |
| `methods`          | [string](#string) | repeated | authorized contract operation methods |

<a name="publicawesome.stargaze.globalfee.v1.Params"></a>

### Params

Params holds parameters for the globalfee module.

| Field                  | Type                                                        | Label    | Description                                                       |
| ---------------------- | ----------------------------------------------------------- | -------- | ----------------------------------------------------------------- |
| `privileged_addresses` | [string](#string)                                           | repeated | Addresses which are whitelisted to modify the gas free operations |
| `minimum_gas_prices`   | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | Minimum stores the minimum gas price(s) for all TX on the chain.  |

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

| Field                     | Type                                                                                | Label    | Description                                     |
| ------------------------- | ----------------------------------------------------------------------------------- | -------- | ----------------------------------------------- |
| `params`                  | [Params](#publicawesome.stargaze.globalfee.v1.Params)                               |          | Module params                                   |
| `code_authorizations`     | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization)         | repeated | Authorizations configured by code id            |
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

| Field         | Type              | Label | Description |
| ------------- | ----------------- | ----- | ----------- |
| `title`       | [string](#string) |       |             |
| `description` | [string](#string) |       |             |
| `code_id`     | [uint64](#uint64) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.RemoveContractAuthorizationProposal"></a>

### RemoveContractAuthorizationProposal

| Field              | Type              | Label | Description |
| ------------------ | ----------------- | ----- | ----------- |
| `title`            | [string](#string) |       |             |
| `description`      | [string](#string) |       |             |
| `contract_address` | [string](#string) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.SetCodeAuthorizationProposal"></a>

### SetCodeAuthorizationProposal

| Field                | Type                                                                        | Label | Description |
| -------------------- | --------------------------------------------------------------------------- | ----- | ----------- |
| `title`              | [string](#string)                                                           |       |             |
| `description`        | [string](#string)                                                           |       |             |
| `code_authorization` | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.SetContractAuthorizationProposal"></a>

### SetContractAuthorizationProposal

| Field                    | Type                                                                                | Label | Description |
| ------------------------ | ----------------------------------------------------------------------------------- | ----- | ----------- |
| `title`                  | [string](#string)                                                                   |       |             |
| `description`            | [string](#string)                                                                   |       |             |
| `contract_authorization` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) |       |             |

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

| Field                     | Type                                                                                | Label    | Description |
| ------------------------- | ----------------------------------------------------------------------------------- | -------- | ----------- |
| `code_authorizations`     | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization)         | repeated |             |
| `contract_authorizations` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) | repeated |             |

<a name="publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationRequest"></a>

### QueryCodeAuthorizationRequest

| Field     | Type              | Label | Description |
| --------- | ----------------- | ----- | ----------- |
| `code_id` | [uint64](#uint64) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationResponse"></a>

### QueryCodeAuthorizationResponse

| Field     | Type              | Label    | Description |
| --------- | ----------------- | -------- | ----------- |
| `methods` | [string](#string) | repeated |             |

<a name="publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationRequest"></a>

### QueryContractAuthorizationRequest

| Field              | Type              | Label | Description |
| ------------------ | ----------------- | ----- | ----------- |
| `contract_address` | [string](#string) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationResponse"></a>

### QueryContractAuthorizationResponse

| Field     | Type              | Label    | Description |
| --------- | ----------------- | -------- | ----------- |
| `methods` | [string](#string) | repeated |             |

<a name="publicawesome.stargaze.globalfee.v1.QueryParamsRequest"></a>

### QueryParamsRequest

<a name="publicawesome.stargaze.globalfee.v1.QueryParamsResponse"></a>

### QueryParamsResponse

| Field    | Type                                                  | Label | Description |
| -------- | ----------------------------------------------------- | ----- | ----------- |
| `params` | [Params](#publicawesome.stargaze.globalfee.v1.Params) |       |             |

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="publicawesome.stargaze.globalfee.v1.Query"></a>

### Query

Query defines the gRPC querier service.

| Method Name             | Request Type                                                                                                | Response Type                                                                                                 | Description | HTTP Verb | Endpoint                                                         |
| ----------------------- | ----------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ----------- | --------- | ---------------------------------------------------------------- |
| `CodeAuthorization`     | [QueryCodeAuthorizationRequest](#publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationRequest)         | [QueryCodeAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.QueryCodeAuthorizationResponse)         |             | GET       | /stargaze/globalfee/v1/code_authorization/{code_id}              |
| `ContractAuthorization` | [QueryContractAuthorizationRequest](#publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationRequest) | [QueryContractAuthorizationResponse](#publicawesome.stargaze.globalfee.v1.QueryContractAuthorizationResponse) |             | GET       | /stargaze/globalfee/v1/contract_authorization/{contract_address} |
| `Params`                | [QueryParamsRequest](#publicawesome.stargaze.globalfee.v1.QueryParamsRequest)                               | [QueryParamsResponse](#publicawesome.stargaze.globalfee.v1.QueryParamsResponse)                               |             | GET       | /stargaze/globalfee/v1/params                                    |
| `Authorizations`        | [QueryAuthorizationsRequest](#publicawesome.stargaze.globalfee.v1.QueryAuthorizationsRequest)               | [QueryAuthorizationsResponse](#publicawesome.stargaze.globalfee.v1.QueryAuthorizationsResponse)               |             | GET       | /stargaze/globalfee/v1/authorizations                            |

 <!-- end services -->

<a name="publicawesome/stargaze/globalfee/v1/tx.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## publicawesome/stargaze/globalfee/v1/tx.proto

<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorization"></a>

### MsgRemoveCodeAuthorization

| Field     | Type              | Label | Description |
| --------- | ----------------- | ----- | ----------- |
| `sender`  | [string](#string) |       |             |
| `code_id` | [uint64](#uint64) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorizationResponse"></a>

### MsgRemoveCodeAuthorizationResponse

<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorization"></a>

### MsgRemoveContractAuthorization

| Field              | Type              | Label | Description |
| ------------------ | ----------------- | ----- | ----------- |
| `sender`           | [string](#string) |       |             |
| `contract_address` | [string](#string) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorizationResponse"></a>

### MsgRemoveContractAuthorizationResponse

<a name="publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorization"></a>

### MsgSetCodeAuthorization

| Field                | Type                                                                        | Label | Description |
| -------------------- | --------------------------------------------------------------------------- | ----- | ----------- |
| `sender`             | [string](#string)                                                           |       |             |
| `code_authorization` | [CodeAuthorization](#publicawesome.stargaze.globalfee.v1.CodeAuthorization) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorizationResponse"></a>

### MsgSetCodeAuthorizationResponse

<a name="publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorization"></a>

### MsgPromoteToPrivilegedContract

MsgPromoteToPrivilegedContract defines the Msg/PromoteToPrivilegedContract

### MsgSetContractAuthorization

| Field                    | Type                                                                                | Label | Description |
| ------------------------ | ----------------------------------------------------------------------------------- | ----- | ----------- |
| `sender`                 | [string](#string)                                                                   |       |             |
| `contract_authorization` | [ContractAuthorization](#publicawesome.stargaze.globalfee.v1.ContractAuthorization) |       |             |

<a name="publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorizationResponse"></a>

### MsgSetContractAuthorizationResponse

<a name="publicawesome.stargaze.globalfee.v1.MsgUpdateParams"></a>

### MsgUpdateParams

| Field    | Type                                                  | Label | Description                            |
| -------- | ----------------------------------------------------- | ----- | -------------------------------------- |
| `sender` | [string](#string)                                     |       |                                        |
| `params` | [Params](#publicawesome.stargaze.globalfee.v1.Params) |       | NOTE: All parameters must be supplied. |

<a name="publicawesome.stargaze.globalfee.v1.MsgUpdateParamsResponse"></a>

### MsgCreateDenomResponse
MsgCreateDenomResponse is the return value of MsgCreateDenom
It returns the full string of the newly created denom

 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

<a name="publicawesome.stargaze.globalfee.v1.Msg"></a>

### Msg

Msg defines the alloc Msg service.

| Method Name                    | Request Type                                                                                       | Response Type                                                                                                      | Description                                                             | HTTP Verb | Endpoint |
| ------------------------------ | -------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------- | --------- | -------- |
| `PromoteToPrivilegedContract`  | [MsgPromoteToPrivilegedContract](#publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContract)   | [MsgPromoteToPrivilegedContractResponse](#publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContractResponse)   | PromoteToPrivilegedContract promotes a contract to privileged status.   |           |
| `DemoteFromPrivilegedContract` | [MsgDemoteFromPrivilegedContract](#publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContract) | [MsgDemoteFromPrivilegedContractResponse](#publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContractResponse) | DemoteFromPrivilegedContract demotes a contract from privileged status. |           |
| `UpdateParams`                 | [MsgUpdateParams](#publicawesome.stargaze.cron.v1.MsgUpdateParams)                                 | [MsgUpdateParamsResponse](#publicawesome.stargaze.cron.v1.MsgUpdateParamsResponse)                                 | UpdateParams updates the cron module's parameters.                      |           |

 <!-- end services -->

## Scalar Value Types

| .proto Type                    | Notes                                                                                                                                           | C++    | Java       | Python      | Go      | C#         | PHP            | Ruby                           |
| ------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------ | ---------- | ----------- | ------- | ---------- | -------------- | ------------------------------ |
| <a name="double" /> double     |                                                                                                                                                 | double | double     | float       | float64 | double     | float          | Float                          |
| <a name="float" /> float       |                                                                                                                                                 | float  | float      | float       | float32 | float      | float          | Float                          |
| <a name="int32" /> int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="int64" /> int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="uint32" /> uint32     | Uses variable-length encoding.                                                                                                                  | uint32 | int        | int/long    | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64     | Uses variable-length encoding.                                                                                                                  | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="fixed32" /> fixed32   | Always four bytes. More efficient than uint32 if values are often greater than 2^28.                                                            | uint32 | int        | int         | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64   | Always eight bytes. More efficient than uint64 if values are often greater than 2^56.                                                           | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum                         |
| <a name="sfixed32" /> sfixed32 | Always four bytes.                                                                                                                              | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes.                                                                                                                             | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="bool" /> bool         |                                                                                                                                                 | bool   | boolean    | boolean     | bool    | bool       | boolean        | TrueClass/FalseClass           |
| <a name="string" /> string     | A string must always contain UTF-8 encoded or 7-bit ASCII text.                                                                                 | string | String     | str/unicode | string  | string     | string         | String (UTF-8)                 |
| <a name="bytes" /> bytes       | May contain any arbitrary sequence of bytes.                                                                                                    | string | ByteString | str         | []byte  | ByteString | string         | String (ASCII-8BIT)            |
