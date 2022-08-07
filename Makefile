IMAGES_DIR ?= images
FOODWHEEL_IMAGE ?= ${IMAGES_DIR}/foodwheel
MONGODB_IMAGE ?= ${IMAGES_DIR}/testMongoDB
STARTUP_SCRIPT ?= ${MONGODB_IMAGE}/scripts

.PHONY: all
all: lint test build 

build: lint test generate
	go build -a -o bin/foodwheel cmd/server/main.go

.PHONY: image
image: generate
	docker build -t foodwheel -f ${FOODWHEEL_IMAGE}/Dockerfile .

.PHONY: build-test-db
build-test-db:
	docker build -t test-mongo-db -f ${MONGODB_IMAGE}/Dockerfile .

.PHONY: run-test-db
run-test-db: build-test-db stop-test-db
	mkdir -p /tmp/data
	docker run --rm -it -v /tmp/data:/data/db --name test-mongodb -d test-mongo-db

.PHONY: stop-test-db
stop-test-db:
	-@docker stop test-mongodb
	rm -rf /tmp/data

.PHONY: deploy
deploy: stop build
	docker run --rm -d --name foodwheel foodwheel

.PHONY: run
run: lint
	go run cmd/server/main.go

.PHONY: stop
stop:
	-docker container stop foodwheel

generate:
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

.PHONY: ginkgo
ginkgo: $(GINKGO)
$(GINKGO): $(LOCALBIN)
	test -s $(LOCALBIN)/gikngo || GOBIN=$(LOCALBIN) go install github.com/onsi/ginkgo/v2/ginkgo@latest

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI)
$(GOLANGCI): $(LOCALBIN)
	test -s $(LOCALBIN)/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCALBIN)

.PHONY: protolint
protolint: $(PROTOLINT)
$(PROTOLINT): $(LOCALBIN)
	test -s $(LOCALBIN)/protolint || GOBIN=$(LOCALBIN) go install github.com/yoheimuta/protolint/cmd/protolint@latest
