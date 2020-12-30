package main

import (
	"log"
	"net"

	user "github.com/sofiukl/explore-grpc/user"
	"google.golang.org/grpc"
)

const (
	port = ":9098"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, &user.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
