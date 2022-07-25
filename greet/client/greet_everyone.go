package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone function was invoked")

	// Create a stream with the server
	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	// Prepare the number of requests we are gonna stream to server
	reqs := []*pb.GreetRequest{
		{FirstName: "Ramy"},
		{FirstName: "Jenny"},
		{FirstName: "Jess"},
	}

	// Prepare a channel for go routines communication
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(6 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
