package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Category 商品库分类
type Category struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Parent    string        `bson:"parent" json:"parent"`
	Seq       int16         `bson:"seq" json:"seq"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	HTML      string        `bson:"-" json:"html"`
}
