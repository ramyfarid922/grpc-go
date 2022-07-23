package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ramyfarid922/grpc-go/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("---listBlog was invoked---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	// Remember handling a stream response is always an infinte loop
	// The infinite for loop serves as a polling mechanism to keep checking if we received something onthe stream
	for {
		res, err := stream.Recv()
		// time.Sleep(2 * time.Second)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
