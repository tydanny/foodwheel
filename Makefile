build: validate
	docker build -t foodwheel .

run: build
	docker run -d -p 8080:3000 --name foodwheel foodwheel

stop:
	docker container stop foodwheel

clean: stop
	docker container rm foodwheel

validate: fmt vet

fmt:
	go fmt ./...

vet:
	go vet ./...