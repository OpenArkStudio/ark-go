VERSION := $(shell git describe --exact-match --tags 2>/dev/null)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := $(LDFLAGS) -s -w -X main.commit=$(COMMIT) -X main.branch=$(BRANCH)

ifdef VERSION
	LDFLAGS += -X main.version=$(VERSION)
endif

.PHONY: app
app:
	CGO_ENABLED=1 go build -ldflags "$(LDFLAGS)" -o bin/app ./app/main.go
	go build -buildmode=plugin -o bin/so/AFHttpPlugin.so plugin/http/AFHttpPlugin.go
	go build -buildmode=plugin -o bin/so/AFLogPlugin.so plugin/log/AFLogPlugin.go
