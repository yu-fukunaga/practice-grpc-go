.PHONY: env
env:
	cp .envrc-develop .envrc
	direnv allow .

.PHONY:setup
setup:
	go install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: protoc-gen
protoc-gen:
	protoc -Iproto --go_out=plugins=grpc:. proto/*.proto