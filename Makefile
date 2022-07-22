PROJECTNAME=terraform-provider-pingdirectory

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)

# ---------------------- targets -------------------------------------

.PHONY: default
default: help

.PHONY: cit
cit: clean build test ## clean build and test

.PHONY: clean
clean: ## clean build output
	rm -rf ./bin

.PHONY: gen
gen: ## invoke go generate
	@CGO_ENABLED=1 go generate ./...

.PHONY: build
build: clean gen ## build for current platform
	mkdir -p ./bin
	go build -o ./bin/terraform-provider-pingdirectory
	go build -o ./bin/ping-api-client pkg/cmd/ping-api-client/main.go

.PHONY: test
test: testacc ## Run all tests (unit + acceptance)

.PHONY: testacc
testacc: ## Run acceptance tests
	#TF_ACC=1 PING_SYNC_BASE_URL=https://lv1stgpingsync01.dm.nfl.com:8443 PING_SYNC_DN_NAME="${STG_PING_DIRECTORY_MANAGER_USERNAME}" PING_SYNC_DN_PASS="${STG_PING_DIRECTORY_MANAGER_PASSWORD}" go test ./... -v $(TESTARGS) -timeout 120m
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

test-action-push: ## Test github actions with event 'push'
	# https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#webhook-payload-example-37
	act push --container-architecture linux/amd64 --eventpath .local/push-tags-payload.json

.PHONY: help
help: Makefile
	@echo
	@echo " ${PROJECTNAME} - available targets:"
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo
