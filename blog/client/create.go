package main

import (
	"context"
	"log"

	pb "github.com/ramyfarid922/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {

	// We will create a pb.Blog{} request and we are gonna send that to ther server
	// The server is gonna respond with either an error or a pb.BlogId{} response

	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Ramy",
		Title:    "My First blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
