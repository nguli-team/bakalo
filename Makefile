MAIN_FILE = "main.go"
BIN_NAME = "bakalo"
DOCKER_DIR = "docker"

setup-dev:
	go get -d -v ./...
	[ -d media ] || mkdir media
	docker-compose -f ${DOCKER_DIR}/docker-compose.local.yml pull

start-dev:
	docker-compose -f ${DOCKER_DIR}/docker-compose.local.yml up -d

clean:
	docker-compose -f ${DOCKER_DIR}/docker-compose.local.yml down -v
	rm -rf bin media

serve:
	go run ${MAIN_FILE} serve

migrate:
	go run ${MAIN_FILE} migrate

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/${BIN_NAME} .
