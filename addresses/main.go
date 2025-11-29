package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Wachayathorn/grpc-beginner/addresses/handler"
	pb "github.com/Wachayathorn/grpc-beginner/addresses/pb/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

func main() {
	grpcServer := grpc.NewServer()

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
