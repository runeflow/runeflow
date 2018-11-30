# Metadata
PACKAGE = runeflow
VERSION ?= 0.0.0
DESCRIPTION = Runeflow helps you manage and monitor servers
ARCH ?= amd64
MAINTAINER = Ben Burwell

# Auto-generated metadata
DATE = $(shell date -u +%Y-%m-%dT%H:%M:%S)
GIT_REVISION=$(shell git rev-parse HEAD)

# Build environment
GO_SOURCES = $(shell find . -name '*.go')
RELEASE_SOURCES = $(shell find release -type f)
DIST = dist
BUILD = build

# go uses 386, but dpkg expects i386
DEB_ARCH = $(ARCH)
ifeq ($(DEB_ARCH),386)
	DEB_ARCH = i386
endif

DEB_NAME = $(PACKAGE)_$(VERSION)_$(DEB_ARCH).deb

# handle annoying differences between BSD and GNU sed
UNAME_S = $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	SED = sed -i''
else
	SED = sed -i ''
endif

# fill in vars in binary
LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.buildDate=$(DATE) -X main.gitRevision=$(GIT_REVISION)"

$(DIST)/$(DEB_NAME): $(BUILD)
	mkdir -p "$(DIST)"
	dpkg-deb --build $(BUILD)/$(PACKAGE) "$@"

$(BUILD)/$(PACKAGE)/usr/bin/$(PACKAGE): $(GO_SOURCES)
	mkdir -p $(shell dirname "$@")
	GOOS=linux GOARCH=$(ARCH) go build $(LDFLAGS) -o "$@" ./cli

$(BUILD): $(RELEASE_SOURCES)
	cp -R release "$@"
	find $(BUILD) -type f -exec $(SED) "s/__PACKAGE__/$(PACKAGE)/g" {} \;
	find $(BUILD) -type f -exec $(SED) "s/__VERSION__/$(VERSION)/g" {} \;
	find $(BUILD) -type f -exec $(SED) "s/__DESCRIPTION__/$(DESCRIPTION)/g" {} \;
	find $(BUILD) -type f -exec $(SED) "s/__DEB_ARCH__/$(DEB_ARCH)/g" {} \;
	find $(BUILD) -type f -exec $(SED) "s/__MAINTAINER__/$(MAINTAINER)/g" {} \;
	make $(BUILD)/$(PACKAGE)/usr/bin/$(PACKAGE)

.PHONY: clean
clean:
	rm -rf "$(BUILD)" "$(DIST)"

