package main

import (
	"log"
	"net"

	pb "github.com/ramyfarid922/grpc-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	// Now hook our grpc server to the tcp port
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve %s\n", err)
	}

}
