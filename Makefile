VERSION := "$(shell git describe --abbrev=0 --tags 2> /dev/null || echo 'v0.0.0')+$(shell git rev-parse --short HEAD)"

meminero: $(shell find . -name '*.go')
	go build -ldflags "-X main.buildVersion=$(VERSION)"

reset: meminero
	./meminero reset --force
	./meminero migrate

sync-mainnet: meminero
	./meminero sync --syncer.network mainnet --syncer.datasets labels,monitored-accounts,monitored-erc20,smart-alpha-pools,smart-alpha-reward-pools,smart-exposure-pools,smart-yield-pools,smart-yield-reward-pools,tokens

sync-kovan: meminero
	./meminero sync --syncer.network kovan --syncer.datasets labels,monitored-accounts,monitored-erc20,smart-alpha-pools,smart-alpha-reward-pools,smart-exposure-pools,smart-yield-pools,smart-yield-reward-pools,tokens

sync-poly: meminero
	./meminero sync --syncer.network polygon --syncer.datasets monitored-erc20,smart-alpha-pools,smart-exposure-pools,smart-yield-pools,smart-yield-reward-pools,tokens

sync-fuji: meminero
	./meminero sync --config config-sync.yml --syncer.network fuji --syncer.datasets smart-alpha-pools,tokens

sync-avalanche: meminero
	./meminero sync --config config-sync-fwd.yml --syncer.network avalanche --syncer.datasets smart-alpha-pools,tokens

sync-bsctestnet: meminero
	./meminero sync --config config-sync.yml --syncer.network bsctestnet --syncer.datasets smart-alpha-pools,tokens

sync-bsc: meminero
	./meminero sync --config config-sync-fwd.yml --syncer.network bsc --syncer.datasets smart-alpha-pools,tokens

sync-arbitrumrinkeby: meminero
	./meminero sync --config config-sync.yml --syncer.network arbitrumrinkeby --syncer.datasets smart-alpha-pools,tokens

sync-arbitrum: meminero
	./meminero sync --config config-sync-fwd.yml --syncer.network arbitrum --syncer.datasets smart-alpha-pools,tokens

sync-optimistickovan: meminero
	./meminero sync --config config-sync.yml --syncer.network optimistickovan --syncer.datasets smart-alpha-pools,tokens

gen: meminero
	./meminero generate-eth-types
