.PHONY: build test install clean version

VERSION ?= dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS = -s -w -X github.com/ovitorvalente/gssh/internal/version.Version=$(VERSION) \
	-X github.com/ovitorvalente/gssh/internal/version.Commit=$(COMMIT) \
	-X github.com/ovitorvalente/gssh/internal/version.BuildDate=$(DATE)

build:
	go build -ldflags "$(LDFLAGS)" -o gssh ./cmd/gssh

test:
	go test ./... -v -race

install: build
	install -m 755 gssh $(shell go env GOPATH)/bin/gssh

clean:
	rm -f gssh gssh-*

version:
	@echo "Version: $(VERSION)"
	@echo "Commit:  $(COMMIT)"
	@echo "Date:    $(DATE)"
