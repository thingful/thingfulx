# Simple Makefile for building and testing thingfulx

PACKAGES := $(shell go list ./... | grep -v /vendor/ )
LINTER = golangci-lint
ARTEFACT_DIR = ./build
GOCMD = go
GOTEST = go test -v
GOLINT = $(LINTER)
GOCOVER = go tool cover
GOGET = $(GOCMD) get -u

default: full

.PHONY: setup
setup:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(shell go env GOPATH)/bin v1.21.0
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: test
test:
	mkdir -p $(ARTEFACT_DIR)
	$(GOTEST) -coverprofile=$(ARTEFACT_DIR)/coverage.out ./...

.PHONY: coverage
coverage: test
	$(GOCOVER) -func=$(ARTEFACT_DIR)/coverage.out
	$(GOCOVER) -html=$(ARTEFACT_DIR)/coverage.out -o $(ARTEFACT_DIR)/coverage.html

.PHONY: generate
generate:
	protoc --go_out=paths=source_relative:. channels.proto

.PHONY: lint
lint:
	$(GOLINT) run

.PHONY: clean
clean:
	rm -rf $(ARTEFACT_DIR)

full: test coverage lint
