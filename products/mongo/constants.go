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

//CategoryCollectionName category table name
var CategoryCollectionName = "categories"

//ProductLibsCollectionName ...
var ProductLibsCollectionName = "product_libs"

//ProductItemsCollectionName ...
var ProductItemsCollectionName = "product_items"

func fatalError(err error) {
	log.Printf("mongodb error: %s", err.Error())
	os.Exit(1)
}
