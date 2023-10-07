package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MartinLupa/gRPC/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg %v\n", err)
	}

	reqs := []*pb.NumRequest{
		{Num: 2},
		{Num: 2},
		{Num: 2},
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg %v\n", err)
	}

	log.Printf("Avg: %v\n", res.Result)
}
