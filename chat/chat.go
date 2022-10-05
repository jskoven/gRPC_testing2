package chat

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SendFunction(ctx context.Context, msg *Message) (*Message, error) {

	log.Printf("Received message from client: %t", msg.Body)
	return &Message{Body: "Received message!"}, nil

}
