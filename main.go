package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/calvinchengx/gochat/pb"
)

const (
	port = ":12021"
)

type server struct{}

func main() {
	fmt.Println("gRPC chat server")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// initializes the gRPC server
	s := grpc.NewServer()

	// register the server with gRPC
	pb.RegisterChatServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
