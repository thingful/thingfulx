# Simple Makefile for building and testing thingfulx

ARTEFACT_DIR = ./build
GOCMD = go
GODEP = godep
GOTEST = $(GODEP) go test
GOLINT = golint
GOCOVER = $(GODEP) go tool cover
GOGET = $(GOCMD) get -u
GORESTORE = $(GODEP) go restore

default: test

setup:
	$(GOGET) github.com/tools/godep
	$(GOGET) github.com/golang/lint/golint

restore:
	$(GORESTORE)

test:
	mkdir -p $(ARTEFACT_DIR)
	$(GOTEST) -coverprofile=$(ARTEFACT_DIR)/cover.out ./...

coverage: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -func=$(ARTEFACT_DIR)/cover.out

html: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -html=$(ARTEFACT_DIR)/cover.out -o $(ARTEFACT_DIR)/coverage.html

lint:
	$(GOLINT) ./...

clean:
	rm -rf $(ARTEFACT_DIR)

full: test coverage html

.PHONY: setup restore test lint coverage html clean full
