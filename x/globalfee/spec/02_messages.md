# Messages

This section describes the processing of the module messages

## MsgSetCodeAuthorization

```protobuf
message MsgSetCodeAuthorization {
  string sender_address = 1;
  uint64 code_id = 2;
  repeated string methods = 3;
}
```

A new Code Id authorization is set using the MsgSetCodeAuthorization message

On Success:
* Code authorization is set/updated 

This message is expected to fail if:
* Given Code Id does not exist
* Sender is not part of the whitelist configured in the module params
* Code authorization methods are empty

This data can also be updated via governance

## MsgRemoveCodeAuthorization

```protobuf
message MsgRemoveCodeAuthorization {
  string sender_address = 1;
  uint64 code_id = 2;
}
```

Existing Code Id authorization is deleted using the MsgRemoveCodeAuthorization message

On Success:
* Code authorization is removed

This message is expected to fail if:
* Sender is not part of the whitelist configured in the module params

This data can also be updated via governance

## MsgSetContractAuthorization

```protobuf
message MsgSetContractAuthorization {
  string sender_address = 1;
  string contract_address = 2;
  repeated string methods = 3;
}
```

A new contract authorization is set using the MsgSetContractAuthorization message

On Success:
* Contract authorization is set/updated 

This message is expected to fail if:
* No contract exists for given address
* Sender is not part of the whitelist configured in the module params
* Contract authorization methods are empty

This data can also be updated via governance

## MsgRemoveContractAuthorization

```protobuf
message MsgRemoveContractAuthorization {
  string sender_address = 1;
  string contract_address = 2;
}
```

Existing contract authorization is removed using the MsgRemoveContractAuthorization message

On Success:
* Contract authorization is deleted

This message is expected to fail if:
* Sender is not part of the whitelist configured in the module params

This data can also be updated via governance