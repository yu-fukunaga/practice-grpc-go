package main

import (
	"context"
	"flag"
	"log"
	"time"

	"practice/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:")
	}
	defer conn.Close()

	// Bearer token
	md := metadata.New(map[string]string{"authorization": "barer test_token"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Contact the server and print out its response.
	client := proto.NewPancakeBakerServiceClient(conn)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := client.Bake(ctx, &proto.BakeRequest{
		Menu: proto.Pancake_BACON_AND_CHEESE,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Pancake: %s", res.GetPancake())
}
