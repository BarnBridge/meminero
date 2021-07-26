VERSION := "$(shell git describe --abbrev=0 --tags 2> /dev/null || echo 'v0.0.0')+$(shell git rev-parse --short HEAD)"

build:
	go build -ldflags "-X main.buildVersion=$(VERSION)"

reset: build
	./meminero reset --force
	./meminero migrate

sync-all: build
	./meminero sync --syncer.datasets labels,tokens,monitored-accounts,monitored-erc20,smart-exposure-pools,smart-yield-pools,smart-yield-reward-pools
