
GO ?= go

.PHONY: build test vet bench run clean

build:
	$(GO) build ./...

test:
	$(GO) build ./...
	$(GO) test ./... -race -count=1

vet:
	$(GO) vet ./...

bench:
	$(GO) test ./... -bench=. -benchmem -run=^$$

run:
	$(GO) run ./cmd/marlind -config marlin.yaml

clean:
	rm -rf bin coverage.out
