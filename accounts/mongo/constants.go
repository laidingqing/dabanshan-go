package mongo

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session
var authCollectionName = "auths"
var accountCollectionName = "accounts"

func fatalError(err error) {
	log.Printf("mongodb error")
	os.Exit(1)
}
