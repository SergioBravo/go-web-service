.PHONY: build
build:
	go build -v ./cmd/webservice

.DEFAULT_GOAL: build