# Simple Makefile for building and testing thingfulx
#
ARTEFACT_DIR = ./build
GOCMD = go
GOTEST = $(GOCMD) test
GOLINT = golint
GOCOVER = $(GOCMD) tool cover

default: test

setup:
	$(GOCMD) get -u github.com/golang/lint/golint

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

.PHONY: setup test lint coverage full html clean
