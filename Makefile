build: fmt vet
	docker build -t foodwheel .

deploy: build
	docker run -d -p 8080:8080 --name foodwheel foodwheel

run: fmt vet
	go run cmd/main.go

stop:
	docker container stop foodwheel

clean: stop
	docker container rm foodwheel

fmt:
	go fmt ./...

vet:
	go vet ./...