package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"practice/app/server/reversi/handler"
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

	proto.RegisterReversiMathingServiceServer(server, handler.NewMatchingHandler())
	proto.RegisterGameServiceServer(server, handler.NewGameHandler())

	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		err := server.Serve(lis)
		if err != nil {
			return
		}
	}()

	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
