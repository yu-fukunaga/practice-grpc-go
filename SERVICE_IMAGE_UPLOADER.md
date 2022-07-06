# ImageUploderService

- Start Server

```
go run app/server/image_uploader/server.go
```

- Start Client

```
go run app/client/image_uploader/client.go -file_path=<image_path>
```

### grpc_cli

- List all the services exposed at a given port
  ```
  grpc_cli ls localhost:50051
  ```
- List ImageUploaderService with details

  ```
  grpc_cli ls localhost:50051 iamge.uploader.ImageUploaderService -l
  ```
