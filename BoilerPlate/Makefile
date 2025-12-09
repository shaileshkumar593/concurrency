APP=gin-jwt-swagger
IMAGE=your-registry/${APP}
TAG=latest

.PHONY: build docker-build

build:
\tgo build -o bin/${APP} ./cmd/server

docker-build:
\tdocker build -t ${IMAGE}:${TAG} .
