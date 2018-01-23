package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//AuthCheckResult 认证审核结果
type AuthCheckResult int

const (
	_ AuthCheckResult = iota
	//CREATED 新建未处理
	CREATED
	//PASS 通过
	PASS
	//REJECT 拒绝
	REJECT
)

//AccountType 账号角色类型
type AccountType int

const (
	_ AccountType = iota
	//SALER 卖家，合作社
	SALER
	//BUYER 买家，中间商等
	BUYER
)

// Account 用户账号信息
type Account struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserName    string        `bson:"username" json:"username,omitempty"`
	Password    string        `bson:"password" json:"password,omitempty"`
	Salt        string        `bson:"salt" json:"-"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	AccountType []AccountType `bson:"accountType" json:"accountType,omitempty"`
}

//AuthInfo 用户认证信息
type AuthInfo struct {
	ID        bson.ObjectId   `bson:"_id" json:"id"`
	AccountID string          `bson:"accountId" json:"accountId"`
	Name      string          `bson:"name" json:"name"`
	Address   string          `bson:"address" json:"address,omitempty"`
	Province  string          `bson:"province" json:"province,omitempty"`
	City      string          `bson:"city" json:"city,omitempty"`
	County    string          `bson:"county" json:"county,omitempty"`
	Contact   string          `bson:"contact" json:"contact,omitempty"`
	Passport  string          `bson:"passport" json:"passport,omitempty"`
	Message   string          `bson:"message" json:"message,omitempty"`
	Result    AuthCheckResult `bson:"result" json:"result,omitempty"`
	CreatedAt time.Time       `bson:"createdAt" json:"createdAt,omitempty"`
}
