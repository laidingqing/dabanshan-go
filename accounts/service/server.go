package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/laidingqing/dabanshan/accounts/model"
	"github.com/laidingqing/dabanshan/accounts/mongo"
	"github.com/laidingqing/dabanshan/common/util"
	"github.com/laidingqing/dabanshan/pb"
)

// RPCAccountServer is used to implement user_service.UserServiceServer.
type RPCAccountServer struct{}

var (
	accountManager = mongo.NewAccountManager()
	authManager    = mongo.NewAuthInfoManager()
)

// CreateAccount implements account_service.UserServiceServer
func (s *RPCAccountServer) CreateAccount(context context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	fmt.Printf("Receive is %s\n", time.Now())
	rev, err := accountManager.Insert(model.Account{
		UserName: request.Username,
		Password: util.CalculatePassHash(request.Password, request.Username),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountResponse{Account: &pb.Account{
		Id:       rev.ID.Hex(),
		Username: rev.UserName,
	}}, nil
}

// GetAccount implements account_service.UserServiceServer
func (s *RPCAccountServer) GetAccount(context context.Context, request *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	log.Printf("Get user account for: %v", request.Id)
	return &pb.GetAccountResponse{Account: &pb.Account{
		Username: "laidingqing",
	}}, nil
}

// GetByUsername implements account_service.UserServiceServer
func (s *RPCAccountServer) GetByUsername(context context.Context, request *pb.GetByUsernameRequest) (*pb.GetByUsernameResponse, error) {
	log.Printf("Get user by username for: %v", request.Username)
	return &pb.GetByUsernameResponse{Account: &pb.Account{
		Username: "laidingqing",
	}}, nil
}

//CreateAuthInfo implements account_service.CreateAuthInfo
func (s *RPCAccountServer) CreateAuthInfo(context context.Context, request *pb.CreateAuthInfoRequest) (*pb.CreateAuthInfoResponse, error) {
	var reqAuth = request.AuthInfo
	rev, err := authManager.Insert(encodeAuthInfo(reqAuth))

	log.Printf("created auth info for : %v", rev)

	if err != nil {
		return nil, err
	}

	return &pb.CreateAuthInfoResponse{
		Created: true,
	}, nil
}

//CheckAuthInfo 审核认证信息
func (s *RPCAccountServer) CheckAuthInfo(ctx context.Context, in *pb.CreateCheckAuthInfoRequest) (*pb.CreateCheckAuthInfoResponse, error) {

	return &pb.CreateCheckAuthInfoResponse{
		Result: &pb.CheckResult{
			Message: "",
		},
	}, nil
}

//GetCurrentAuthInfo 获取当前审核通过认证信息
func (s *RPCAccountServer) GetCurrentAuthInfo(ctx context.Context, in *pb.GetAuthInfoRequest) (*pb.GetAuthInfoResponse, error) {
	return &pb.GetAuthInfoResponse{
		Info: &pb.AuthInfo{},
	}, nil
}
