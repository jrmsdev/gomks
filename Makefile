PKG ?= ./...
MOD := github.com/jrmsdev/gomks

.PHONY: default
default: build

.PHONY: build
build: _build/version
	@mkdir -p ./_build/cmd
	@go build -v -mod vendor -i -o ./_build/cmd/mks.bin \
		-ldflags "-X $(MOD).version=`cat ./_build/version`" ./cmd/mks

.PHONY: _build/version
_build/version:
	@mkdir -p ./_build
	@echo "`cat VERSION.txt`-`date -u '+%Y%m%d.%H%M%S'`-`git describe --always --dirty`" >./_build/version

.PHONY: clean
clean:
	@go clean -mod vendor -i -cache $(PKG)
	@rm -vf ./mks
	@rm -vrf ./_build ./_testing

.PHONY: distclean
distclean:
	@go clean -mod vendor -i -cache -modcache -testcache $(PKG)

.PHONY: test
test: build
	@go test -mod vendor $(PKG)

.PHONY: testclean
testclean:
	@go clean -mod vendor -testcache $(PKG)

.PHONY: cover
cover: build
	@mkdir -p ./_testing
	@go test -mod vendor -coverprofile ./_testing/cover.out $(PKG)
	@go tool cover -html ./_testing/cover.out -o ./_testing/coverage.html

.PHONY: vendor
vendor:
	@go mod vendor -v
	@go mod tidy
	@$(MAKE) fmt >/dev/null

.PHONY: fmt
fmt:
	@gofmt -w -l -s .

.PHONY: demo
demo: build
	@./_build/cmd/mks.bin ./demo/build.mks
