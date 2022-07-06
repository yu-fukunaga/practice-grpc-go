# About

This repository is for learning gRPC.  
.proto, server, and client are included in one repository.

- [gRPC official](https://grpc.io/docs/languages/go/quickstart/)
- [grpc_cli](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md)

# Quick Start

- [PancakeBakerService](./SERVICE_PANCAKE_BAKER.md)  
  PancakeBakerService is simple gRPC service.

- [ImageUploaderService](./SERVICE_IMAGE_UPLOADER.md)  
  ImageUploaderService is a client-side streaming service that sends image data to a server.
  Return information on images uploaded to the server.

# Development

direnv is recommended.  
If you use direnv, run the following command.

```
make env
```

Install packages necessary for development.

```
make setup
```
