package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cheesydev/go-microservice/pi"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type piServer struct {
}

func (s *piServer) AproximatePi(c context.Context, e *empty.Empty) (*pi.PiAproximation, error) {
	return &pi.PiAproximation{Value: 3.1}, nil
}

func main() {

	// create a listener
	const port = 7070
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pi.RegisterPiServer(grpcServer, &piServer{})

	fmt.Printf("grpc server listening at port %d\n", port)
	grpcServer.Serve(lis)
}
