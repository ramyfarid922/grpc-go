package main

import (
	"log"

	pb "github.com/ramyfarid922/grpc-go/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Define the grpc service address and port
var addr string = "localhost:50051"

func main() {
	// Get a connection instance by dialing the grpc into our grpc server address
	// Notice how we connect with insecure creds to disable SSL
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	//Next, what to do with this connection instance
	// Check if we even got error or not
	if err != nil {
		log.Fatalf("Failed to connect:  %v\n", err)
	}

	//And if there's no error, just keep the conn open
	// Defer will just execute the conn.Close() at the end of the main func
	defer conn.Close()

	// get a new Blog Service Client instance
	// plug our open connection to the Service Client instance
	c := pb.NewBlogServiceClient(conn)

	// Do whatever you want here
	// I expect I will do what I want with the FleetManager response continuing here

	id := createBlog(c)
	readBlog(c, id)
	// readBlog(c, "SomeInvalidID")
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
