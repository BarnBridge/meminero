VERSION := "$(shell git describe --abbrev=0 --tags 2> /dev/null || echo 'v0.0.0')+$(shell git rev-parse --short HEAD)"

meminero: $(shell find . -name '*.go')
	go build -ldflags "-X main.buildVersion=$(VERSION)"

reset: meminero
	./meminero reset --force
	./meminero migrate

sync-mainnet: meminero
	./meminero sync --syncer.network mainnet --syncer.datasets labels,monitored-accounts,monitored-erc20,smart-alpha-pools,smart-exposure-pools,smart-yield-pools,smart-yield-reward-pools,tokens

sync-kovan: meminero
	./meminero sync --syncer.network kovan --syncer.datasets labels,monitored-accounts,monitored-erc20,smart-alpha-pools,smart-exposure-pools,smart-yield-pools,smart-yield-reward-pools,tokens

sync-poly: meminero
	./meminero sync --syncer.network polygon --syncer.datasets monitored-erc20,smart-alpha-pools,smart-exposure-pools,smart-yield-pools,smart-yield-reward-pools,tokens

sync-fuji: meminero
	./meminero sync --syncer.network fuji --syncer.datasets smart-alpha-pools,tokens

gen: meminero
	./meminero generate-eth-types
