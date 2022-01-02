package services

import (
	"context"

	"github.com/retatu/go-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User{
		id:    "123",
		name:  req.GetName(),
		email: req.GetEmail(),
	}, nil
}
