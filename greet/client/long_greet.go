package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	//Define what we are going to send
	reqs := []*pb.GreetRequest{
		{FirstName: "Ramy"},
		{FirstName: "Jenni"},
		{FirstName: "Jess"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling longGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
