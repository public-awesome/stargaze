# State

This section describes all stored state by the module objects and their storage keys. 

## Params

```protobuf
message Params {
  repeated string privileged_address = 1;
}
```

As stored in the global params store, it stores the list of addresses who have the ability to modify the Global Fee configuration without having to go via governance.

### CodeAuthorization
```protobuf
message CodeAuthorization {
  uint64 code_id = 1;
  repeated string methods = 2;
}
```
CodeAuthorization is used to store the configuration for code Ids which can have zero gas operations. The methods are individual msg operations which benefit from zero gas. e.g `mint`, `burn`. This configuration would allow all contracts instantiated from the given code Id to be gas free.

`*` value can be used in the configuration to enable all methods of the contracts with given code Id to be gas free

Storage keys:

* CodeAuthorization: 0x00 | CodeId -> ProtocolBuffer(CodeAuthorization)

### ContractAuthorization
```protobuf
message ContractAuthorization {
  string contract_address = 1;
  repeated string methods = 2;
}
```
ContractAuthorization is used to store the configuration for contract addresses which can have zero gas operations. The methods are individual msg operations which benefit from zero gas. e.g `mint`, `burn`. 

`*` value can be used in the configuration to enable all methods of the contract to be gas free

Storage keys:

* ContractAuthorization: 0x01 | ContractAddress -> ProtocolBuffer(ContractAuthorization)