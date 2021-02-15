package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Failed to listen on port 4040: %v", err)
	}

	gs := grpc.NewServer()

	if err := gs.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server on port 4040: %v", err)
	}
}

func (s *server) Getcontact(ctx context.Context, request *proto.Request) (*proto.Response, error) {}
