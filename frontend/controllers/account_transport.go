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

//EncodeAccounts encode accounts
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

//DecodeTags model to pb for tags
func DecodeTags(tags []model.Tag) []*pb.Tag {
	var pbtags []*pb.Tag
	for i := range tags {
		pbtags = append(pbtags, &pb.Tag{
			Name: tags[i].Name,
		})
	}
	return pbtags
}
