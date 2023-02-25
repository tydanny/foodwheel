IMAGES_DIR ?= images
FOODWHEEL_IMAGE ?= ${IMAGES_DIR}/foodwheel
REDIS_IMAGE ?= redis/redis-stack:latest

.PHONY: all image deploy deploy-db run stop stop-db generate lint fmt
all: lint test build 

build: lint test generate
	go build -a -o bin/foodwheel cmd/main.go

image: generate lint test
	docker build -t foodwheel -f ${FOODWHEEL_IMAGE}/Dockerfile .

deploy: stop deploy-db
	docker run --rm -d -p 50051:50051 --name foodwheel foodwheel

deploy-db: stop-db
	docker run --rm -d -p 6379:6379 -p 8001:8001 --name foodwheel-redis ${REDIS_IMAGE}

run: lint
	go run cmd/server/main.go

stop: stop-db
	-docker container stop foodwheel

stop-db:
	-docker container stop foodwheel-redis

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/foodwheel/foodwheel.proto

lint: golangci-lint protolint
	$(PROTOLINT) lint .
	$(GOLANGCI) run ./...
	go vet ./...

fmt:
	$(PROTOLINT) lint -fix .
	$(GOLANGCI) run --fix ./...
	go fmt ./...
	

test: ginkgo
	$(GINKGO) run ./...

# This section that downloads tools needed to run some targets

LOCALBIN ?= $(shell pwd)/bin
GINKGO ?= $(LOCALBIN)/ginkgo
GOLANGCI ?= $(LOCALBIN)/golangci-lint
PROTOLINT ?= $(LOCALBIN)/protolint

.PHONY: ginkgo protolint golangci-lint localbin bin

bin: localbin ginkgo protolint golangci-lint

localbin: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

ginkgo: $(GINKGO)
$(GINKGO): $(LOCALBIN)
	test -s $(LOCALBIN)/gikngo || GOBIN=$(LOCALBIN) go install github.com/onsi/ginkgo/v2/ginkgo@latest

golangci-lint: $(GOLANGCI)
$(GOLANGCI): $(LOCALBIN)
	test -s $(LOCALBIN)/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCALBIN)

protolint: $(PROTOLINT)
$(PROTOLINT): $(LOCALBIN)
	test -s $(LOCALBIN)/protolint || GOBIN=$(LOCALBIN) go install github.com/yoheimuta/protolint/cmd/protolint@latest
