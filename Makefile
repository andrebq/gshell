.PHONY: watch build tidy deps gogen

build:
	go build ./...

tidy:
	go fmt ./...
	go mod tidy

watch:
	modd

deps:
	go get -u github.com/tobgu/peds/cmd/peds

gogen:
	go generate ./...
