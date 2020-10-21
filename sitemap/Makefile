#!/bin/sh

test: formatter
	go test ./...

formatter: 
	go fmt ./...

clean:
	rm -rf bin/sitemap

build: clean formatter test
	go build -o bin/sitemap cmd/main.go
