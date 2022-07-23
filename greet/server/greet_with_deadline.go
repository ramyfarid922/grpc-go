package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)

	// We want the server to respond late intentionally
	for i := 0; i < 3; i++ {
		// Checking if the deadline is exceeded
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client cancelled the request")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}

		// Intentional delay in the for loop
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
