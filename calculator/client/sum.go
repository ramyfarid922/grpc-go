package main

import (
	"context"
	"log"

	pb "github.com/ramyfarid922/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	// You can use this function to either return a sum from the server rpc endpoint or just log the sum
	log.Printf("getSum was invoked..")

	// Now hit the server rpc endpoint with a populated SumRequest Message
	// Capture res and error
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		X: 10,
		Y: 3,
	})

	// We got error from the rpc service
	if err != nil {
		log.Fatalf("Couldn't get sum: %v\n", err)
	}

	// We got a response from rpc service
	log.Printf("Sum: %d\n", res.Result)
}
