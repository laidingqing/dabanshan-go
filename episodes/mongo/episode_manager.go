package mongo

import (
	"log"
	"time"

	"github.com/laidingqing/dabanshan/common/config"
	"github.com/laidingqing/dabanshan/episodes/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EpisodeManager ...
type EpisodeManager struct {
	session *mgo.Session
}

//NewEpisodeManager episode data manager
func NewEpisodeManager() *EpisodeManager {
	session, err := mgo.Dial(config.Database.HostURI)
	if err != nil {
		fatalError(err)
	}
	return &EpisodeManager{
		session: session,
	}
}

//CopySession ...
func (em *EpisodeManager) CopySession() *mgo.Session {
	return em.session.Copy()
}

//Insert insert a episode
func (em *EpisodeManager) Insert(episode model.Episode) (string, error) {
	copySession := em.session.Copy()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(episodeCollectionName)
	episode.ID = bson.NewObjectId()
	episode.CreatedAt = time.Now()
	query.Insert(episode)
	em.InsertItems(episode.ID.Hex(), episode.Items)
	return episode.ID.Hex(), nil
}

//InsertItems insert items by Episode
func (em *EpisodeManager) InsertItems(episodeID string, items []model.EpisodeItem) error {
	copySession := em.session.Copy()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(episodeItemCollectionName)

	for i := range items {
		items[i].ID = bson.NewObjectId()
		items[i].EpisodeID = episodeID
		items[i].CreatedAt = time.Now()
		err := coll.Insert(items[i])
		if err != nil {
			log.Printf("insert items: %s", err.Error())
		}
	}

	return nil
}

//FindByID find episode by id
func (em *EpisodeManager) FindByID(id string) (model.Episode, error) {
	var episode model.Episode
	copySession := em.session.Copy()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(episodeCollectionName)
	coll.FindId(bson.ObjectIdHex(id)).One(&episode)
	return episode, nil
}

//FindEpisodeItemsByID find episode items by id
func (em *EpisodeManager) FindEpisodeItemsByID(episodeID string) ([]model.EpisodeItem, error) {
	var items []model.EpisodeItem
	copySession := em.session.Copy()
	defer copySession.Close()
	var query = bson.M{"episodeId": episodeID}
	coll := copySession.DB(config.Database.DatabaseName).C(episodeCollectionName)
	coll.Find(query).One(&items)
	return items, nil
}

//FindEpisodeItemsByAcctID find episode items by id
func (em *EpisodeManager) FindEpisodeItemsByAcctID(episodeID string, accountID string) ([]model.EpisodeItem, error) {
	var items []model.EpisodeItem
	copySession := em.session.Copy()
	defer copySession.Close()
	var query = bson.M{"episodeId": episodeID, "accountId": accountID}
	coll := copySession.DB(config.Database.DatabaseName).C(episodeCollectionName)
	coll.Find(query).One(&items)
	return items, nil
}

//FindEpisodesByAcctID 根据账号查询发布的Episode
func (em *EpisodeManager) FindEpisodesByAcctID(accountID string, page int, size int) (model.PagnationEpisode, error) {
	var episodes []model.Episode
	copySession := em.session.Copy()
	defer copySession.Close()
	var bson = bson.M{"accountId": accountID}
	coll := copySession.DB(config.Database.DatabaseName).C(episodeCollectionName)
	q := coll.Find(bson).Skip(page).Limit(size)
	countQuery := coll.Find(bson)
	total, err := countQuery.Count()
	if err != nil {
		return model.PagnationEpisode{}, err
	}
	q.All(episodes)

	return model.PagnationEpisode{
		Data:       episodes,
		TotalCount: total,
	}, nil
}

//UpdateEpisodeWithExpire ..
func (em *EpisodeManager) UpdateEpisodeWithExpire(episodeID string) error {
	copySession := em.session.Copy()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(episodeCollectionName)
	coll.UpdateId(bson.ObjectIdHex(episodeID), bson.M{"$set": bson.M{
		"offerStatus": model.EXPIRED,
	}})
	return nil
}
