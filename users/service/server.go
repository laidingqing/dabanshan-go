package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/laidingqing/dabanshan/pb"
)

// RPCUserServer is used to implement user_service.UserServiceServer.
type RPCUserServer struct{}

// CreateUser implements user_service.UserServiceServer
func (s *RPCUserServer) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Printf("Receive is %s\n", time.Now())
	return &pb.CreateUserResponse{User: &pb.User{
		Username: "laidingqing",
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
