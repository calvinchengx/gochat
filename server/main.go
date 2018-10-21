package main

import (
	"fmt"
	"io"
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

// ListenToClient listens on the incoming stream for any messages. It adds those messages to the channel. It doesn't return anything.
func ListenToClient(stream pb.Chat_RouteChatServer, messages chan<- pb.ChatMessage) {

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// do nothing
		}
		if err != nil {
			// do nothing
		} else {
			log.Printf("[ListenToClient] Client " + msg.Sender + " sent " + msg.Receiver + " a message: " + msg.Message)
			messages <- *msg
		}
	}

}

// RouteChat implements the RouteChat logic
func (s *server) RouteChat(stream pb.Chat_RouteChatServer) error {

	msg, err := stream.Recv()

	if err != nil {
		return err
	}

	log.Printf("[RouteChat]: Client " + msg.Sender + " sent " + msg.Receiver + " a message: " + msg.Message)

	outbox := make(chan pb.ChatMessage, 100)

	go ListenToClient(stream, outbox)

	for {
		select {
		case outMsg := <-outbox:
			// Broadcast
			log.Println(outMsg)
		}
	}

}

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
