IMG ?= foodwheel

GOLANGCI_VERSION ?= v1.62.2

.PHONY: all
all: lint test build protoall

.PHONY: build
build:
	go build -a -o bin/foodwheel cmd/main.go

.PHONY: image
image: generate lint test
	docker build -t ${IMG} -f images/foodwheel/Dockerfile .

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

generate:
	buf generate

protolint:
	buf lint

GOLANGCI ?= golangci-lint

.PHONY: lint
lint:
	$(GOLANGCI) run ./...

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

# This section that downloads tools needed to run some targets

LOCALBIN ?= $(shell pwd)/bin

.PHONY: localbin bin

bin: localbin protolint

localbin: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)
