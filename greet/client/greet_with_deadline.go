package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MartinLupa/gRPC/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Martin",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("Non gRPC error: %v\n", e)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
