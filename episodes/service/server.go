package service

import (
	"context"

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

	return &pb.CreateEpisodeResponse{}, nil
}
