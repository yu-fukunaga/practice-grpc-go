# PancakeBakerService

- Start Server

```
go run app/grpc/server/pancake_baker/server.go
```

- Start Client

```
go run app/grpc/client/pancake_baker/client.go
```

### grpc_cli

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
