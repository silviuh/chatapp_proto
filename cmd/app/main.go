package main

import (
	pb "chat_app/api/proto/gen" // Import the pb package
	"chat_app/internal/chat"
	"chat_app/pkg/db"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	// Create a new gRPC server
	mongoURI := "mongodb://localhost:27017" // Your MongoDB URI
	dbName := "chatapp"                     // Your database name
	db.ConnectMongo(mongoURI, dbName)

	db.InitializeRedis("localhost:6379")

	// Verify the connection (Optional, as ConnectMongo already pings the database)
	if err := db.MI.Client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	} else {
		log.Println("Connected to MongoDB successfully.")
	}

	s := grpc.NewServer()

	// Register your service with the gRPC server
	pb.RegisterChatServiceServer(s, &chat.Server{}) // Use pb.RegisterChatServiceServer

	// Register the reflection service on the gRPC server.
	reflection.Register(s)

	// Listen on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
