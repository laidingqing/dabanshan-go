package mongo

import (
	"github.com/laidingqing/dabanshan/common/config"
	mgo "gopkg.in/mgo.v2"
)

//ActivityManager user manager struct
type ActivityManager struct {
	session *mgo.Session
}

// NewActivityManager new manager
func NewActivityManager() *ActivityManager {
	session, err := mgo.Dial(config.Database.HostURI)
	if err != nil {
		fatalError(err)
	}
	// Ensure Index TODO
	return &ActivityManager{
		session: session,
	}
}
