package main

import (
	"io"
	"log"

	pb "github.com/ramyfarid922/grpc-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	// Now you have a stream of requests from the client
	// Each request contains a number. We need to average the numbers in each request
	var sum int32 = 0
	count := 0

	// The way to usually handle an incoming stream is by starting an infinite loop with stream.Recv()
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving number: %d\n", req.Number)
		sum += req.Number
		count++
	}

}
