# Cron

## Abstract

This module enables smart contracts on chain to receive callbacks at the beginning and end of each block. Useful for scheduling actions which need to happen every block.


## Concepts

### Priviledged contracts
An existing contract instantiated on chain which has been elevated to the status of priviledged contract via on-chain governance proposal. 
Privileged contracts receive some extra superpowers where they can receive callbacks from the [`BeginBlocker & EndBlocker`](https://docs.cosmos.network/main/building-modules/beginblock-endblock.html) into their [`sudo`](https://book.cosmwasm.com/basics/entry-points.html?highlight=sudo#entry-points) entrypoint. More details on how to use this in [here](#how-to-use-in-cw-contract)

## State

The module state is quite simple. It stores the bech32 address of all the Cosmwasm smart contracts that have been elevated via governance. Once the contract priviledge is demoted via governance, it is removed from state.

Storage keys:

- PriviledgedContract: `0x00 | contractAddress -> []byte{1}`


## Begin Blockers

In the ABCI beginblock, the module fetches the list of all priviledged contracts and loops through them and calls the `sudo` entry point with the `BeginBlocker` msg. None of the `abci.RequestBeginBlock` values are passed in.

## End Blockers

In the ABCI endblock, the module fetches the list of all priviledged contracts and loops through them and calls the `sudo` entry point with the `EndBlocker` msg.

## Events

The module emits the following events:

| Source type | Source name                |
| ----------- | -------------------------- |
| Keeper      | `SetContractPriviledge`    |
| Keeper      | `UnsetContractPriviledge`  |

## Client

### CLI - Query

#### **list-privileged**

```
starsd q cron list-privileged
```

List all privileged contract addresses in bech32 format

### CLI - Tx

#### **promote-to-privilege-contract** || **promote-contract**

```
starsd tx gov submit-proposal promote-contract {contractAddr} --title {proposalTitle} --deposit {depositAmount}
starsd tx gov submit-proposal promote-contract stars19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseues57n420 --title "Promote Contract Proposal" --deposit 1000ustars
```

Creates a governance proposal which on passing will add the given contract address to the priviledged contract list

#### **demote-from-privilege-contract** || **demote-contract**

```
starsd tx gov submit-proposal demote-contract {contractAddr} --title {proposalTitle} --deposit {depositAmount}
starsd tx gov submit-proposal demote-contract  stars19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseues57n420 --title "Demote Contract Proposal" --deposit 1000ustars
```

Creates a governance proposal which on passing will remove the given contract address from the priviledged contract list

## How to use in CW contract

To use the `BeginBlocker` & `EndBlocker` callback from this module, you need to add the following msg type to your contract msgs. 

```rust
// msg.rs

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum SudoMsg {
    BeginBlock {},
    EndBlock {},
}
```

```rust
// contract.rs

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn sudo(deps: DepsMut, env: Env, msg: SudoMsg) -> Result<Response, ContractError> {
    match msg {
        SudoMsg::BeginBlock {} => !unimplemented(),
        SudoMsg::EndBlock {} => !unimplemented(),
    }
}
```
Ensure you have a sudo entrypoint in your contract and add your code to be called by the `SudoMsg::BeginBlock` & `SudoMsg::EndBlock`. Both endpoints are hit for every privileged contract in the block lifecycle, so **ensure both are implemented**.

This endpoint is hit at the beginning and end of every block by the x/cron module. The `deps` and `env` values are set as expected. It is not possible to receive any input in the `BeginBlock` & `EndBlock` call and as such the contract must maintain any relevant state it needs to use. 


## Attribution
Thanks to Confio and the Tgrade open source project for providing a base for x/cron where most of it's bits were extracted and adapted to only do begin_blocker & end_blocker callbacks.

Found more about Tgrade's twasm module:

[x/twasm](https://github.com/confio/tgrade/tree/main/x/twasm)
