.PHONY: watch build tidy

build:
	go build ./...

tidy:
	go fmt ./...
	go mod tidy

watch:
	modd
