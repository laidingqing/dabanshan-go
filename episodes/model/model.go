package model

import (
	"time"

	"github.com/laidingqing/dabanshan/common/util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EpisodeStatus Episode Offer Status
type EpisodeStatus int

const (
	_ EpisodeStatus = iota
	//NORMAL 正常
	NORMAL
	//EXPIRED 过期的不可报价
	EXPIRED
)

//EpisodeFeedStatus 关注的Episode状态
type EpisodeFeedStatus int

const (
	_ EpisodeFeedStatus = iota
	//CREATED 新建
	CREATED
	//OFFERED 已报
	OFFERED
)

//Episode 发布的需求(销售)事件
type Episode struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name,omitempty"`
	Expire    util.JsonTime `bson:"expire" json:"expire,omitempty"`
	AccountID string        `bson:"accountId" json:"accountId"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"cupdatedAt" json:"cupdatedAt,omitempty"`
	Items     []EpisodeItem `bson:"-" json:"items,omitempty"`
	Status    EpisodeStatus `bson:"status" json:"status,omitempty"`
}

//EpisodeItem 明细项
type EpisodeItem struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	AccountID string        `bson:"accountId" json:"accountId"`
	EpisodeID string        `bson:"episodeId" json:"episodeId"`
	ProductID string        `bson:"productid" json:"productid"`
	Name      string        `bson:"name" json:"name"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt,omitempty"`
	Weight    float32       `bson:"weight" json:"weight,omitempty"`
	Quantity  float32       `bson:"quantity" json:"quantity,omitempty"`
	Unit      string        `bson:"unit" json:"unit,omitempty"`
	TagIDs    []mgo.DBRef   `bson:"tags" json:"tags,omitempty"`
}

// EpisodeFeed ..关注的Episode Feed流
type EpisodeFeed struct {
	ID        bson.ObjectId     `bson:"_id" json:"id"`
	AccountID string            `bson:"accountId" json:"accountId"`
	EpisodeID mgo.DBRef         `bson:"episodeId" json:"episodeId"`
	Status    EpisodeFeedStatus `bson:"status" json:"status"`
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
