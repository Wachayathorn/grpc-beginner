package handler

import (
	"context"

	pb "github.com/wachayathorn/grpc-beginner/addresses/pb/proto"
)

type AddressHandler struct {
	pb.UnimplementedAddressServiceServer
}

func NewAddressHandler() *AddressHandler {
	return &AddressHandler{}
}

func (h *AddressHandler) ListAddresses(ctx context.Context, req *pb.ListAddressesRequest) (*pb.ListAddressesResponse, error) {
	return &pb.ListAddressesResponse{
		Addresses: []*pb.Address{},
		Total:     0,
		Page:      req.Page,
		PageSize:  req.PageSize,
	}, nil
}
