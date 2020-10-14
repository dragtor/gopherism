#!/bin/sh

test: 
	go test ./...

formatter: 
	go fmt cmt/main.go
	
build: clean  test
	go build -o bin/quiz cmd/main.go

clean: 
	rm -rf bin/quiz
