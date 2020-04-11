package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cheesydev/go-microservice/pi"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:7070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
		os.Exit(-1)
	}
	defer conn.Close()

	client := pi.NewPiClient(conn)
	piAprox, err := client.AproximatePi(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("failed to call grpc: %v", err)
		os.Exit(-1)
	}
	fmt.Printf("Pi aproximation %v\n", piAprox.Value)
}
