VERSION := $(shell git describe --exact-match --tags 2>/dev/null)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := $(LDFLAGS) -s -w -X main.commit=$(COMMIT) -X main.branch=$(BRANCH)

ifdef VERSION
	LDFLAGS += -X main.version=$(VERSION)
endif

.PHONY: demo
demo:
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o bin/demo ./server/demo/main.go
