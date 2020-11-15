.PHONY: watch

deps:
	go get -u github.com/pointlander/peg

build:
	go build ./...

generate:
	peg -switch -inline internal/grammar/gshell.peg

watch:
	modd
