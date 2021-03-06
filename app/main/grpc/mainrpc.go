package main

import (
	_"asyncMessageSystem/app/config"
	pb "asyncMessageSystem/app/controller/producer/grpc"
	_"asyncMessageSystem/app/middleware"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":3334"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterProducerServer(s, &server{})
	pb.RegisterProducerServer(s, &pb.Producerpc{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
