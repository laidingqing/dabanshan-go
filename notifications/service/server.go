package service

import (
	"context"

	"github.com/laidingqing/dabanshan/pb"
)

// RPCNotifactionServer is used to implement user_service.UserServiceServer.
type RPCNotifactionServer struct{}

//Send send a notifaction message.
func Send(ctx context.Context, in *pb.CreateNotifactionRequest) (*pb.CreateNotifactionResponse, error) {

	return &pb.CreateNotifactionResponse{
		Created: true,
	}, nil
}
