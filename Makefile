.PHONY: watch build tidy deps gogen test

build:
	go build ./...

test: build
	go test ./...

tidy:
	go fmt ./...
	go mod tidy

watch:
	modd

deps:
	go get -u github.com/tobgu/peds/cmd/peds

gogen:
	go generate ./...
