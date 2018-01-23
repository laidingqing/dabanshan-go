package mongo

import (
	"time"

	"github.com/laidingqing/dabanshan/accounts/model"
	"github.com/laidingqing/dabanshan/common/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AccountManager user manager struct
type AccountManager struct {
	session *mgo.Session
}

// NewAccountManager new manager
func NewAccountManager() *AccountManager {
	session, err := mgo.Dial(config.Database.HostURI)
	defer session.Close()
	if err != nil {
		fatalError(err)
	}
	session.DB(config.Database.DatabaseName)
	return &AccountManager{
		session: session,
	}
}

// CopySession for more flexible use
func (um *AccountManager) CopySession() *mgo.Session {
	return um.session.Copy()
}

//Insert 根据用户名查询用户
func (um *AccountManager) Insert(user model.Account) (model.Account, error) {
	copySession := um.session.Copy()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(accountCollectionName)
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	query.Insert(user)
	return user, nil
}

//FindByUserName 根据用户名查询用户
func (um *AccountManager) FindByUserName(username string) (interface{}, error) {
	copySession := um.session.Copy()
	defer copySession.Close()
	var user model.Account
	query := copySession.DB(config.Database.DatabaseName).C(accountCollectionName)
	query.Find(bson.M{"username": username}).One(&user)
	return user, nil
}
