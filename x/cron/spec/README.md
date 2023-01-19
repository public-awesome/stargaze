# Cron

## Abstract

This module enables smart contracts on chain to recive callbacks at the end of each block. Useful for scheduling actions which need to happen every block.


## Concepts

### Priviledged contracts
An existing contract instantiated on chain which has been elevated to the status of priviledged contract via on-chain governance proposal. 
Priviledged contracts receive some extra superpowers where they can receive callbacks from the [`EndBlocker`](https://docs.cosmos.network/main/building-modules/beginblock-endblock.html) into their [`sudo`](https://book.cosmwasm.com/basics/entry-points.html?highlight=sudo#entry-points) entrypoint

## State

The module state is quite simple. It stores an array of bech32 address of all the Cosmwasm smart contracts that have been elevated via governance. Once the contract priviledge is demoted via governance, it is removed from state.

## End Blockers

In the ABCI endblock, the module fetches the list of all priviledged contracts and loops through them and calls the `sudo` entry point.

## Events

The module emits the following events:

| Source type | Source name                |
| ----------- | -------------------------- |
| Keeper      | `SetContractPriviledge`    |
| Keeper      | `UnsetContractPriviledge`  |

## Client

### CLI - Query

#### **list-privileged**

`starsd q cron list-privileged`

List all privileged contract addresses in bech32 format

## WASM bindings

//todo