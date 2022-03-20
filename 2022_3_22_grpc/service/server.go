package service

import (
	"context"
	"errors"
	"fmt"

	pb "test_environment/2022_3_22_grpc/proto"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := getUser(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: user,
	}, nil
}

func getUser(id int64) (*pb.User, error) {

	switch id {
	case int64(1):
		return &pb.User{
			Id:      id,
			Name:    "Sazae",
			Age:     int64(24),
			Comment: "Hello! I'm Sazae!",
		}, nil
	case int64(2):
		return &pb.User{
			Id:      id,
			Name:    "Katsuo",
			Age:     int64(11),
			Comment: "Hello! I'm Katsuo!",
		}, nil
	case int64(3):
		return &pb.User{
			Id:      id,
			Name:    "Tara",
			Age:     int64(3),
			Comment: "Hello! I'm Tara!",
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Error: %d does not exist", id))
}
