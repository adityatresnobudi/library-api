package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adityatresnobudi/library-api/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up the gRPC client to connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewAuthClient(conn)

	// Prepare the request message
	req := &pb.LoginRequest{
		Email:    "natnat@xyz.com",
		Password: "natnat",
	}

	// Send the request to the server
	login, err := client.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send the request: %v", err)
	}

	// Print the server's response
	fmt.Println("Message:", login.Message)
	fmt.Println("Token:", login.Token)
}
