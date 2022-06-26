# About

This repository is for learning gRPC.  
.proto, server, and client are included in one repository.

- [gRPC official](https://grpc.io/docs/languages/go/quickstart/)
- [grpc_cli](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md)

# Quick Start

- Start Server

```
go run app/server/pancake_baker/main.go
```

### Method1. use grpc_cli

- List all the services exposed at a given port
  ```
  grpc_cli ls localhost:50051
  ```
- List PancakeBakerService with details

  ```
  grpc_cli ls localhost:50051 pancake.maker.PancakeBakerService -l
  ```

- Call a unary method Send a rpc to a helloworld server at localhost:50051

  ```
  grpc_cli call localhost:50051 pancake.maker.PancakeBakerService.Bake 'menu: 1'
  ```

### Method2. start client

```
go run app/client/main.go
```

# Development

direnv is recommended.  
If you use direnv.

```
make env
```

Install packages necessary for development.

```
make setup
```
