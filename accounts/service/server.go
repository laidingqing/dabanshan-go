package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/laidingqing/dabanshan/accounts/model"
	"github.com/laidingqing/dabanshan/accounts/mongo"
	"github.com/laidingqing/dabanshan/common/util"
	pb "github.com/laidingqing/dabanshan/pb"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RPCAccountServer is used to implement user_service.UserServiceServer.
type RPCAccountServer struct{}

var (
	accountManager = mongo.NewAccountManager()

	authManager = mongo.NewAuthInfoManager()
)

// CreateAccount implements account_service.UserServiceServer
func (s *RPCAccountServer) CreateAccount(context context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	fmt.Printf("Receive is %s\n", time.Now())
	rev, err := accountManager.Insert(model.Account{
		UserName: request.Username,
		Password: util.CalculatePassHash(request.Password, request.Username),
		Salt:     request.Username,
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
	res, err := accountManager.FindByUserName(request.Username)
	if err != nil {
		return &pb.GetByUsernameResponse{}, err
	}
	return &pb.GetByUsernameResponse{Account: &pb.Account{
		Id:       res.ID.Hex(),
		Username: res.UserName,
		Password: res.Password,
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

//UpdateToken 更新当前会话TOKEN
func (s *RPCAccountServer) UpdateToken(ctx context.Context, in *pb.UpdateSessionTokenRequest) (*pb.UpdateSessionTokenResonse, error) {
	err := accountManager.UpdateCurrentToken(in.Accountid, in.Token)
	if err != nil {
		log.Printf("UpdateCurrentToken Err: %s", err.Error())
		return &pb.UpdateSessionTokenResonse{Updated: false}, err
	}
	return &pb.UpdateSessionTokenResonse{Updated: true}, nil
}

//GetByToken get Account by token
func (s *RPCAccountServer) GetByToken(ctx context.Context, in *pb.GetByTokenRequest) (*pb.GetByTokenResponse, error) {
	acct, err := accountManager.FindAccountByToken(in.Token)
	if err != nil {
		return &pb.GetByTokenResponse{}, err
	}
	return &pb.GetByTokenResponse{
		Account: dencodeAccountInfo(acct)}, nil
}

//FollowUser 关注用户
func (s *RPCAccountServer) FollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserResponse, error) {
	account, err := accountManager.FindByID(in.Accountid)
	if err != nil {
		return nil, err
	}
	flowAcct, err := accountManager.FindByID(in.Followid)
	if err != nil {
		return nil, err
	}

	log.Printf("found account: %s, flow acct: %s", account.ID, flowAcct.ID)
	var follows = []*model.Follow{}
	if err := accountManager.FindFollows(account.AccountID, flowAcct.AccountID, follows); err != nil {
		return nil, err
	}

	if len(follows) > 0 {
		return &pb.FollowUserResponse{Followed: true}, nil
	}
	rev, err := accountManager.InsertFollow(model.Follow{
		AccountID: mgo.DBRef{Collection: mongo.AccountCollectionName, Id: account.ID, Database: mongo.DatabaseName},
		FollowID:  mgo.DBRef{Collection: mongo.AccountCollectionName, Id: flowAcct.ID, Database: mongo.DatabaseName},
		CreatedAt: time.Now(),
	})
	log.Printf("new follow[%s] by acct id: %s", rev, account.AccountID)
	if err != nil {
		return nil, err
	}
	return &pb.FollowUserResponse{Followed: true}, nil
}

//GetFollows get follows by account
func (s *RPCAccountServer) GetFollows(ctx context.Context, in *pb.GetFollowsRequest) (*pb.GetFollowsResponse, error) {
	res, err := accountManager.FindAccountFollows(in.Accountid)
	if err != nil {
		return nil, err
	}

	var followAccts = []*pb.Account{}
	for i := range res {
		followAccts = append(followAccts, dencodeAccountInfo(res[i].Follow))
	}

	return &pb.GetFollowsResponse{
		Accounts: followAccts,
	}, nil
}

//CreateInterests create interests by account
func (s *RPCAccountServer) CreateInterests(ctx context.Context, in *pb.CreateTagsRequest) (*pb.CreateTagsResonse, error) {
	for i := range in.Tags {
		rev, err := accountManager.UpdateTagName(in.Tags[i].Name)
		if err == nil {
			res, _ := accountManager.CreateInterest(model.Interest{
				AccountID: mgo.DBRef{Collection: mongo.AccountCollectionName, Id: in.Accountid},
				TagID:     mgo.DBRef{Collection: mongo.TagsCollectionName, Id: bson.ObjectIdHex(rev)},
			})
			log.Printf("create interest rev: %s", res)
		} else {
			log.Printf("update tag err: %s", err.Error())
		}
	}
	return &pb.CreateTagsResonse{Created: true}, nil
}

//GetInterests get all interests of tags
func (s *RPCAccountServer) GetInterests(ctx context.Context, in *pb.GetTagsRequest) (*pb.GetTagsResponse, error) {

	return &pb.GetTagsResponse{}, nil
}
