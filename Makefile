VERSION := $(shell git describe --tags | grep -Eo '^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$$' | tr -d v)
IMAGE := bjoernalbers/ovpnd

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build binary
	@go build -ldflags '-X main.Version=$(VERSION)'

test: unit integration ## Run unit and integration tests

unit: ## Run unit tests
	@go test ./...

integration: build ## Run integration tests
	@go test integration_test.go

image: ## Build docker image
	docker build --platform=linux/amd64 -t '$(IMAGE):latest' -t '$(IMAGE):$(VERSION)' .

publish: ## Publish docker image
	docker push '$(IMAGE)' --all-tags
