IMAGE_NAME="kavu/curraunt:latest"

all: build

build: scripts/build.sh
	@bash scripts/build.sh

docker_start:
	@docker run --name curraunt -d -p 8080:8080 $(IMAGE_NAME)

docker_stop:
	@docker rm -v -f curraunt

docker_clean:
	@docker rmi $(IMAGE_NAME)

docker_build:
	@docker build -t kavu/curraunt:$(shell git rev-parse --short HEAD) .
	@docker tag kavu/curraunt:$(shell git rev-parse --short HEAD) $(IMAGE_NAME)

docker_rebuild: docker_clean docker_build

docker_push:
	@docker push $(IMAGE_NAME)
