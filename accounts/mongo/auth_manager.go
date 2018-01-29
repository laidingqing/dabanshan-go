package mongo

import (
	"errors"
	"time"

	"github.com/laidingqing/dabanshan/accounts/model"
	"github.com/laidingqing/dabanshan/common/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	//ErrExistedInfo 存在未审核的信息错误
	ErrExistedInfo = errors.New("存在未审核的信息")
)

//AuthInfoManager auth manager struct
type AuthInfoManager struct {
	session *mgo.Session
}

// NewAuthInfoManager new manager
func NewAuthInfoManager() *AuthInfoManager {
	session, err := mgo.Dial(config.Database.HostURI)
	defer session.Close()
	if err != nil {
		fatalError(err)
	}
	session.DB(config.Database.DatabaseName)
	return &AuthInfoManager{
		session: session,
	}
}

// CopySession for more flexible use
func (am *AuthInfoManager) CopySession() *mgo.Session {
	return am.session.Copy()
}

//Insert 新增认证信息
func (am *AuthInfoManager) Insert(auth model.AuthInfo) (string, error) {
	infos, err := am.FindAuthByStatus(auth.AccountID, model.CREATED)
	if err != nil {
		return "", err
	}

	if len(infos) > 0 {
		return "", ErrExistedInfo
	}

	copySession := am.session.Copy()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	auth.ID = bson.NewObjectId()
	auth.CreatedAt = time.Now()
	query.Insert(auth)
	return auth.ID.Hex(), nil
}

//FindAuthByStatus 根据审核状态查询认证信息
func (am *AuthInfoManager) FindAuthByStatus(accountID string, status model.AuthCheckResult) ([]model.AuthInfo, error) {
	var authInfos []model.AuthInfo
	copySession := am.session.Copy()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	var query = bson.M{"accountId": accountID, "result": status}
	coll.Find(query).All(authInfos)
	return authInfos, nil
}
