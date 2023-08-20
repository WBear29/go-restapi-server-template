include .env
export

VERSION=test
LINTER_VERSION=v1.54.1

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

define generate
	rm -rf internal/http/model/model_*.go
	docker run --rm -v ${PWD}:/spec redocly/cli bundle openapi/openapi.yaml -o openapi/bundle.yaml
	docker run --rm -v ${PWD}:/local -v ${PWD}/internal/controller/http:/project \
	openapitools/openapi-generator-cli:v5.2.0 \
	generate -i /local/openapi/bundle.yaml \
	-c /project/.openapi-gen.json \
	-g go-server \
	-o /project/model \
	--global-property model \
	--ignore-file-override /project/.openapi-generator-ignore
	rm -rf openapi/bundle.yaml
endef

generate:
	$(call generate)
.PHONY: generate

build:
	$(call generate)
	go build -ldflags "-X main.version=$(VERSION)" -o $(APP_NAME) main.go
.PHONY: build

lint:
	docker run --rm -v `pwd`:/app -w /app golangci/golangci-lint:${LINTER_VERSION} golangci-lint run
.PHONY: lint

lint-fix:
	docker run --rm -v `pwd`:/app -w /app golangci/golangci-lint:${LINTER_VERSION} golangci-lint run --fix
.PHONY: lint-fix

licenses-show:
	go-licenses report ./... 2>/dev/null
.PHONY: licenses-show

licenses-report:
	go-licenses report ./... > licenses.csv 2>/dev/null
.PHONY: licenses-report

vulncheck:
	govulncheck ./...
.PHONY: vulncheck