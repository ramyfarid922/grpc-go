package main

import (
	"context"
	"log"

	pb "github.com/ramyfarid922/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %s\n", in)

	// Transform the string in.Id into a mongoDB objectID type using the primitive package
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	// How to read from mongo
	// Basically creating an empty model instance
	data := &BlogItem{}
	// Means we  are filtering by ID
	filter := bson.M{"_id": oid}
	// Now use the collection API
	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Can't find blog with ID provided",
		)
	}

	// Now return our proto response message
	// Don't forget the nil because of the return signature of the function
	return documentToBlog(data), nil
}
