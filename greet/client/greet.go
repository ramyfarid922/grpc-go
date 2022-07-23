package main

import (
	"context"
	"log"

	// The pb variable contains our protobuf schemas
	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

// Looks like we are passing the service client instance to the implementation of the corresponding rpc in that service
// doGreet() implements the client version of the Greet() rpc in the Greet service
// Note that we can also use a struct style to proceed into this implementation
func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		// Why's firstname here passed in CamelCase although in the protobuf the field was snake_case?
		FirstName: "Ramy",
	})

	// We got error from the rpc service
	if err != nil {
		log.Fatalf("Couldn't greet: %v\n", err)
	}

	// We got a response from rpc service
	log.Printf("Greeting: %s\n", res.Result)
}
