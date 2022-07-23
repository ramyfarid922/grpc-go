package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	// using context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Ramy",
	}

	// instead of always calling the server rpc with context.Background()
	// This time we are passing ctx which is a context with timeout
	// Meaning we are telling the server that we are putting a deadline on this request
	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			// Meaning we have a grpc error
			// The error here is related to the service input
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.Canceled {
				log.Println("We probably exceeded deadline!")
				return
			}
		} else {
			log.Fatalf("A non grpc error: %v\n", res.Result)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
