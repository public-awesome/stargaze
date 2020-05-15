GOARCH=amd64 GOOS=linux go build -o bin/staked github.com/rocket-protocol/stakebird/cmd/staked
GOARCH=amd64 GOOS=linux  go build -o bin/stakecli github.com/rocket-protocol/stakebird/cmd/stakecli
