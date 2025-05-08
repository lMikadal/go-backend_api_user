N = main

all:
	go run $(N).go

.PHONY: build
build:
	go build $(N).go

.PHONY: build-run
build-run:
	go build $(N).go && ./$(N)