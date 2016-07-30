# Simple Makefile for building and testing thingfulx

ARTEFACT_DIR = ./build
GOCMD = go
GOTEST = go test
GOLINT = gometalinter --deadline=10s --disable-all --enable=errcheck varcheck
GOCOVER = go tool cover
GOGET = $(GOCMD) get -u

default: full

setup:
	$(GOGET) github.com/alecthomas/gometalinter
	$(GOPATH)/bin/gometalinter --install --update

test:
	mkdir -p $(ARTEFACT_DIR)
	$(GOTEST) -coverprofile=$(ARTEFACT_DIR)/cover.out .

coverage: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -func=$(ARTEFACT_DIR)/cover.out

html: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -html=$(ARTEFACT_DIR)/cover.out -o $(ARTEFACT_DIR)/coverage.html

lint:
	$(GOLINT) .

clean:
	rm -rf $(ARTEFACT_DIR)

full: test coverage html lint

.PHONY: setup test lint coverage html clean full
