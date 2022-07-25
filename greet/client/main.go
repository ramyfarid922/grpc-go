package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/ramyfarid922/grpc-go/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	// Define tls parameter
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// Dialing with disabled SSL
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	// Defer will just execute the conn.Close() at the end of the main func
	defer conn.Close()

	// Some stuff to do..
	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetWithDeadline(c, 2*time.Second)
	doGreetEveryone(c)
}
