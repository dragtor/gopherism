#!/bin/sh

test: formatter
	go test ./...

formatter: 
	go fmt ./...

clean:
	rm -rf bin/task
	rm -rf /tmp/task.db

build: clean formatter test
	go build -o bin/task main.go
