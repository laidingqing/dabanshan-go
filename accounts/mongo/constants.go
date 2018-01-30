package mongo

import (
	"log"
	"os"

	"github.com/laidingqing/dabanshan/common/config"
	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

//DatabaseName ..
var DatabaseName = config.Database.DatabaseName

//AuthCollectionName auth table name
var AuthCollectionName = "auths"

//AccountCollectionName account table name
var AccountCollectionName = "accounts"

//FollowsCollectionName follow table name
var FollowsCollectionName = "follows"

//TagsCollectionName tags table name
var TagsCollectionName = "tags"

//InterestsCollectionName account's interest table name
var InterestsCollectionName = "interests"

func fatalError(err error) {
	log.Printf("mongodb error")
	os.Exit(1)
}
