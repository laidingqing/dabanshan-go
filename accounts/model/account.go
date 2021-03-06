package model

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AuthenticationType 授权平台
type AuthenticationType int

const (
	_ AuthenticationType = iota
	//WEIXIN 微信开放平台
	WEIXIN
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
	AccountID   string        `bson:"-" json:"id"`
	UserName    string        `bson:"username" json:"username,omitempty"`
	Password    string        `bson:"password" json:"password,omitempty"`
	Salt        string        `bson:"salt" json:"-"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	AccountType []AccountType `bson:"accountType" json:"accountType,omitempty"`
	Token       string        `bson:"token" json:"-,omitempty"`
	UserInfoRef mgo.DBRef     `bson:"infoRef" json:"-"`
	UserInfo    UserInfo      `bson:"userInfo" json:"userInfo,omitempty"`
}

//Authentication 第三方授权认证登记表
type Authentication struct {
	ID        bson.ObjectId      `bson:"_id" json:"id"`
	AccountID string             `bson:"accountId" json:"accountId,omitempty"`
	OpenID    string             `bson:"openId" json:"openId,omitempty"`
	Type      AuthenticationType `bson:"plateform" json:"plateform,omitempty"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
}

//UserInfo ...
type UserInfo struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	AvatarURL string        `bson:"avatarURL" json:"avatarURL,omitempty"`
	City      string        `bson:"city" json:"city,omitempty"`
	Country   string        `bson:"country" json:"country,omitempty"`
	Gender    int           `bson:"gender" json:"gender,omitempty"`
	NickName  string        `bson:"nickName" json:"nickName,omitempty"`
	Province  string        `bson:"province" json:"province,omitempty"`
	AccountID string        `bson:"accountId" json:"accountId,omitempty"`
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

//Interest 账号关注的食材标签
type Interest struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	AccountID mgo.DBRef     `bson:"accountId" json:"accountId"`
	Account   Account       `bson:"-" json:"account,omitempty"`
	TagID     mgo.DBRef     `bson:"tagId" json:"tagId"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
}

//Tag all tags.
type Tag struct {
	ID        bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Name      string        `bson:"name" json:"name,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
}

//Follow 关注
type Follow struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	AccountID mgo.DBRef     `bson:"accountId" json:"accountId"`
	FollowID  mgo.DBRef     `bson:"followId" json:"followId,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	Account   Account       `bson:"-" json:"account,omitempty"`
	Follow    Account       `bson:"-" json:"follow,omitempty"`
}
