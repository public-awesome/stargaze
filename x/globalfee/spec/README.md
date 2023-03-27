# Global Fee

## Abstract

This module enables Stargaze to enforce gas fees on all operations unless explicitly configured to be zero-gas.

Examples of explicit configuration:
1. Allow Code ID 1 to use free gas on all operations
2. Allow Code ID 2 to use free gas on only `mint` operation
3. Allow contract starsblah123blah456 to use free gas on all operations
4. Allow contract starsblah123blah987 to use free gas on `list, unlist` operations

## Concepts

A cosmos-sdk blockchain allows the network of validators to set minimum gas prices as part of the node configuration. It is set in the `config/app.toml` configuration file.

x/globalfee module allows the chain to enforce gas fees on transactions bypassing the validator configuration which might allow for zero gas. 

Additionally, the module also implements a way for the network to set a few operations on specific contracts which can still use free gas (irrespective of validator configuration). This configuration can be updated either via on-chain governance or via whitelisted addresses. 

## Contents

1. [State](./01_state.md)
2. [Messages](./02_messages.md)
3. [Ante Handlers](./03_ante_handlers.md)
4. [Events](./04_events.md)
5. [Params](./05_params.md)
6. [Client](./06_client.md)