package mongo

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session
var authCollectionName = "auths"
var accountCollectionName = "accounts"
var followsCollectionName = "follows"

func fatalError(err error) {
	log.Printf("mongodb error")
	os.Exit(1)
}
