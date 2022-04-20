.DEFAULT_GOAL := default
.PHONY: all

test:
	go test ./...

default: test