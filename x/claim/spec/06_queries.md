<!--
order: 6
-->

# Queries

## GRPC queries

Claim module provides below GRPC queries to query claim status

```protobuf
service Query {
  rpc ModuleAccountBalance(QueryModuleAccountBalanceRequest) returns (QueryModuleAccountBalanceResponse) {}
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse) {}
  rpc ClaimRecord(QueryClaimRecordRequest) returns (QueryClaimRecordResponse) {}
  rpc ClaimableForAction(QueryClaimableForActionRequest) returns (QueryClaimableForActionResponse) {}
  rpc TotalClaimable(QueryTotalClaimableRequest) returns (QueryTotalClaimableResponse) {}
}
```

## CLI commands

For the following commands, you can change `$(stargazed keys show -a {your key name})` with the address directly.

Query the claim record for a given address

```sh
stargazed query claim claim-record $(stargazed keys show -a {your key name})
```

Query the claimable amount that would be earned if a specific action is completed right now.

```sh

stargazed query claim claimable-for-action $(stargazed keys show -a {your key name}) ActionAddLiquidity
```

Query the total claimable amount that would be earned if all remaining actions were completed right now.
Note that even if the decay process hasn't begun yet, this is not always _exactly_ the same as `InitialClaimableAmount`, due to rounding errors.

```sh
stargazed query claim total-claimable $(stargazed keys show -a {your key name}) ActionAddLiquidity
```
