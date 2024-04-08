VERSION := $(shell git describe --tags | tr -d v)
IMAGE := bjoernalbers/ovpnd

.PHONY: all build unit integration_test release help

all: unit integration ## Build binary and run all tests

build: ## Build binary
	@go build -ldflags '-X main.Version=$(VERSION)'

unit: ## Run unit tests
	@go test ./...

integration: build ## Run integration tests
	@go test integration_test.go

image: ## Build docker image
	docker build --platform=linux/amd64 -t '$(IMAGE):latest' -t '$(IMAGE):$(VERSION)' .

release: ## Push changes to GitHub and Docker Hub
	git push origin
	git push origin "v$(VERSION)"
	docker push '$(IMAGE):latest'
	docker push '$(IMAGE):$(VERSION)'

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
