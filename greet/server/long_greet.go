package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

// Implementing the LongGreet() rpc endpoint server handler
// The function signature is obtained from the greet_grpc.pb.go file after recompiling the proto files
// why doesn't the signature include the returned type although the rpc service definition denotes otherwise
func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	// Empty string to be concatenated into
	res := ""

	// The way we handle a stream input

	// Infinite for loop
	for {
		req, err := stream.Recv()

		// Check if we reached the end of the stream
		if err == io.EOF {
			// Client not sending anymore
			// Send greet response and close stream
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		// Build the response by concatenating into our empty string
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}

}
