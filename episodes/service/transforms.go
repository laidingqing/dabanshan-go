package service

import (
	"github.com/laidingqing/dabanshan/episodes/model"
	"github.com/laidingqing/dabanshan/pb"
)

//DecodeEpisode ..
func DecodeEpisode(th *pb.Episode) model.Episode {
	return model.Episode{
		Name: th.Head.Name,
	}
}
