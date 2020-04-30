# Stake module specification

The stake module performs time-constrained validator delegations.

## Dependencies

* x/bank
* x/staking
    - `UnbondingPeriod` = 3 days
    - `MaxEntries` = 7+

## State

### DelegationQueue (FIFO)

* DelegationQueue: 0x01 | format(expire_time) | vendor_id | post_id -> Delegation

```go
type Delegation struct {
    DelegatorAddr sdk.AccAddress
    ValidatorAddr sdk.ValAddress
    Amount        sdk.Dec
}
```

## Messages

### MsgDelegate

* Call `Delegate()` in the staking module
* Add a new `Delegation` to the end of `DelegationQueue`

## End-Block

* Check `expire_time` in `DelegationQueue` against the current block time
* Process all delegations that have an `expire_time` that exceeds the current block time
    - Undelegate
    - Distribute rewards

## Events

### EndBlocker

| Type              | Attribute Key         | Attribute Value           |
| ----------------- | --------------------- | ------------------------- |
| complete_stake    | amount                | {Amount}                  |
| complete_stake    | delegator             | {delegatorAddress}        |

