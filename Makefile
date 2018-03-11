# General
WORKDIR = $(PWD)

# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test -v

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

test:
	@cd $(WORKDIR); \
    $(GOTEST) ./...

check:
	@for gofile in $$(find . -path './vendor' -o -path './src' -prune -o -type f -iname '*.go' -print); do \
			gofmt -w $$gofile; \
			golint $$gofile; \
			go vet -v $$gofile; \
    done