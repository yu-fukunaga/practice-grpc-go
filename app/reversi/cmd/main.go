package main

import (
	"os"
	client "practice/app/grpc/client/reversi"
)

func main() {
	os.Exit(client.NewReversi().Run())
}
