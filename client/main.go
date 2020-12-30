package main

import (
	"context"
	"log"
	"time"

	user "github.com/sofiukl/explore-grpc/user"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9098"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SaveUser(ctx, &user.UserRequest{Name: "Sofikul", Age: "20", MobileNo: "901XXX31X6", Email: "user.email.com"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
