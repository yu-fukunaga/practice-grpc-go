package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"time"

	"practice/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	addr      = flag.String("addr", "localhost:50051", "the address to connect to")
	file_path string
)

func init() {
	flag.StringVar(&file_path, "file_path", "", "Uploadする画像ファイルパス")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:")
	}
	defer conn.Close()

	md := metadata.New(map[string]string{"authorization": "barer test_token"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client := proto.NewImageUploadServiceClient(conn)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Uploadするファイル
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// メソッドを呼び出すことでstreamが作成される
	stream, err := client.Upload(ctx)
	if err != nil {
		log.Fatalf("%v.Upload(_) = _, %v", client, err)
	}

	// メタデータを作成
	meta := &proto.ImageUploadRequest{
		File: &proto.ImageUploadRequest_FileMeta_{
			FileMeta: &proto.ImageUploadRequest_FileMeta{
				Filename: file_path,
			},
		},
	}

	// streamにメタデータを送る
	err = stream.Send(meta)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data := &proto.ImageUploadRequest{
			File: &proto.ImageUploadRequest_Data{
				Data: buf,
			},
		}
		err = stream.Send(data)
		if err != nil {
			log.Fatal(err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", res)

}
