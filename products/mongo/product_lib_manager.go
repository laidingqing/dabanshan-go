package mongo

import (
	"github.com/laidingqing/dabanshan/common/config"
	"github.com/laidingqing/dabanshan/products/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ProductLibManager user manager struct
type ProductLibManager struct {
	session *mgo.Session
}

// NewProductLibManager new manager
func NewProductLibManager() *ProductLibManager {
	session, err := mgo.Dial(config.Database.HostURI)
	if err != nil {
		fatalError(err)
	}
	return &ProductLibManager{
		session: session,
	}
}

// CopySession for more flexible use
func (plm *ProductLibManager) CopySession() *mgo.Session {
	return plm.session.Copy()
}

//Find find all match word's productlibs.
func (plm *ProductLibManager) Find(word string) ([]model.ProductLib, error) {
	copySession := plm.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(ProductLibsCollectionName)
	var libs []model.ProductLib
	err := coll.Find(bson.M{"name": bson.M{"$regex": bson.RegEx{Pattern: word}}}).All(&libs)
	if err != nil {
		return nil, err
	}
	return libs, nil
}
