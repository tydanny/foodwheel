IMAGES_DIR ?= "images"
FOODWHEEL_IMAGE ?= "${IMAGES_DIR}/foodwheel"
MONGODB_IMAGE ?= "${IMAGES_DIR}/testMongoDB"

.PHONY: all
all: lint build 

.PHONY: build
build: lint
	docker build -t foodwheel -f ${FOODWHEEL_IMAGE}/Dockerfile .

.PHONY: build-test
build-test:
	docker build -t test-mongodb -f ${MONGODB_IMAGE}/Dockerfile .

.PHONY: run-test-db
run-test-db: build-test
	docker run --rm -it -v mongoTestData:/data/db --name test-mongodb -d mongo

.PHONY: stop-test-db
stop-test-db:
	docker stop test-mongodb

.PHONY: deploy
deploy: stop clean build
	docker run --rm -d --name foodwheel foodwheel

.PHONY: run
run: lint
	go run cmd/main.go

.PHONY: stop
stop:
	docker container stop foodwheel

.PHONY: clean
clean: stop
	docker container rm foodwheel

.PHONY: lint
lint:
	golangci-lint run  --fix ./... -E gosec,gofmt,misspell,testpackage,whitespace
	protolint lint -fix .

.PHONY: test
test:
	ginkgo run ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...