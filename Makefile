.DEFAULT_GOAL := build

.PHONY:fmt
fmt:
	go fmt ./...

.PHONY:build
build: fmt
	go build -o bin/ helloweb.go