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

func EncodeAccounts(accts []*pb.Account) []model.Account {
	var accounts []model.Account
	for i := range accts {
		accounts = append(accounts, model.Account{
			AccountID: accts[i].Id,
			UserName:  accts[i].Username,
		})
	}
	return accounts
}
