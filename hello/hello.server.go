package hello

import (
	context "context"
	"log"
)

// Server - hello server
type Server struct{}

// SayHello - SayHello method
func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello123 " + in.GetName()}, nil
}
