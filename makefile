PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

GO = CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go

LDFLAGS = "-s -w -X ${PKG}/internal/version.Version=${VERSION}+sha.${COMMIT_SHA}"

GOBUILD=$(GO) build -a -ldflags $(LDFLAGS)

WORKSPACE ?= $(NAME)
GORUN = $(GO) run

local-run:
	cd ./cmd/$(WORKSPACE) && $(GORUN) .

build:
	cd ./cmd/$(WORKSPACE) && $(GOBUILD) -o ./$(WORKSPACE)

docker:
	docker build -t $(NAME):$(VERSION) .

swagger:
	swag init --pd -g ./cmd/$(WORKSPACE)/main.go -o ./cmd/$(WORKSPACE)/docs