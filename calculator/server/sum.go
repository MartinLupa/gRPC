package main

import (
	"context"
	"log"

	pb "github.com/MartinLupa/gRPC/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.NumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)

	sum := in.Num1 + in.Num2

	return &pb.NumResponse{
		Result: sum,
	}, nil
}
