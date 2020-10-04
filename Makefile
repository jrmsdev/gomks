.PHONY: default
default: build

.PHONY: build
build:
	@go build -v -i ./cmd/mks

.PHONY: clean
clean:
	@go clean -i ./...
	@rm -vf ./mks
