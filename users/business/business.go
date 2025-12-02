package business

import (
	"context"

	"github.com/wachayathorn/grpc-beginner/addresses/client"
	addresspb "github.com/wachayathorn/grpc-beginner/addresses/pb/proto"
	userpb "github.com/wachayathorn/grpc-beginner/users/pb/proto"
)

type Business interface {
	ListUsersWithAddresses(ctx context.Context) (ListUsersWithAddressesResponse, error)
}

type business struct {
	addressClient client.Client
}

func New(
	addressClient client.Client,
) Business {
	return &business{
		addressClient,
	}
}

func (b *business) ListUsersWithAddresses(ctx context.Context) (ListUsersWithAddressesResponse, error) {
	addresses, err := b.addressClient.ListAddresses(ctx, &addresspb.ListAddressesRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		return ListUsersWithAddressesResponse{}, err
	}
	return ListUsersWithAddressesResponse{
		Users: &userpb.ListUsersResponse{
			Page:  0,
			Total: 0,
		},
		Addresses: addresses,
	}, nil
}
