PACKAGE = runeflow
MAINTAINER = Ben Burwell
CLI_SOURCES = $(shell find ./cli -name '*.go')
TARGET = target
DIST = dist
VERSION ?= 0.0.0
DATE = $(shell date -u +%Y-%m-%dT%H:%M:%S)
GIT_REVISION=$(shell git rev-parse HEAD)
DESCRIPTION = Runeflow helps you manage and monitor servers
ARCH ?= amd64

# handle annoying differences between BSD and GNU sed
UNAME_S = $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	SED = sed -i''
else
	SED = sed -i ''
endif

LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.buildDate=$(DATE) -X main.gitRevision=$(GIT_REVISION)"

$(DIST)/$(PACKAGE)_$(VERSION)_$(ARCH).deb: $(TARGET)/$(PACKAGE)_$(VERSION)_$(ARCH)
	rm -rf $(TARGET)/$(PACKAGE)
	cp -R release/$(PACKAGE) $(TARGET)/$(PACKAGE)
	find $(TARGET)/$(PACKAGE) -type f -exec $(SED) "s/__PACKAGE__/$(PACKAGE)/g" {} \;
	find $(TARGET)/$(PACKAGE) -type f -exec $(SED) "s/__VERSION__/$(VERSION)/g" {} \;
	find $(TARGET)/$(PACKAGE) -type f -exec $(SED) "s/__DESCRIPTION__/$(DESCRIPTION)/g" {} \;
	find $(TARGET)/$(PACKAGE) -type f -exec $(SED) "s/__ARCH__/$(ARCH)/g" {} \;
	find $(TARGET)/$(PACKAGE) -type f -exec $(SED) "s/__MAINTAINER__/$(MAINTAINER)/g" {} \;
	cp "$<" $(TARGET)/$(PACKAGE)/usr/bin/$(PACKAGE)
	mkdir -p "$(DIST)"
	dpkg-deb --build $(TARGET)/$(PACKAGE) "$@"

$(TARGET)/$(PACKAGE)_$(VERSION)_$(ARCH): $(CLI_SOURCES)
	GOOS=linux GOARCH=$(ARCH) go build $(LDFLAGS) -o "$@" ./cli

.PHONY: clean
clean:
	rm -rf "$(TARGET)"

