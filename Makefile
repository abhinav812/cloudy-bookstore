# README !!
# Use this Makefile only in a xterm window; e.g. GitBash (on Windows), any sh shell (bash/zsh on *nix)
#

include .env

# Go parameters
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: build help
all: help

## Init
verify-install: ## Verifies where go binaries are installed on machine
	@echo "  >  Verifying GO installation..."
	go version

install-dependencies: ## Installs go libraries used in build
	@echo " >  Installing build dependencies..."
	go get -u golang.org/x/lint/golint

## Build:
clean: ## Remove build related file
	@echo " > Cleaning up project output directories..."
	rm -fr ./bin
	rm -fr ./out

go-vet: ## Run go vet on the *.go files
	@echo "  >  Running GO VET..."
	go vet ./...

go-lint: ## Run go linter on the *.go files
	@echo "  >  Running GO LINT..."
	golint ./...

go-build: ## Build the project and put the output binary in out/bin/
	@echo "  >  Running GO BUILD..."
	mkdir -p out/bin
	GO111MODULE=on go build -o out/bin/${BINARY_NAME} ./cmd/bookstore/main.go

## Test:
go-test: ## Run the tests of the project
	@echo " > Running tests..."
	go test -v ./...

## All-in-one build:
build: clean verify-install install-dependencies go-vet  go-lint go-test go-build ## Build this project by running vet/lint/test/build targets and generated binary in out/bin/

## Docker:
docker-build: ## Use the dockerfile to build the container
	@echo " > Building Docker image for CI_PLATFORM - $(CI_PLATFORM) ..."
	docker-compose  build --force-rm


docker-release: ## Release the container with tag latest and version
	@echo " > Tagging Docker image... | $(CI_PLATFORM)"
	docker tag $(IMAGE_NAME) $(IMAGE_NAME):latest
	docker tag $(IMAGE_NAME) $(IMAGE_NAME):$(TAG)
ifneq ($(CI_PLATFORM), local) # Do not push docker images from local
	# Push the docker images
	@echo " > Pushing Docker image to docker registry for CI_PLATFORM - $(CI_PLATFORM)..."
	docker push $(IMAGE_NAME):latest
	docker push $(IMAGE_NAME):$(TAG)
endif


docker-build-release: docker-build docker-release ## Build and release docker images in one go

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)