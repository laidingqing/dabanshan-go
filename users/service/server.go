package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/laidingqing/dabanshan/common/util"
	"github.com/laidingqing/dabanshan/pb"
	"github.com/laidingqing/dabanshan/users/model"
	"github.com/laidingqing/dabanshan/users/mongo"
)

// RPCUserServer is used to implement user_service.UserServiceServer.
type RPCUserServer struct{}

var (
	manager = mongo.NewUserManager()
)

// CreateUser implements user_service.UserServiceServer
func (s *RPCUserServer) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Printf("Receive is %s\n", time.Now())
	rev, err := manager.Insert(model.User{
		UserName: request.Username,
		Password: util.CalculatePassHash(request.Password, request.Username),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{User: &pb.User{
		Id:       rev.ID.Hex(),
		Username: rev.UserName,
	}}, nil
}

// GetUser implements user_service.UserServiceServer
func (s *RPCUserServer) GetUser(context context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Get user account for: %v", request.Id)
	return &pb.GetUserResponse{User: &pb.User{
		Username: "laidingqing",
	}}, nil
}

// GetByUsername implements user_service.UserServiceServer
func (s *RPCUserServer) GetByUsername(context context.Context, request *pb.GetByUsernameRequest) (*pb.GetByUsernameResponse, error) {
	log.Printf("Get user by username for: %v", request.Username)
	return &pb.GetByUsernameResponse{User: &pb.User{
		Username: "laidingqing",
	}}, nil
}
