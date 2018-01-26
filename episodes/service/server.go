package service

import (
	"context"
	"log"

	"github.com/laidingqing/dabanshan/episodes/mongo"
	"github.com/laidingqing/dabanshan/pb"
)

// RPCEpisodeServer is used to implement user_service.UserServiceServer.
type RPCEpisodeServer struct{}

var (
	episodeManager = mongo.NewEpisodeManager()
	offerManager   = mongo.NewOfferManager()
)

//CreateEpisode create a episode entry
func (s *RPCEpisodeServer) CreateEpisode(context context.Context, request *pb.CreateEpisodeRequest) (*pb.CreateEpisodeResponse, error) {
	rev, err := episodeManager.Insert(DecodeEpisode(request.Episode))
	if err != nil {
		return nil, err
	}
	log.Printf("saved episode, rev: %s ", rev)
	//TODO 分发供货消息
	return &pb.CreateEpisodeResponse{
		Revid: rev,
	}, nil
}
