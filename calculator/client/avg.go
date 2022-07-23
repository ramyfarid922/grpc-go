package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ramyfarid922/grpc-go/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg function was invoked")

	// I have a hard time understanding this line
	// Looks like the rpc endpoint returns a stream to be used in communication with it?
	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
		time.Sleep(1 * time.Second)
	}

	// Getting response from the server
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)

}
