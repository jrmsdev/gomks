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

.PHONY: demo
demo: build
	@./demo/build.sh
