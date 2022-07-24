IMAGES_DIR ?= images
FOODWHEEL_IMAGE ?= ${IMAGES_DIR}/foodwheel
MONGODB_IMAGE ?= ${IMAGES_DIR}/testMongoDB
STARTUP_SCRIPT ?= ${MONGODB_IMAGE}/scripts

.PHONY: all
all: lint build 

.PHONY: build
build: lint
	docker build -t foodwheel -f ${FOODWHEEL_IMAGE}/Dockerfile .

.PHONY: build-test
build-test:
	docker build -t test-mongo-db -f ${MONGODB_IMAGE}/Dockerfile .

.PHONY: run-test-db
run-test-db: build-test stop-test-db
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
	go run cmd/main.go

.PHONY: stop
stop:
	-docker container stop foodwheel

.PHONY: lint
lint:
	golangci-lint run  --fix ./... -E gosec,gofmt,misspell,testpackage,whitespace
	protolint lint -fix .

.PHONY: test
test:
	ginkgo run ./...
