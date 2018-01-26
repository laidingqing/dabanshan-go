package controllers

import (
	"log"

	model "github.com/laidingqing/dabanshan/episodes/model"
	"github.com/laidingqing/dabanshan/pb"
)

//DecodeEpisode ...
func DecodeEpisode(episode model.Episode) *pb.CreateEpisodeRequest {
	var expire = episode.Expire
	v, _ := expire.MarshalJSON()
	log.Printf("expire: %s", v)
	return &pb.CreateEpisodeRequest{
		Episode: &pb.Episode{
			Head: &pb.EpisodeHead{
				Name: episode.Name,
			},
			Items: DecodeEpisodeItems(episode.Items),
		},
	}
}

//DecodeEpisodeItems ...
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
