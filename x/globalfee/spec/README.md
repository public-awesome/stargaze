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

x/globalfee module exploses a param called `MinimumGasPrices`. The param can hold multiple denoms and the fees must be paid in at least one denom from the configured param list and the sent amount per unit gas must be greater than or equal to the corresponding amount in the param list. e.g `[1stars, 1atom]`

Between the minimum-gas-prices set by the validator and the MinimumGasPrices configured on chain, if contain the same denom, the higher of the two is used. Therefore, the validators cannot undercut the configured param but can set higher values.

> **Note**
> The denoms in `min-gas-prices` that are not present in `MinimumGasPrices` are ignored.

x/globalfee module implements a way for the network to set a few operations on specific cosmwasm contracts which can use zero gas fees (irrespective of validator configuration or MinimumGasPrices value).

The transaction benefits from zero gas fee operations only when all the msgs in the tranaction have been authorized. Else, typical fee has to be paid for the transaction. 

This configuration of code/contract authorizations can be updated either via on-chain governance or via whitelisted addresses. 

The whitelisted addresses are set via governance as a module param.



## Contents

1. [State](./01_state.md)
2. [Messages](./02_messages.md)
3. [Ante Handlers](./03_ante_handlers.md)
4. [Params](./04_params.md)
5. [Client](./05_client.md)