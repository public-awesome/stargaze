# WASM Contract Scripts

## Start chain

```sh
../../startnode.sh
```

## Upload contracts

```sh
BINARY='starsd'
DENOM='ustarx'
CHAIN_ID='localnet-1'
NODE='http://localhost:26657'

bash upload-contracts.sh <key-name>
```
