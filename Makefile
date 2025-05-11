N = main

######
# GO #
######

all:
	go run $(N).go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: build
build:
	go build $(N).go

.PHONY: build-run
build-run:
	go build $(N).go && ./$(N)

#######
# SQLC #
#######

generate:
	sqlc generate