package grpc

import (
	"context"

	pb "github.com/Wachayathorn/grpc-beginner/users/pb/service"
)

// UserHandler implements the gRPC UserService interface
type UserHandler struct {
	pb.UnimplementedUserServiceServer
}

// NewUserHandler creates a new gRPC user handler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// ListUsers returns a simple OK response
func (h *UserHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	// Return empty response with OK status
	return &pb.ListUsersResponse{
		Users:    []*pb.User{},
		Total:    0,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
