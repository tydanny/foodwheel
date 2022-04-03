build:
	docker build -t foodwheel .

run:
	docker run -d -p 8080:8080 --name foodwheel foodwheel

stop:
	docker container stop foodwheel

rm:
	docker container rm foodwheel