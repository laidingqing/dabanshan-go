package service

import (
	"github.com/laidingqing/dabanshan/accounts/model"
	pb "github.com/laidingqing/dabanshan/pb"
)

func encodeAccountInfo(request *pb.Account) model.Account {
	return model.Account{
		AccountID: request.Id,
		UserName:  request.Username,
	}
}

func dencodeAccountInfo(request model.Account) *pb.Account {
	return &pb.Account{
		Id:       request.ID.Hex(),
		Username: request.UserName,
	}
}

// transport auth info

func encodeAuthInfo(request *pb.AuthInfo) model.AuthInfo {
	return model.AuthInfo{
		AccountID: request.AccountId,
		Name:      request.Name,
		Address:   request.Address,
		Province:  request.Province,
		City:      request.City,
		County:    request.County,
		Passport:  request.Passport,
		Message:   request.Result.Message,
		// Result:    request.Result.Result,
	}
}
