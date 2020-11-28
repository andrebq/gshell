.PHONY: generate

generate:
	docker run --rm -v $(CURDIR):/workdir -w /workdir andrebq/antlr4:latest antlr -Dlanguage=Go -o . GShell.g4
