package main

import (
	"context"
	"log"

	pb "github.com/ramyfarid922/grpc-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	req := &pb.SqrtRequest{
		Number: n,
	}

	res, err := c.Sqrt(context.Background(), req)

	if err != nil {
		// log.Fatalf("Error while calling Sqrt: %v\n", err)
		e, ok := status.FromError(err)

		if ok {
			// Meaning we have a grpc error
			// The error here is related to the service input
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			// We have a non-grpc error
			log.Fatalf("A non grpc error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
