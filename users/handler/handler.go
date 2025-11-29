package handler

import (
	"context"
	"fmt"

	"github.com/Wachayathorn/grpc-beginner/users/business"
	pb "github.com/Wachayathorn/grpc-beginner/users/pb/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userBusiness business.Business
}

// NewUserHandler creates a new gRPC user handler
func NewUserHandler(userBusiness business.Business) *UserHandler {
	return &UserHandler{
		userBusiness: userBusiness,
	}
}

// ListUsers returns a simple OK response
func (h *UserHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	res, err := h.userBusiness.ListUsersWithAddresses(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("response : %v", res)

	return &pb.ListUsersResponse{
		Users:    []*pb.User{},
		Total:    0,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
