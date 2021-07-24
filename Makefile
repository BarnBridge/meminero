VERSION := "$(shell git describe --abbrev=0 --tags 2> /dev/null || echo 'v0.0.0')+$(shell git rev-parse --short HEAD)"

build:
	go build -ldflags "-X main.buildVersion=$(VERSION)"

reset: build
	./meminero reset --force
	./meminero migrate
