package business

import (
	addresspb "github.com/Wachayathorn/grpc-beginner/addresses/pb/proto"
	userpb "github.com/Wachayathorn/grpc-beginner/users/pb/proto"
)

type ListUsersWithAddressesResponse struct {
	Users     *userpb.ListUsersResponse
	Addresses *addresspb.ListAddressesResponse
}
