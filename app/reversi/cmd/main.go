package main

import (
	"os"
	client "practice/app/client/reversi"
)

func main() {
	os.Exit(client.NewReversi().Run())
}
