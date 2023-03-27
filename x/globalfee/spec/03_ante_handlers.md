# AnteHandlers

This section describes the module ante handlers

## FeeDecorator

FeeDecorator parses every msg in the transactions.
If the msg is of type `MsgExecuteContract`, we decode the msg to identify the following:
* the contract address
* the method being called
* the code id for the given contract address

If we find any authorizations for the given values, we allow the msg to `next()` even if there are no fees in the tx.

If the tx contains many msgs, including ones with zero gas authorization, //todo