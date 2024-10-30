package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	address string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{address: addr}
}

func (s *gRPCServer) Run() error {
	listener, err := net.Listen("tcp", s.address)

	if err != nil {
		log.Fatalf("Failed to listen to the port %v : %v", s.address, err)
	}

	grpcServer := grpc.NewServer()

	//Services to inject...

	log.Println("Starting GRPC server in the port", s.address)

	return grpcServer.Serve(listener)

}
