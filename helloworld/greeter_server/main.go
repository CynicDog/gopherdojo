package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/cynicdog/grpc-prep/proto/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()

	// Create a TCP listener on the specified port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server instance
	s := grpc.NewServer()

	// Register the Greeter service implementation with the gRPC server
	helloworld.RegisterGreeterServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	// Start serving incoming connections; block here until server stops or error occurs
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
