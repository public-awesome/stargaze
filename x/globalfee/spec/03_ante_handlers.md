# AnteHandlers

This section describes the module ante handlers

## FeeDecorator

FeeDecorator parses every msg in the transactions.

If the msg it of type `authz.MsgExec`, we unpack its wrapped msg and proceed.

If the msg is of type `wasmd.MsgExecuteContract`, we decode the msg to identify the following:
* the contract address
* the method being called
* the code id for the given contract address

If all the msgs in the tx are eligible for zero fees, then we allow the msg to `next()` irrespective of any fees provided.
In all other cases, we take the required fees value from node's local minimum gas prices and check if the fees provided are sufficient.

//todo
handle cases where tx has multiple msgs some of which are eligible to be zero fees while others are not