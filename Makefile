.PHONY: env
env:
	cp .envrc-develop .envrc
	direnv allow .

.PHONY:setup
setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.46.2
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: protoc-gen
protoc-gen:
	protoc --go_out gen/. \
				--go_opt paths=source_relative \
				--go-grpc_out gen/. \
				--go-grpc_opt paths=source_relative \
				proto/*.proto

.PHONY: lint
lint:
	golangci-lint run