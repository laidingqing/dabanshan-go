package service

import (
	"github.com/laidingqing/dabanshan/accounts/model"
	"github.com/laidingqing/dabanshan/pb"
)

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
