package mongo

import (
	"github.com/laidingqing/dabanshan/common/config"
	"github.com/laidingqing/dabanshan/products/model"
	mgo "gopkg.in/mgo.v2"
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

	return "", nil
}
