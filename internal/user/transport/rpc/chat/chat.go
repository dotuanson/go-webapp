package chat

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message from client: %s", message.Body)
	return &Message{Body: "Hello from server"}, nil
}
