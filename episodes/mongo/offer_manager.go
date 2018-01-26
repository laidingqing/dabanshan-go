package mongo

import (
	"time"

	"github.com/laidingqing/dabanshan/common/config"
	"github.com/laidingqing/dabanshan/episodes/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//OfferManager ...
type OfferManager struct {
	session *mgo.Session
}

//NewOfferManager episode data manager
func NewOfferManager() *OfferManager {
	session, err := mgo.Dial(config.Database.HostURI)
	if err != nil {
		fatalError(err)
	}
	return &OfferManager{
		session: session,
	}
}

//CopySession ...
func (om *OfferManager) CopySession() *mgo.Session {
	return om.session.Copy()
}

//Insert insert offers
func (om *OfferManager) Insert(episodeID string, items []model.OfferItem) error {
	copySession := om.session.Copy()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(offerCollectionName)
	for i := range items {
		items[i].ID = bson.NewObjectId()
		items[i].EpisodeID = episodeID
		items[i].CreatedAt = time.Now()
		coll.Insert(items[i])
	}

	return nil
}

//UpdateOfferPrice update offer price
func (om *OfferManager) UpdateOfferPrice(offerID string, price float64) error {
	copySession := om.session.Copy()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(offerCollectionName)
	coll.UpdateId(bson.ObjectIdHex(offerID), bson.M{"$set": bson.M{
		"price": price,
	}})
	return nil
}
