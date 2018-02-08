package mongo

import (
	"time"

	"github.com/laidingqing/dabanshan/common/config"
	"github.com/laidingqing/dabanshan/products/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ProductItemManager user manager struct
type ProductItemManager struct {
	session *mgo.Session
}

// NewProductItemManager new manager
func NewProductItemManager() *ProductItemManager {
	session, err := mgo.Dial(config.Database.HostURI)
	if err != nil {
		fatalError(err)
	}
	return &ProductItemManager{
		session: session,
	}
}

// CopySession for more flexible use
func (pim *ProductItemManager) CopySession() *mgo.Session {
	return pim.session.Copy()
}

//Insert add a product item.
func (pim *ProductItemManager) Insert(item model.ProductItem) (string, error) {
	copySession := pim.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(ProductItemsCollectionName)
	item.ID = bson.NewObjectId()
	item.CreatedAt = time.Now()
	err := coll.Insert(item)
	if err != nil {
		return "", err
	}
	return item.ID.Hex(), nil
}

//FindAll find all product itmes by params
func (pim *ProductItemManager) FindAll(keyword string, store string, category string, offset int64, limit int64) ([]model.ProductItem, error) {
	params := bson.M{}
	if keyword != "" {
		params["name"] = keyword
	}
	if store != "" {
		params["accountId"] = store
	}
	if category != "" {
		params["categoryId"] = category
	}
	var products []model.ProductItem
	copySession := pim.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(ProductItemsCollectionName)
	coll.Find(params).All(&products)

	return products, nil
}
