package main

import (
	"io"
	"log"

	pb "github.com/MartinLupa/gRPC/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")

	var max int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if num := req.Num; num > max {
			max = num
			stream.Send(&pb.NumResponse{
				Result: max,
			})
		}
	}
}
