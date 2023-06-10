PACKAGE_NAME=common

deps:
	@go mod download

proto:
	protoc --proto_path=proto \
	--go_out=gen --go_opt=paths=source_relative \
    --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=gen --grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    proto/*/**.proto

.PHONY: proto

build:
	go build -o bin/${PACKAGE_NAME}

clean:
	rm -rf bin

check:
	pre-commit run
