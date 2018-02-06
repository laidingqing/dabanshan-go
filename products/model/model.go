package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//ActivityType 活动类型: 直接折扣
type ActivityType int

const (
	_ ActivityType = iota
	//SALES 折扣
	SALES
)

//Activity ...
type Activity struct {
	ID           bson.ObjectId          `bson:"_id" json:"id,omitempty"`
	AccountID    string                 `bson:"accoutId" json:"accoutId,omitempty"`
	Name         string                 `bson:"name" json:"name"`
	ActivityType ActivityType           `bson:"activityType" json:"activityType,omitempty"`
	Props        map[string]interface{} `bson:"props" json:"props,omitempty"` //规则、参与方式（全场/部份参与/部分不参与）等
	CreatedAt    time.Time              `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt    time.Time              `bson:"updatedAt" json:"updatedAt,omitempty"`
}

//Category 商品库分类
type Category struct {
	ID         bson.ObjectId `bson:"_id" json:"id,omitempty"`
	CategoryID string        `bson:"-" json:"categoryID"`
	Name       string        `bson:"name" json:"name"`
	Parent     string        `bson:"parent" json:"parent"`
	Seq        int16         `bson:"seq" json:"seq"`
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	HTML       string        `bson:"-" json:"-,omitempty"`
}

//ProductLib 商品库，保存所有商品基本信息，用于查询及快速录入商品。
type ProductLib struct {
	ID         bson.ObjectId `bson:"_id" json:"id,omitempty"`
	ProductID  string        `bson:"-" json:"productId"`
	CategoryID string        `bson:"categoryID" json:"categoryID"`
	Name       string        `bson:"name" json:"name,omitempty"`
	SKU        string        `bson:"sku" json:"sku,omitempty"`
	Tags       []string      `bson:"tags" json:"tags,omitempty"`
	Unit       string        `bson:"unit" json:"unit,omitempty"`
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
}

//ProductItem 销售中的商品项
type ProductItem struct {
	ID         bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Name       string        `bson:"name" json:"name,omitempty"`
	SKU        string        `bson:"sku" json:"sku,omitempty"`
	AccountID  string        `bson:"accountId" json:"accountId,omitempty"`
	ImgURL     []string      `bson:"imgURL" json:"imgURL,omitempty"`
	CategoryID string        `bson:"categoryID" json:"categoryID"`
	BasicPrice float64       `bson:"basicPrice" json:"basicPrice,omitempty"`
	LastPrice  float64       `bson:"lastPrice" json:"lastPrice,omitempty"`
	Unit       string        `bson:"unit" json:"unit,omitempty"`
	Activity   []Activity    `bson:"-" json:"activities,omitempty"`
	Status     int           `bson:"status" json:"status,omitempty"`
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt  time.Time     `bson:"updatedAt" json:"updatedAt,omitempty"`
}

//WishList 愿望清单
type WishList struct {
}

//WishListItem 愿望清单列表
type WishListItem struct {
}
