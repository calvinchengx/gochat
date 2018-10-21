package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/calvinchengx/gochat/pb"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:12021", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	} else {
		log.Printf("Successfully connected.")
	}
	defer conn.Close()

	// create the client
	c := pb.NewChatClient(conn)
	ctx := context.Background()
	stream, serr := c.RouteChat(ctx)
	if serr != nil {
		log.Fatalf("Stream error %v", err)
	}

	// accept chat message text from stdin
	reader := bufio.NewReader(os.Stdin)
	log.Printf("Send your message...")
	for {
		text, _ := reader.ReadString('\n')
		stream.Send(&pb.ChatMessage{
			Sender:   "alice",
			Receiver: "bob",
			Message:  text,
		})
	}

}
