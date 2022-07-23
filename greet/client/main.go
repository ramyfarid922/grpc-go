package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	// Dialing with disabled SSL
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	// Defer will just execute the conn.Close() at the end of the main func
	defer conn.Close()

	// Some stuff to do..
	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	// doGreetManyTimes(c)
	doLongGreet(c)

}
