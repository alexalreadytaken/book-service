SERVICE_NAME=book-service.bin

gen_proto:
	./protogen.sh

build:
	go build -o ${SERVICE_NAME} ./cmd/main.go

run:
	./${SERVICE_NAME}

build_and_run: build run

test:
	go test ./..