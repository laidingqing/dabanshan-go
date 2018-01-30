package mongo

import (
	"log"
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
	if err != nil {
		fatalError(err)
	}
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
	copySession := um.CopySession()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	err := query.Insert(user)
	if err != nil {
		log.Printf("db manager err: %s", err.Error())
		return model.Account{}, err
	}
	return user, nil
}

//FindByID find account by id
func (um *AccountManager) FindByID(id string) (model.Account, error) {
	var acct model.Account
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	err := coll.FindId(bson.ObjectIdHex(id)).One(&acct)
	if err != nil {
		return model.Account{}, err
	}
	return acct, nil
}

//FindByUserName 根据用户名查询用户
func (um *AccountManager) FindByUserName(username string) (model.Account, error) {
	copySession := um.CopySession()
	defer copySession.Close()
	var user model.Account
	query := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	err := query.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return model.Account{}, err
	}
	return user, nil
}

//FindAccountByToken 根据当前TOKEN查询账号
func (um *AccountManager) FindAccountByToken(token string) (model.Account, error) {
	var acct model.Account
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	err := coll.Find(bson.M{"token": token}).One(&acct)
	if err != nil {
		return model.Account{}, err
	}
	return acct, nil
}

//UpdateCurrentToken 更新会话账号TOKEN
func (um *AccountManager) UpdateCurrentToken(id string, token string) error {
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	err := coll.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"token": token}})
	if err != nil {
		return err
	}
	return nil
}

//InsertFollow insert a follow
func (um *AccountManager) InsertFollow(follow model.Follow) (string, error) {
	copySession := um.CopySession()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(FollowsCollectionName)
	follow.ID = bson.NewObjectId()
	follow.CreatedAt = time.Now()
	err := query.Insert(follow)
	if err != nil {
		log.Printf("db manager err: %s", err.Error())
		return "", err
	}
	return follow.ID.Hex(), nil
}

//FindFollows find list follow by acctId
func (um *AccountManager) FindFollows(acctID string, followID string, follows []*model.Follow) error {
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(FollowsCollectionName)
	err := coll.Find(bson.M{"accountId": acctID, "followId": followID}).All(&follows)
	if err != nil {
		return err
	}
	return nil
}

//FindAccountFollows find follows by account
func (um *AccountManager) FindAccountFollows(acctID string) ([]*model.Follow, error) {
	var follows = []*model.Follow{}
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(FollowsCollectionName)
	err := coll.Find(bson.M{"accountId.$id": bson.ObjectIdHex(acctID)}).All(&follows)
	log.Printf("found follows: %d", len(follows))
	if err != nil {
		return nil, err
	}
	for i := range follows {
		var acct model.Account
		var follow model.Account
		copySession.DB(config.Database.DatabaseName).FindRef(&follows[i].AccountID).One(&acct)
		follows[i].Account = acct
		copySession.DB(config.Database.DatabaseName).FindRef(&follows[i].FollowID).One(&follow)
		follows[i].Follow = follow
	}

	return follows, nil
}

// Interest && Tag

//CreateInterest ...
func (um *AccountManager) CreateInterest(interest model.Interest) (string, error) {
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(InterestsCollectionName)
	interest.ID = bson.NewObjectId()
	interest.CreatedAt = time.Now()
	err := coll.Insert(interest)
	if err != nil {
		return "", err
	}
	return interest.ID.Hex(), nil
}

//UpdateTagName ...
func (um *AccountManager) UpdateTagName(name string) (string, error) {
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(TagsCollectionName)
	var tags []model.Tag
	coll.Find(bson.M{"name": name}).All(&tags)
	if len(tags) > 0 {
		return tags[0].ID.Hex(), nil
	}
	tag := model.Tag{Name: name, CreatedAt: time.Now()}
	tag.ID = bson.NewObjectId()
	err := coll.Insert(tag)
	if err != nil {
		return "", err
	}
	log.Printf("Upsert Tags: %s", tag.ID)
	return tag.ID.Hex(), err
}
