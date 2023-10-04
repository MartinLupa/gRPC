package main

import (
	"context"
	"log"

	pb "github.com/MartinLupa/gRPC/greet/proto"
)

// Function signature taken from GreetServiceServer interface inside greet_grpc.pb.go
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
