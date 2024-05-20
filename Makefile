.DEFAULT_GOAL := run

fmt:
	gofmt -w .
.PHONY:fmt

test: fmt
	go test ./... -count=1
.PHONY: test

build: test
	go build 
.PHONY:build

run: fmt 
	go build && ./pokedex-repl
.PHONY:run