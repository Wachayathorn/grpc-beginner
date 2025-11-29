package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Wachayathorn/grpc-beginner/addresses/client"
	"github.com/Wachayathorn/grpc-beginner/users/business"
	"github.com/Wachayathorn/grpc-beginner/users/handler"
	pb "github.com/Wachayathorn/grpc-beginner/users/pb/proto"
)

const (
	port = ":50051"
)

func main() {
	addressClient := client.New()

	userBusiness := business.New(addressClient)

	userHandler := handler.NewUserHandler(userBusiness)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("🚀 Users gRPC Server running on %s\n", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
