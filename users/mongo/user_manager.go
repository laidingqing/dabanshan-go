package mongo

import (
	"log"
	"os"
	"time"

	"github.com/laidingqing/dabanshan/common/config"
	"github.com/laidingqing/dabanshan/users/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserManager user manager struct
type UserManager struct {
	session *mgo.Session
}

var session *mgo.Session
var userCollectionName = "users"

func fatalError(err error) {
	log.Printf("mongodb error")
	os.Exit(1)
}

// NewUserManager new manager
func NewUserManager() *UserManager {
	session, err := mgo.Dial(config.Database.HostURI)
	defer session.Close()
	if err != nil {
		fatalError(err)
	}
	session.DB(config.Database.DatabaseName)
	return &UserManager{
		session: session,
	}
}

// CopySession for more flexible use
func (um *UserManager) CopySession() *mgo.Session {
	return um.session.Copy()
}

//findByUserName 根据用户名查询用户
func (um *UserManager) Insert(user model.User) (model.User, error) {
	copySession := um.session.Copy()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(userCollectionName)
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	query.Insert(user)
	return user, nil
}

//findByUserName 根据用户名查询用户
func (um *UserManager) FindByUserName(username string) (interface{}, error) {
	copySession := um.session.Copy()
	defer copySession.Close()
	var user model.User
	query := copySession.DB(config.Database.DatabaseName).C(userCollectionName)
	query.Find(bson.M{"username": username}).One(&user)
	return user, nil
}
