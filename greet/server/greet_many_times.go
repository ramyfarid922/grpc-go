package main

import (
	"fmt"
	"log"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

// What are we trying to do here
// We are implementing the rpc handler func
// How?

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes func was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s. number %d", in.FirstName, i)

		// This is how we respond with a stream of messages
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
