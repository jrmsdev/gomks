.PHONY: default
default: build

.PHONY: build
build:
	@mkdir -p ./_build/cmd
	@go build -v -mod vendor -i -o ./_build/cmd/mks.bin ./cmd/mks

.PHONY: clean
clean:
	@go clean -mod vendor -i -cache ./...
	@rm -vf ./mks
	@rm -vrf ./_build ./_testing

.PHONY: distclean
distclean:
	@go clean -mod vendor -i -cache -modcache -testcache ./...

.PHONY: test
test: build
	@go test -mod vendor ./...

.PHONY: testclean
testclean:
	@go clean -mod vendor -testcache ./...

.PHONY: cover
cover: build
	@mkdir -p ./_testing
	@go test -mod vendor -coverprofile ./_testing/cover.out ./...
	@go tool cover -html ./_testing/cover.out -o ./_testing/coverage.html

.PHONY: vendor
vendor:
	@go mod vendor
	@go mod tidy

.PHONY: fmt
fmt:
	@gofmt -w -l -s .

.PHONY: demo
demo: build
	@./_build/cmd/mks.bin ./demo/build.mks
