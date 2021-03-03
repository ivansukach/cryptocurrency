PACKAGES=$(shell go list ./... | grep -v '/simulation')

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
VERSION = $(BRANCH)-$(COMMIT)
COMMIT := $(shell git log -1 --format='%H')

# TODO: Update the ldflags with the app, client & server names
ldflags = -X github.com/ivansukach/modified-cosmos-sdk/version.Name=octa \
	-X github.com/ivansukach/modified-cosmos-sdk/version.AppName=octadaemon \
	-X github.com/ivansukach/modified-cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/ivansukach/modified-cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/octadaemon

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify


# # look into .golangci.yml for enabling / disabling linters
# lint:
# 	@echo "--> Running linter"
# 	@golangci-lint run
# 	@go mod verify