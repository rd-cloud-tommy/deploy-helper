BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

VERSION ?= $(shell git describe --tags 2>/dev/null || echo "dev")
GIT_SHA ?= $(shell git rev-parse --short HEAD)
BUILD_DATE ?= $(shell date "+%Y-%m-%d")

GO_LDFLAGS := -X deploy-helper/cmd.version=$(VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X deploy-helper/cmd.githash=$(GIT_SHA) $(GO_LDFLAGS)

bin/deploy-helper:
	go build -ldflags "${GO_LDFLAGS}" -o "$@" ./

test:
	go test ./... -cover
.PHONY: test

clean:
	rm -rf ./bin
.PHONY: clean

DESTDIR :=
prefix  := /usr/local
bindir  := ${prefix}/bin

install: bin/deploy-helper
	install -d ${DESTDIR}${bindir}
	install -m755 bin/deploy-helper ${DESTDIR}${bindir}/
.PHONY: install

uninstall:
	rm -f ${DESTDIR}${bindir}/deploy-helper
.PHONY: uninstall
