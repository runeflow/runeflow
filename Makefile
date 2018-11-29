BINARY = runeflow
CLI_SOURCES = $(shell find ./cli -name '*.go')
BUILD_DIR = build
ARCH = amd64
VERSION ?= dev

LDFLAGS = "-X main.version=$(VERSION) -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitRevision=$(shell git rev-parse HEAD)"

$(BUILD_DIR)/$(BINARY)_$(VERSION)_$(ARCH): $(CLI_SOURCES)
	git rev-parse HEAD
	GOOS=linux GOARCH="$(ARCH)" go build -ldflags $(LDFLAGS) -o "$@" ./cli

.PHONY: clean
clean:
	rm -rf "$(BUILD_DIR)"
