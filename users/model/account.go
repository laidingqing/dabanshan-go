package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User 用户账号信息
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	UserName  string        `bson:"username" json:"username,omitempty"`
	Password  string        `bson:"password" json:"password,omitempty"`
	Salt      string        `bson:"salt" json:"-"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
}
