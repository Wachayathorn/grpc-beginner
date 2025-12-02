package business

import (
	addresspb "github.com/wachayathorn/grpc-beginner/addresses/pb/proto"
	userpb "github.com/wachayathorn/grpc-beginner/users/pb/proto"
)

type ListUsersWithAddressesResponse struct {
	Users     *userpb.ListUsersResponse
	Addresses *addresspb.ListAddressesResponse
}
