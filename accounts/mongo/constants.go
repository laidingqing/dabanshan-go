package mongo

import (
	"log"
	"os"

	"github.com/laidingqing/dabanshan/common/config"
	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session
var DatabaseName = config.Database.DatabaseName
var AuthCollectionName = "auths"
var AccountCollectionName = "accounts"
var FollowsCollectionName = "follows"

func fatalError(err error) {
	log.Printf("mongodb error")
	os.Exit(1)
}
