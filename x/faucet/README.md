# Faucet Module

This module will enable mint function. Every address can mint 100(bonded tokens) in every 24 hours by sending a mint message. 

For security consideration, you can add this module to your project as you want, but this module would *NOT* be active by default. unless you active it manually by adding `"-tags faucet"` when you build or install. 

这个水龙头模块提供铸币功能，每一个地址都可以发送mint消息为自己铸造一定数量的币，时间间隔为24小时。
出于安全考虑，你可以随意将本模块加入到你的项目代码中，但是默认是不会生效的，除非在编译的时候加上`-tags faucet`手动激活这个模块

## Developer Tutorial

Step 1: Import to your app.go
```go
import (
	"github.com/cosmos/modules/incubator/faucet"
)
```

Step 2: Declare faucet module and permission in app.go
```go
ModuleBasics = module.NewBasicManager(
    ..., // the official basic modules

    faucet.AppModule{},  // add faucet module
)
// account permissions
maccPerms = map[string][]string{
    staking.BondedPoolName:    {supply.Burner, supply.Staking},
    staking.NotBondedPoolName: {supply.Burner, supply.Staking},
    faucet.ModuleName:          {supply.Minter}, // add permissions for faucet
}
	
type nameServiceApp struct {
    *bam.BaseApp
    cdc *codec.Codec

    // Other Keepers ... ...
    
    // Declare faucet keeper here
    faucetKeeper faucet.Keeper

    // Module Manager
    mm *module.Manager

    // simulation manager
    sm *module.SimulationManager
}
```

Step 3: Initialize faucet keeper and faucet module in func NewNameserviceApp() in app.go
```go
keys := sdk.NewKVStoreKeys(
		bam.MainStoreKey,
		auth.StoreKey,
		staking.StoreKey,
		supply.StoreKey,
		distr.StoreKey,
		slashing.StoreKey,
		params.StoreKey,
		faucet.StoreKey) // add faucet key

... // some stuff in between

app.faucetKeeper = faucet.NewKeeper(
    app.supplyKeeper, 
    app.stakingKeeper, 
    10 * 1000000,  // amount for mint
    24 * time.Hour, // rate limit by time
    keys[faucet.StoreKey], 
    app.cdc,)

app.mm = module.NewManager(
    ..., // other modules
    
    faucet.NewAppModule(app.faucetKeeper), // add faucet module
    
)
```

Step 4: Enable faucet in [Makefile](Makefile_Sample)
```
installWithFaucet: go.sum
		go install -mod=readonly $(BUILD_FLAGS) -tags faucet ./cmd/nsd
		go install -mod=readonly $(BUILD_FLAGS) -tags faucet ./cmd/nscli
```

Step 5: Build your app
```
make installWithFaucet
```

Step 6: Initialize faucet and publish to blockchain.

Create an account with default password "12345678", and publish it to blockchain, therefore others are able to load this account to mint coins.

```
nscli tx faucet publish --from faucet --chain-id test
{
  "chain_id": "test",
  "account_number": "3",
  "sequence": "2",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "faucet/FaucetKey",
      "value": {
        "Sender": "cosmos14pkakt8apdm0e49tzp6gy3lwe8u04ajched5qm",
        "Armor": "-----BEGIN TENDERMINT KEY INFO-----\ntype: Info\nversion: 0.0.0\n\nZA2tFT0KBHBpbmcSJuta6YchA8yOxBXjUwLzUeBxeECHpLU2GwILK/7OVbMF6uiX\nl/PNGiXhsPebIP3XBLEM0VlX6/whk4LtlqqvYLOduCLGh1yS0OE4SQFWIglzZWNw\nMjU2azE=\n=WAST\n-----END TENDERMINT KEY INFO-----"
      }
    }
  ],
  "memo": ""
}
```

Congratulations, you are ready to mint 

## Usage / 用法

1: Initialize faucet 

``` 
nscli tx faucet init --chain-id test
The faucet has been loaded successfully.
```

2: Mint coins.

```
iMac:~ liangping$ nscli tx faucet mintfor cosmos17l3gw079cn5x9d3pqa0jk0xhrw2mt358xvw555 --from faucet --chain-id test -y
{
  "height": "0",
  "txhash": "40F2AB8AD75B39532622302A71CB84523847D2E43D36B185E0CE65CE60208AB0",
  "raw_log": "[]"
}
iMac:~ liangping$ nscli query account cosmos17l3gw079cn5x9d3pqa0jk0xhrw2mt358xvw555 --chain-id test
{
  "type": "cosmos-sdk/Account",
  "value": {
    "address": "cosmos17l3gw079cn5x9d3pqa0jk0xhrw2mt358xvw555",
    "coins": [
      {
        "denom": "stake",
        "amount": "100000000"
      }
    ],
    "public_key": "",
    "account_number": 7,
    "sequence": 0
  }
}

```

Also, you are able to mint for yourself if your address has actived/existed on blockchain.
```
$ nscli tx faucet mint --from YOUR_OTHER_ACCOUNT --chain-id test -y
{
  "height": "0",
  "txhash": "40F2AB8AD75B39532622302A71CB84523847D2E43D36B185E0CE65CE60208AB0",
  "raw_log": "[]"
}
```

## Compatible Version

 cosmos-sdk v0.38.0 or above

## Contact Us

Author: liangping from [Ping.pub](https://ping.pub)

18786721@qq.com

If you like this module, welcome to delegate to Ping.pub on [Cosmoshub](https://cosmos.ping.pub), [IRISHub](https://iris.ping.pub), [KAVA](https://kava.ping.pub).
