package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"practice/app/grpc/server/image_uploader/handler"
	"practice/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	proto.RegisterImageUploadServiceServer(server, handler.NewImageUploaderHandler())
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
