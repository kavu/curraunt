NAME="curraunt"
IMAGE_NAME="kavu/$(NAME):latest"
CONTAINER_NAME=$(NAME)

all: build

build: scripts/build.sh
	@bash scripts/build.sh

docker_start:
	@docker run --name $(CONTAINER_NAME) -d -p 8080:80 $(IMAGE_NAME)

docker_stop:
	@docker rm -v -f $(CONTAINER_NAME)

docker_clean:
	@docker rmi $(IMAGE_NAME)

docker_build:
	@docker build -t $(IMAGE_NAME) .

docker_build_tag:
	@git co $(TAG)
	@docker build -t kavu/$(NAME):$(TAG) .

docker_push_tag:
	@docker push kavu/$(NAME):$(TAG)

docker_rebuild: docker_clean docker_build

docker_push:
	@docker push $(IMAGE_NAME)
