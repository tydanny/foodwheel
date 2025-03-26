IMG ?= foodwheel

GOLANGCI_VERSION ?= v1.62.2

.PHONY: all
all: protoall lint test build

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell TZ=UTC0 git show --quiet --date='format-local:%Y-%m-%dT%T%z' --format="%cd")
LDFLAGS="-X 'main.GitCommit=${GIT_COMMIT}${GIT_DIRTY}' -X 'main.BuildDate=${BUILD_DATE}'"

.PHONY: build
build:
	go build -ldflags=${LDFLAGS} -a -o bin/foodwheel cmd/*.go

.PHONY: docker-build
docker-build: generate
	docker build -t ${IMG} \
		--build-arg COMMIT=${LDFLAGS} \
		.

.PHONY: start
start:
	podman compose --file ./test/docker-compose.yaml up --detach

.PHONY: stop
stop:
	podman compose --file ./test/docker-compose.yaml down

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: protoall generate protolint
protoall: generate protolint

generate: protoc-gen-go protoc-gen-go-grpc
	buf generate

protolint:
	buf lint

GOLANGCI ?= golangci-lint

.PHONY: lint
lint:
	$(GOLANGCI) run ./...
	$(API-LINTER) --config aip-lint.yaml -I api/foodwheel -I api/googleapis $(shell find api/foodwheel -iname "*.proto")

.PHONY: format
format:
	$(GOLANGCI) run --fix ./...

.PHONY: fmt
fmt:
	$(GOLANGCI) run --fix ./...

GINKGO ?= go run github.com/onsi/ginkgo/v2/ginkgo \
-r \
--randomize-all \
--randomize-suites \
--fail-on-pending \
--fail-on-empty \
--keep-going \
--cover \
--coverprofile=cover.profile \
--race \
--trace

GINKGO_RUNNER ?= $(GINKGO) \
-p

.PHONY: test
test:
	$(GINKGO_RUNNER) run ./...

GINKGO_RUNNER_CI ?= $(GINKGO) \
--procs=3 \
--compilers=3 \
--timeout=120s \
--poll-progress-after=60s \
--poll-progress-interval=30s \
--github-output

.PHONY: test-ci
test-ci:
	$(GINKGO_RUNNER_CI) run ./...

# This section downloads tools needed to run some targets

API-LINTER-VERSION ?= v1.69.2
PROTOC-VERSION ?= v1.36.5
GRPC-VERSION ?= v1.5.1

LOCALBIN ?= $(shell pwd)/bin
API-LINTER ?= $(LOCALBIN)/api-linter
PROTOC ?= $(LOCALBIN)/protoc-gen-go
PROTOC-GEN-GO-GRPC ?= $(LOCALBIN)/protoc-gen-go-grpc

.PHONY: localbin bin

bin: localbin protolint

localbin: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

.PHONY: api-linter
api-linter: $(APILINTER)
$(APILINTER): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/googleapis/api-linter/cmd/api-linter@$(API-LINTER-VERSION)

.PHONY: protoc-gen-go
protoc-gen-go: $(PROTOC)
$(PROTOC): $(LOCALBIN)
	@GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC-VERSION)

.PHONY: protoc-gen-go-grpc
protoc-gen-go-grpc: $(PROTOC-GEN-GO-GRPC)
$(PROTOC-GEN-GO-GRPC): $(LOCALBIN)
	@GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(GRPC-VERSION)

.PHONY: api-linter
api-linter: $(APILINTER)
$(APILINTER): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/googleapis/api-linter/cmd/api-linter@$(APILINTER_VERSION)
