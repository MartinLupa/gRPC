package main

import (
	"log"
	"net"

	pb "github.com/MartinLupa/gRPC/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(server, &Server{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
