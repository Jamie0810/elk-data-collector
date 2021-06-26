install:
	cd src && go mod tidy && \
	go mod vendor

docker-build:
	./build/build.sh

local-run:
	cd src && go build && \
	_LAMBDA_SERVER_PORT=8001 go run data-collector

invoke:
	cd src/example && go run lambda-invoke.go