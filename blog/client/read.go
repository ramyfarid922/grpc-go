package main

import (
	"context"
	"log"

	pb "github.com/ramyfarid922/grpc-go/blog/proto"
)

// Note that the Client functions start with lower case
func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")

	req := &pb.BlogId{Id: id}
	// Why do we pass context.Background
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
