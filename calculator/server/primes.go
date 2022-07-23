package main

import (
	"log"
	"time"

	pb "github.com/ramyfarid922/grpc-go/calculator/proto"
)

// Implement the rpc handler Primes
func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with: %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		// This divisor is a prime
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			time.Sleep(1 * time.Second)

			// updating number = number / divisor
			number /= divisor
		} else {
			// Try the next divisor
			divisor++
		}
	}

	return nil

}
