# User Module Specification

## Abstract

This module provides functionality to add Sybil-resistance for other modules to use.

## Inspiration from Bitcoin

The simplest and the safest way is to do what Bitcoin does in its protocol, which has been battle-tested over a decade. Some highlights of the Bitcoin protocol:

- Each block refers to a previous block, therefore, forming a chain of blocks. In other words, except for the first block, a block is always required to create a new block.
- To create a new block, the miner has to do some work (i.e. guess the nonce).

Based on the above facts, the inspiration has been derived to design Sybil-resistance mechanism like the following:

- Except for the first set of users (Genesis users), a new user can only start earning reward once vouched for by a previous user.
- To earn the right to vouch (RtV) a user, the user must do some work (i.e. earn a minimum of x reward).

Note: If multiple users vouch for the same new user, only the first voucher's vouch would be considered, and the others' vouches would be left unused.

## State

### Parameters

Each of these parameters can be voted on by governance, but will be assigned sane defaults for genesis.

```go
type Params struct {
    // minimum rewards to be earned to unlock the right to vouch
    // it is an array because on networks like Reddit's,
    // the threshold can be defined as a combination of multiple coins
    ThresholdAmount     sdk.Coins
    
    // number of vouches made available to the user,
    // when they reach the threshold amount
    VouchCount          uint32
}
```

### Base Types

```go
type Vouch struct {
    Voucher         sdk.AccAddress
    Vouched         sdk.AccAddress
    Comment         string // can be used by the voucher to add some notes
}
```

### Stores

_Stores are KVStores in the multi-store. The key to find the store is the first parameter in the list_.

We will use one KVStore `vouches` to store the following two mappings:

- A mapping from `0x01|voucherAddress` to `Vouch`.
- A mapping from `0x02|vouchedAddress` to `Vouch`.

## Messages

### MsgVouch

Vouching a user is registered on-chain with `MsgVouch`. Vouched user starts earning reward as soon as the vouch is registered on the chain.

```go
type MsgVouch struct {
	Voucher       sdk.Address
	Vouched       sdk.Address
	Comment       string // optional
}
```

Vouches are persisted at two places with the keys `0x01|voucherAddress` and `0x02|vouchedAddress`. The first key is used to validated that the voucher cannot vouch more users that they can, while the second key is used to calculate which users should earn the reward during curating's `EndBlock`.

### Handlers

#### MsgVouch

| Type  | Attribute Key   | Attribute Value  |
| ----- | --------------- | ---------------- |
| vouch | voucher_address | {voucherAddress} |
| vouch | vouched_address | {vouchedAddress} |
