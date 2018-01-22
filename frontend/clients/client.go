package clients

import (
	"context"
	"time"

	"github.com/laidingqing/dabanshan/common/config"
	grpclb "github.com/laidingqing/dabanshan/common/registry"
	"github.com/laidingqing/dabanshan/pb"
	"google.golang.org/grpc"
)

var userCli pb.UserServiceClient
var userServiceName = "user_service"

//GetClient get grpc client
func GetClient() pb.UserServiceClient {
	if userCli == nil {
		r := grpclb.NewResolver(userServiceName)
		b := grpc.RoundRobin(r)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		conn, err := grpc.DialContext(ctx, config.Service.RegistryLocation, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
		if err != nil {
			panic(err)
		}
		userCli = pb.NewUserServiceClient(conn)
	}
	return userCli
}
