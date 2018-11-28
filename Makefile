BINARY=runeflow
CLI_SOURCES=$(shell find ./cli -name '*.go')
BUILD_DIR=build

$(BUILD_DIR)/$(BINARY): $(CLI_SOURCES)
	GOOS=linux go build -o "$@" ./cli

