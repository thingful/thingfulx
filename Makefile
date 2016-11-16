# Simple Makefile for building and testing thingfulx

LINTER = gometalinter
ARTEFACT_DIR = ./build
GOCMD = go
GOTEST = go test -v
GOLINT = $(LINTER) --deadline=30s --vendor --debug
GOCOVER = go tool cover
GOGET = $(GOCMD) get -u

default: full

.PHONY: setup
setup:
	$(GOGET) github.com/alecthomas/$(LINTER)
	$(LINTER) --install

.PHONY: test
test:
	mkdir -p $(ARTEFACT_DIR)
	$(GOTEST) -coverprofile=$(ARTEFACT_DIR)/cover.out .

.PHONY: coverage
coverage: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -func=$(ARTEFACT_DIR)/cover.out

.PHONY: html
html: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -html=$(ARTEFACT_DIR)/cover.out -o $(ARTEFACT_DIR)/coverage.html

.PHONY: lint
lint:
	$(GOLINT) ./...

.PHONY: clean
clean:
	rm -rf $(ARTEFACT_DIR)

full: test coverage html lint
