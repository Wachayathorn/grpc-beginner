package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wachayathorn/grpc-beginner/addresses/handler"
	pb "github.com/wachayathorn/grpc-beginner/addresses/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

func main() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	addressHandler := handler.NewAddressHandler()
	pb.RegisterAddressServiceServer(grpcServer, addressHandler)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("🚀 Users gRPC Server is running on port %s\n", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
