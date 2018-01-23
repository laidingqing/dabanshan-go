package mongo

import (
	"errors"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	//ErrNoFoundEpisode 存在未审核的信息错误
	ErrNoFoundEpisode = errors.New("未发现的数据")
)

var session *mgo.Session
var episodeCollectionName = "episodes"
var episodeItemCollectionName = "episode_items"
var offerCollectionName = "offers"

func fatalError(err error) {
	log.Printf("mongodb error")
	os.Exit(1)
}
