package client

import (
	"context"
	"log"

	pb "github.com/wachayathorn/grpc-beginner/addresses/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	ListAddresses(ctx context.Context, req *pb.ListAddressesRequest) (*pb.ListAddressesResponse, error)
}

type client struct {
	client pb.AddressServiceClient
}

func New() Client {
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	c := pb.NewAddressServiceClient(conn)
	return &client{client: c}
}

func (c *client) ListAddresses(ctx context.Context, req *pb.ListAddressesRequest) (*pb.ListAddressesResponse, error) {
	return c.client.ListAddresses(ctx, req)
}
