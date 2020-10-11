PKG ?= ./...
ARGS ?=
SITE ?= _site

MOD := github.com/jrmsdev/gomks

.PHONY: default
default: build

.PHONY: build
build: _build/cmd/mks.bin

.PHONY: check
check: build test demo cover
	@./_build/cmd/mks.bin -version

.PHONY: _build/cmd/mks.bin
_build/cmd/mks.bin: _build/version
	@mkdir -p ./_build/cmd
	@go build -v -mod vendor -i -o ./_build/cmd/mks.bin \
		-ldflags "-X $(MOD).build=`cat ./_build/version`" ./cmd/mks

GIT_DESCRIBE ?= --always --dirty --first-parent --tags

.PHONY: _build/version
_build/version:
	@mkdir -p ./_build
	@echo "`date -u '+%Y%m%d.%H%M%S'`-`git describe $(GIT_DESCRIBE)`" >./_build/version

.PHONY: clean
clean:
	@go clean -mod vendor -i -cache $(PKG)
	@rm -vf ./mks
	@rm -vrf ./_build ./_testing

.PHONY: distclean
distclean:
	@$(MAKE) clean testclean
	@go clean -mod vendor -i -cache -modcache -testcache $(PKG)

.PHONY: test
test: build
	@(MKSLOG=debug go test -mod vendor $(ARGS) $(PKG))

.PHONY: testclean
testclean:
	@go clean -mod vendor -testcache $(PKG)

.PHONY: cover
cover: build
	@mkdir -p ./_testing
	@(MKSLOG=debug go test -mod vendor -coverprofile ./_testing/cover.out $(PKG))
	@go tool cover -html ./_testing/cover.out -o ./_testing/coverage.html

.PHONY: vendor
vendor:
	@go mod vendor -v
	@go mod tidy
	@$(MAKE) fmt >/dev/null

.PHONY: vendor-upgrade
vendor-upgrade:
	@mkdir -p ./_build
	@(echo 'github.com/jrmsdev/gomks'; echo ''; echo 'go 1.15') >./_build/go.mod
	@$(MAKE) vendor

.PHONY: fmt
fmt:
	@gofmt -w -l -s .

.PHONY: demo
demo: build
	@(cd ./demo && ../_build/cmd/mks.bin build.mks)

.PHONY: demo-serve
demo-serve: demo
	@$(MAKE) serve SITE=demo/_site

.PHONY: serve
serve: build
	@./_build/cmd/mks.bin -serve $(SITE) $(ARGS)
