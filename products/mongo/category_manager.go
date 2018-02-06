package mongo

import (
	"time"

	"github.com/laidingqing/dabanshan/common/config"
	"github.com/laidingqing/dabanshan/products/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//CategoryManager user manager struct
type CategoryManager struct {
	session *mgo.Session
}

// NewCagegoryManager new manager
func NewCagegoryManager() *CategoryManager {
	session, err := mgo.Dial(config.Database.HostURI)
	if err != nil {
		fatalError(err)
	}
	return &CategoryManager{
		session: session,
	}
}

// CopySession for more flexible use
func (cm *CategoryManager) CopySession() *mgo.Session {
	return cm.session.Copy()
}

//Remove remove a category entry
func (cm *CategoryManager) Remove(id string) error {
	copySession := cm.CopySession()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(CategoryCollectionName)
	err := query.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}

//RemoveAll remove all categories data
func (cm *CategoryManager) RemoveAll() error {
	copySession := cm.CopySession()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(CategoryCollectionName)
	_, err := query.RemoveAll(bson.M{})
	if err != nil {
		return err
	}
	return nil
}

//Insert insert a category entry
func (cm *CategoryManager) Insert(category model.Category) (model.Category, error) {
	copySession := cm.CopySession()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(CategoryCollectionName)
	category.ID = bson.NewObjectId()
	category.CreatedAt = time.Now()
	err := query.Insert(category)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

//FindAll find all category by parent ..
func (cm *CategoryManager) FindAll(parentID string) ([]model.Category, error) {
	copySession := cm.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(CategoryCollectionName)
	var categories []model.Category
	err := coll.Find(bson.M{"parent": parentID}).Sort("seq").All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
