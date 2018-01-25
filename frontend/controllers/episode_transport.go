package controllers

import (
	model "github.com/laidingqing/dabanshan/episodes/model"
	"github.com/laidingqing/dabanshan/pb"
)

func DecodeEpisode(episode model.Episode) *pb.CreateEpisodeRequest {
	return &pb.CreateEpisodeRequest{
		Episode: &pb.Episode{
			Head: &pb.EpisodeHead{
				Name:   episode.Name,
				Expire: episode.Expire.Format("yyyy-MM-dd"),
			},
			Items: DecodeEpisodeItems(episode.Items),
		},
	}
}

func DecodeEpisodeItems(items []model.EpisodeItem) []*pb.EpisodeItem {
	var pEpisodeItems []*pb.EpisodeItem

	for i := range items {
		var item = &pb.EpisodeItem{
			Productid: items[i].ProductID,
			Name:      items[i].Name,
			Weight:    items[i].Weight,
			Quantity:  items[i].Quantity,
		}
		pEpisodeItems = append(pEpisodeItems, item)
	}
	return pEpisodeItems
}
