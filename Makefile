PREFIX?=$(shell pwd)

# Used to populate version variable in main package.
GO_LDFLAGS=-ldflags "-X `go list ./version`.Version `git describe --match 'v[0-9]*' --dirty='.m' --always`"

.PHONY: clean all fmt vet lint build test binaries
.DEFAULT: default
all: AUTHORS clean fmt vet fmt lint build test binaries

AUTHORS: .mailmap .git/HEAD
	 git log --format='%aN <%aE>' | sort -fu > $@

# This only needs to be generated by hand when cutting full releases.
version/version.go:
	./version/version.sh > $@

vet:
	@echo "+ $@"
	@go vet ./...

fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . | grep -v Godeps/_workspace/src/ | tee /dev/stderr)" || \
		echo "+ please format Go code with 'gofmt -s'"

lint:
	@echo "+ $@"
	@test -z "$$(golint ./... | grep -v Godeps/_workspace/src/ | tee /dev/stderr)"

build:
	@echo "+ $@"
	@go build -v ${GO_LDFLAGS} ./...

test:
	@echo "+ $@"
	@go test -test.short ./...

test-full:
	@echo "+ $@"
	@go test ./...

api-docs:
	@echo "+ $@"
	@raml2html specs/api.raml > ./public/index.html
