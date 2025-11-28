package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcHandler "github.com/Wachayathorn/grpc-beginner/users/internal/delivery/grpc"
	pb "github.com/Wachayathorn/grpc-beginner/users/pb/proto"
)

const (
	port = ":50051"
)

func main() {
	// Create handler
	userHandler := grpcHandler.NewUserHandler()

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register service
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	// Enable reflection for grpcurl
	reflection.Register(grpcServer)

	// Start listening
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("🚀 Users gRPC Server is running on port %s\n", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
