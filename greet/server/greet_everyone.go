package main

import (
	"io"
	"log"
	"time"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone func was invoked with")

	// Who should start the communication?

	// Infinite loop to receive from client stream
	for {
		// Capture an item from the stream which is actually a GreetRequest{} message
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		// Construct our response of GreetResponse{} message to send it back
		res := "Hello " + req.FirstName
		time.Sleep(1 * time.Second)
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}

	// Notice that we don't return the nil error here
	// In streaming endpoint we return the function's error which is in signature
	// only when we are done receiving from client
}
