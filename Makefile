# application repository and binary file name
NAME=go-app-template

# application repository path
REPOSITORY=github.com/antonfisher/${NAME}

# version from git
VERSION ?= $(shell git rev-parse --abbrev-ref HEAD | sed -e "s/.*\\///")
COMMIT ?= $(shell git rev-parse HEAD | cut -c 1-7)

# build ldflags
LDFLAGS ?= \
	-X ${REPOSITORY}/pkg/params.Name=${NAME} \
	-X ${REPOSITORY}/pkg/params.Version=${VERSION} \
	-X ${REPOSITORY}/pkg/params.Commit=${COMMIT}


all: build
.PHONY: all

build: build-linux-amd64
.PHONY: build

build-linux-amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(NAME) -ldflags "$(LDFLAGS)" ./cmd
.PHONY: build-amd64