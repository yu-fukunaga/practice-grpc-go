package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"pancake.maker/app/server/pancake_baker/handler"
	"pancake.maker/gen/proto"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	grpc_zap.ReplaceGrpcLogger(zapLogger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(zapLogger),
				grpc_auth.UnaryServerInterceptor(auth),
			),
		),
	)
	proto.RegisterPancakeBakerServiceServer(
		server,
		handler.NewBakerHandler(),
	)
	reflection.Register(server) //required for grpc_cli

	go func() {
		log.Printf("start gRPC server port: %v\n", port)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}

func auth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "barer")
	if err != nil {
		return nil, err
	}
	if token != "test_token" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid barer token")
	}
	return context.WithValue(ctx, "UserName", "uyu"), nil
}
