#!/bin/sh

test:
	go test ./...

formatter:
	go fmt ./...

build: formatter clean test
	go build -o bin/cyoa cmd/main.go

clean:
	rm -rf bin/cyoa
