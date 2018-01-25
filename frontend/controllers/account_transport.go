package controllers

import (
	"github.com/laidingqing/dabanshan/accounts/model"
	"github.com/laidingqing/dabanshan/pb"
)

//EncodeAccount pb to model for Account
func EncodeAccount(acct *pb.Account) model.Account {
	return model.Account{
		AccountID: acct.Id,
		UserName:  acct.Username,
	}
}
