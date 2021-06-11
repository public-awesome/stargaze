#!/bin/bash
npx ts-node \
contract-deployer.ts \
--cosmos-node="http://localhost:26657" \
--eth-node="https://eth-goerli.alchemyapi.io/v2/SMHE5NpwoSbTprgY17Q11IjXYuWXCMPw" \
--eth-privkey="0x7432b3cfe8e764c3ca66b0ea291f9ac656e19464e65e3e4d7abce7688bde1fe0" \
--contract=artifacts/contracts/Gravity.sol/Gravity.json
#--test-mode=true
