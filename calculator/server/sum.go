// We will use this file to implement just the sum rpc endpoint for the calculator grpc service
// So I am changing the name from calculator.go to sum.go
package main

import (
	"context"
	"log"

	pb "github.com/ramyfarid922/grpc-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with: %v\n", in)

	// Again, why do I have to capitalize the SunRequest field names although they are snake_case in the Message definition
	return &pb.SumResponse{
		Result: in.X + in.Y,
	}, nil
}
