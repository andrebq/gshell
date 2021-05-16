.PHONY: watch build tidy deps gogen test benchmark

build:
	go build ./...

test: build
	go test ./...

benchmark: test
	go test -bench=. -benchtime=15s ./mailbox

tidy:
	go fmt ./...
	go mod tidy

watch:
	modd

deps:
	go get -u github.com/tobgu/peds/cmd/peds

gogen:
	go generate ./...
