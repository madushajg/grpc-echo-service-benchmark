package hw

import (
	"fmt"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, in *Message) (*Message, error) {
	//fmt.Printf("Server received hello from %s \n", in.Body)
	message := fmt.Sprintf("Hello %s", in.Body)
	return &Message{Body: message}, nil
}
