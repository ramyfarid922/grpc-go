package main

import (
	"context"
	"log"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

// Basically we are now implementing the rpc endpoints server handlers
// Notice the receiver with server struct is crucial here
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	// Acknowledging the handler invokaion
	log.Printf("Greet function was invoked with %v\n ", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
