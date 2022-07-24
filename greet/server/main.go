// Think of main.go as index.js in Node.js express apps
// Here happens:
// 1- the bootstraping for the grpc server
// 2- Defining the address and port our server will listen on
// 3- Defining the server struct that will implement our rpc endpoint handlers
// 4- Hooking up our grpc service whatever.proto to our powered on grpc server s via pb.RegisterGreetServiceServer(s, &Server{})
// 5- Handling errors of server bootstrapping..to be continued

// Think of greet.go and greet_many_times.go as in exppress route handlers or controllers of the mvc world
package main

import (
	"log"
	"net"

	// Think of the pb as your gateway to your grpc services
	pb "github.com/ramyfarid922/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

// The server struct will be used to implement all the rpc endpoints
type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true // Change to false if needed

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	// Now we have s, which is an rpc server but it doesn't know which grpc service to expose on the port
	// The 3 dots is expanding the array into multiple func parameters
	s := grpc.NewServer(opts...)

	// We are telling the grpc server please expose or GreetService by passing the grpc server variable s
	// to the RegisterGreetServiceServer
	// Basically we are plugging our grpc server cord s to a grpc service definition to make it alive
	// Need to understand the arguments to the register func
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s\n", err)
	}
}
