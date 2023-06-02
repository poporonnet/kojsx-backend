SOURCE_FILES:=$(shell find . -type f -name '*.go' -print)

VERSION:=$(shell cat ./VERSION)
REVISION:=$(shell git rev-parse --short HEAD)

.PHONY: build
build: $(SOURCE_FILES)
	@CGO_ENABLED=0 go build -o kojs -ldflags "-s -w -X main.VERSION=${VERSION} -X main.REVISION=${REVISION}"
