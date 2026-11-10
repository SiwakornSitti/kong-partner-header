BINARY := partner-header
KONG_BINARY := partner-header-linux
KONG_GOARCH ?= $(shell go env GOARCH)

.PHONY: build build-kong

build:
	go build -o $(BINARY) .

build-kong:
	GOOS=linux GOARCH=$(KONG_GOARCH) CGO_ENABLED=0 go build -o $(KONG_BINARY) .
