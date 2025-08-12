package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/cynicdog/grpc-prep/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "World"
)

var (
	// Command-line flag for the gRPC server address, with a default value
	addr = flag.String("addr", "localhost:500051", "the address to connect to")
	// Command-line flag for the name to send in the greeting request
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	// Parse the command-line flags and assign their values
	flag.Parse()

	// Set up a connection to the gRPC server using insecure credentials (no TLS)
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client stub from the connection for the Greeter service
	c := pb.NewGreeterClient(conn)

	// Create a context with a 1-second timeout for the RPC call
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the SayHello method on the Greeter service with the provided name
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	// Log the greeting message received from the server
	log.Printf("Greeting: %s", r.GetMessage())
}
