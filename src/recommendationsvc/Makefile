.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && go build -o ./.bin/app ./cmd/main/main.go

run: build
	./.bin/app

proto:
	rm -rf gen/proto && mkdir -p gen/proto && protoc -I=../../proto/ --go_out=gen/ --go-grpc_out=gen/ ../../proto/shop.proto

lint:
	golangci-lint run

test:
	go test --short -coverprofile=cover.out -v ./...

test.coverage:
	go tool cover -func=cover.out

