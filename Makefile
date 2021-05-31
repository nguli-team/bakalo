MAIN_FILE = "main.go"
BIN_NAME = "bakalo"
DOCKER_DIR = "docker"

setup-dev:
	go get -v ./...
	[ -d media ] || mkdir media
	docker-compose -f ${DOCKER_DIR}/docker-compose.local.yml pull

clean:
	docker-compose -f ${DOCKER_DIR}/docker-compose.local.yml down -v
	rm -rf bin

start-dev:
	docker-compose -f ${DOCKER_DIR}/docker-compose.local.yml up -d

run-dev:
	go run ${MAIN_FILE} serve

build:
	go build -o bin/${BIN_NAME} ${MAIN_FILE}

migrate:
	go run ${MAIN_FILE} migrate
