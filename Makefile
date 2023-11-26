
.PHONY: install-tools
install-tools:
# basic tools
	go install go.uber.org/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest
	go install golang.org/x/tools/cmd/goimports@latest
# protobuf
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

.PHONY: all
all: proto lint test build

proto:
	buf format -w
	rm -r ./pkg/apis ./doc/apis || true
	buf generate --template build/buf.gen.yaml
	go generate ./...

lint:
	yamlfmt .
	buf lint
	go mod tidy
	goimports -w .
	go vet ./...
	golangci-lint run ./...

test:
	go test -v ./...

build:
	go build ./cmd/server

