package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cheesydev/go-microservice/product"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
		os.Exit(-1)
	}
	defer conn.Close()

	client := product.NewRatingsClient(conn)
	p := product.Product{Id: "10"}
	r, err := client.GetRating(context.TODO(), &p)
	if err != nil {
		log.Fatalf("failed to call grpc: %v", err)
		os.Exit(-1)
	}
	fmt.Printf("Rating of product id %s: %v stars.\n", p.Id, r.Value)
}
