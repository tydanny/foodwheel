all: lint build 

build: lint
	docker build -t foodwheel .

deploy: stop clean build
	docker run -d -p 8080:8080 --name foodwheel foodwheel

run: lint
	go run cmd/main.go

stop:
	docker container stop foodwheel

clean: stop
	docker container rm foodwheel

lint:
	golangci-lint run  --fix ./... -E gosec,gofmt,misspell,testpackage,whitespace
	protolint lint -fix .

test:
	ginkgo run ./...

fmt:
	go fmt ./...

vet:
	go vet ./...