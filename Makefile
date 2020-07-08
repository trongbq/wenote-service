IMAGE_NAME := wenote-service
IMAGE_NAME_LATEST := $(IMAGE_NAME):latest
CONTAINER_NAME := $(IMAGE_NAME)-container

run:
	go run main.go

build:
	go build -o main main.go

docker-build:
	docker build -f Dockerfile -t $(IMAGE_NAME_LATEST) .

docker-run:
	docker run -p 8080:8080 -d --name $(CONTAINER_NAME) $(IMAGE_NAME_LATEST)

docker-stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
