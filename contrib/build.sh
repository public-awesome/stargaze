GOARCH=amd64 GOOS=linux go build -o bin/rocketd github.com/rocket-protocol/stakebird/cmd/staked
GOARCH=amd64 GOOS=linux  go build -o bin/rocketcli github.com/rocket-protocol/stakebird/cmd/stakecli
