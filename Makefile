ENV ?= local

IMAGE_NAME := wetodo-service
IMAGE_NAME_LATEST := $(IMAGE_NAME):latest
CONTAINER_NAME := $(IMAGE_NAME)-container

run:
	ENV=$(ENV) go run cmd/wetodo-api/main.go

build:
	go build -tags api -o ./build/wetodo-api ./cmd/wetodo-api/.

docker-build:
	docker build -f Dockerfile -t $(IMAGE_NAME_LATEST) .

docker-run:
	docker run -p 8080:8080 -d --name $(CONTAINER_NAME) $(IMAGE_NAME_LATEST)

docker-stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
