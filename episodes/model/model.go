package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//EpisodeOfferStatus Episode Offer Status
type EpisodeOfferStatus int

const (
	_ EpisodeOfferStatus = iota
	//NORMAL 正常
	NORMAL
	//EXPIRED 过期的不可报价
	EXPIRED
)

//Episode 发布的需求(销售)事件
type Episode struct {
	ID          bson.ObjectId      `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Expire      time.Time          `bson:"expire" json:"expire,omitempty"`
	AccountID   string             `bson:"accountId" json:"accountId"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"cupdatedAt" json:"cupdatedAt,omitempty"`
	Items       []EpisodeItem      `bson:"-" json:"items,omitempty"`
	OfferStatus EpisodeOfferStatus `bson:"offerStatus" json:"offerStatus,omitempty"`
}

//EpisodeItem 明细项
type EpisodeItem struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	AccountID string        `bson:"accountId" json:"accountId"`
	EpisodeID string        `bson:"episodeId" json:"episodeId"`
	ProductID string        `bson:"productid" json:"productid"`
	Name      string        `bson:"name" json:"name"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"cupdatedAt" json:"cupdatedAt,omitempty"`
	Weight    float32       `bson:"weight" json:"weight,omitempty"`
	Quantity  float32       `bson:"quantity" json:"quantity,omitempty"`
}

//OfferItem offer items.
type OfferItem struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	AccountID     string        `bson:"accountId" json:"accountId"`
	EpisodeID     string        `bson:"episodeId" json:"episodeId"`
	EpisodeItemID string        `bson:"episodeItemId" json:"episodeItemId"`
	ProductID     string        `bson:"productid" json:"productid"`
	Name          string        `bson:"name" json:"name"`
	Price         float32       `bson:"price" json:"price"`
	CreatedAt     time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
}

//PagnationEpisode ...
type PagnationEpisode struct {
	Data        interface{} `bson:"-" json:"data"`
	TotalCount  int         `bson:"-" json:"totalCount"`
	CurrentPage int         `bson:"-" json:"currentPage"`
}
