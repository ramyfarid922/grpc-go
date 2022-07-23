package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ramyfarid922/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog function was invoked with: %v\n", in)

	// instantiate an instance of the BlogItem (Model)
	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	// Insert in mongodb
	res, err := collection.InsertOne(ctx, data)
	// Error handling, internal error
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v\n", err),
		)
	}

	// If we succeeded in creating a mongo document
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Can't convert to OID",
		)
	}

	// Now return our proto response message
	// Don't forget the nil because of the return signature of the function
	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil

}
