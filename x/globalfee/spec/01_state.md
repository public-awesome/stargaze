# State

This section describes all stored state by the module objects and their storage keys. 


### CodeAuthorization
```protobuf
message CodeAuthorization {
  uint64 code_id = 1;
  repeated string methods = 2;
}
```
CodeAuthorization is used to store the configuration for Code IDs which can have zero gas fee operations. The methods are individual msg operations which benefit from zero fees. e.g `mint`, `burn`. This configuration would allow all contracts instantiated from the given code ID to benefit zero gas fee.

> **Note**
>
> `'*'` value can be used in the configuration to enable all methods of the contracts with given Code ID to have zero gas fee

Storage keys:

* CodeAuthorization: 0x00 | CodeID -> ProtocolBuffer(CodeAuthorization)

This state can be updated via [MsgSetCodeAuthorization](./02_messages.md#msgsetcodeauthorization) & [MsgRemoveCodeAuthorization](./02_messages.md#msgremovecodeauthorization) by whitelisted addresses or via governance

### ContractAuthorization
```protobuf
message ContractAuthorization {
  string contract_address = 1;
  repeated string methods = 2;
}
```
ContractAuthorization is used to store the configuration for contract addresses which can have zero gas fee operations. The methods are individual msg operations which benefit from zero gas. e.g `mint`, `burn`. 

> **Note**
>
> `'*'` value can be used in the configuration to enable all methods of the contract to have zero gas fee

Storage keys:

* ContractAuthorization: 0x01 | ContractAddress -> ProtocolBuffer(ContractAuthorization)

This state can be updated via [MsgSetContractAuthorization](./02_messages.md#msgsetcontractauthorization) & [MsgRemoveContractAuthorization](./02_messages.md#msgremovecontractauthorization) by whitelisted addresses or via governance