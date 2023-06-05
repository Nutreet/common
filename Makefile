PACKAGE_NAME=common

deps:
	@go mod download

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*/**.proto

.PHONY: proto

build:
	@go build -o bin/${PACKAGE_NAME}

clean:
	@rm -rf bin

check:
	@pre-commit run
