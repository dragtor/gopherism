#!/bin/sh

test: 
	go test ./...

formatter: 
	go fmt cmd/main.go
	
build: formatter clean  test
	go build -o bin/urlshortner cmd/main.go

clean: 
	rm -rf bin/urlshortner
