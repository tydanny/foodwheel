IMAGES_DIR ?= images
FOODWHEEL_IMAGE ?= ${IMAGES_DIR}/foodwheel
REDIS_IMAGE ?= redis/redis-stack:latest

GOLANGCI_VERSION ?= v1.62.2

.PHONY: all
all: lint test build 

.PHONY: build
build:
	go build -a -o bin/foodwheel cmd/main.go

.PHONY: image
image: generate lint test
	docker build -t foodwheel -f ${FOODWHEEL_IMAGE}/Dockerfile .

.PHONY: deploy
deploy: stop deploy-db
	docker run --rm -d -p 50051:50051 --name foodwheel foodwheel

.PHONY: deploy-db
deploy-db: stop-db
	docker run --rm -d -p 6379:6379 -p 8001:8001 --name foodwheel-redis ${REDIS_IMAGE}

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: stop
stop: stop-db
	-docker container stop foodwheel

.PHONY: stop-db
stop-db:
	-docker container stop foodwheel-redis

.PHONY: generate
generate:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/foodwheel/foodwheel.proto

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
PROTOLINT ?= $(LOCALBIN)/protolint

.PHONY: protolint localbin bin

bin: localbin protolint

localbin: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

protolint: $(PROTOLINT)
$(PROTOLINT): $(LOCALBIN)
	test -s $(LOCALBIN)/protolint || GOBIN=$(LOCALBIN) go install github.com/yoheimuta/protolint/cmd/protolint@latest
