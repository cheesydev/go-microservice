package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cheesydev/go-microservice/product"
	"google.golang.org/grpc"
)

func main() {
	const port = 10000

	// create a listener in the desired port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register the RatingService implementing the functionality
	product.RegisterRatingsServer(grpcServer, &product.RatingService{})

	log.Printf("prodrating grpc server listening at port %d\n", port)
	grpcServer.Serve(lis)
}
