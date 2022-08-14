# About

This repository is for learning gRPC.  
I'm studying with ["スターティング gRPC"](https://tatsu-zine.com/books/starting-grpc).  
.proto, server, and client are included in this repository.

# Reference

- gRPC official  
  https://grpc.io/docs/languages/go/quickstart/
- grpc_cli  
  https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md

# Quick Start

- [PancakeBakerService](./SERVICE_PANCAKE_BAKER.md)  
  PancakeBakerService is simple gRPC service.

- [ImageUploaderService](./SERVICE_IMAGE_UPLOADER.md)  
  ImageUploaderService is a client-side streaming service that sends image data to a server.
  Return information on images uploaded to the server.

- [ReversiService](./SERVICE_REVERSI.md)  
  ReversiService is a bidirectional streaming service that allows you to play reversi between two clients.

# Development

direnv is recommended.  
If you use direnv, run the following command.

```
make env
```

"setup" installs the packages needed for development.

```
make setup
```
