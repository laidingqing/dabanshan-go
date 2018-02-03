package service

import (
	"log"
	"time"

	"github.com/laidingqing/dabanshan/common/util"
	"github.com/laidingqing/dabanshan/episodes/model"
	"github.com/laidingqing/dabanshan/pb"
)

//DecodeEpisode ..
func DecodeEpisode(th *pb.Episode) model.Episode {
	expire, err := time.Parse("2006-01-02 00:00:00", th.Head.Expire)
	if err != nil {
		log.Printf("parse time err: %s", err.Error())
	}
	return model.Episode{
		Name:      th.Head.Name,
		AccountID: th.Head.Accountid,
		Expire:    util.JsonTime(expire),
		Items:     DecodeEpisodeItems(th.Items),
	}
}

//DecodeEpisodeItems .
func DecodeEpisodeItems(items []*pb.EpisodeItem) []model.EpisodeItem {
	var episodeItems []model.EpisodeItem
	for i := range items {
		episodeItems = append(episodeItems, model.EpisodeItem{
			Name:     items[i].Name,
			Weight:   items[i].Weight,
			Quantity: items[i].Quantity,
			Unit:     items[i].Unit,
		})
	}
	return episodeItems
}
