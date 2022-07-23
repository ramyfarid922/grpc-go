package main

import (
	"context"
	"log"

	pb "github.com/ramyfarid922/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	// Populating the blog that will be updated
	// Notice here how we don't really send an id for update, we send an entire new updated blog instance including the id
	// that the old one is gonna be fetched with
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Ramy",
		Title:    "New Title",
		Content:  "Content with some addition",
	}

	// So yea we are always passing this context.Background() here because the GRPC handler signature on the server
	// is expecting a context.Context type as the first parameter
	// Call the server rpc
	_, err := c.UpdateBlog(context.Background(), newBlog)

	// Check errors and if there's no errors then we succeeded
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")

}
