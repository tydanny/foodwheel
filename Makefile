IMAGES_DIR ?= images
FOODWHEEL_IMAGE ?= ${IMAGES_DIR}/foodwheel

.PHONY: all
all: lint test build 

build: lint test generate
	go build -a -o bin/foodwheel cmd/server/main.go

.PHONY: image
image: generate lint test
	docker build -t foodwheel -f ${FOODWHEEL_IMAGE}/Dockerfile .

.PHONY: deploy
deploy: stop build
	docker run --rm -d --name foodwheel foodwheel

.PHONY: run
run: lint
	go run cmd/server/main.go

.PHONY: stop
stop:
	-docker container stop foodwheel

.PHONY: generate
generate: $(PROTOC)
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/foodwheel/foodwheel.proto

.PHONY: lint
lint: golangci-lint protolint
	golangci-lint run  --fix ./... -E gosec,gofmt,misspell,testpackage,whitespace
	protolint lint -fix .

.PHONY: test
test: ginkgo
	$(GINKGO) run ./...

LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

GINKGO ?= $(LOCALBIN)/ginkgo
GOLANGCI ?= $(LOCALBIN)/golangci-lint
PROTOLINT ?= $(LOCALBIN)/protolint
PROTOC ?= $(LOCALBIN)/protoc

.PHONY: ginkgo protolint golangci-lint

ginkgo: $(GINKGO)
$(GINKGO): $(LOCALBIN)
	test -s $(LOCALBIN)/gikngo || GOBIN=$(LOCALBIN) go install github.com/onsi/ginkgo/v2/ginkgo@latest

golangci-lint: $(GOLANGCI)
$(GOLANGCI): $(LOCALBIN)
	test -s $(LOCALBIN)/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCALBIN)

protolint: $(PROTOLINT)
$(PROTOLINT): $(LOCALBIN)
	test -s $(LOCALBIN)/protolint || GOBIN=$(LOCALBIN) go install github.com/yoheimuta/protolint/cmd/protolint@latest

protoc: $(PROTOC)
$(PROTOC): $(LOCALBIN)
	test -s $(PROTOC)/protoc || curl -LO \
		https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip \


