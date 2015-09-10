
GO ?= go

.PHONY: build test vet bench run clean

build:
	$(GO) build ./...

test:
	$(GO) build ./...
	$(GO) test ./... -race -count=1

vet:
	$(GO) vet ./...

