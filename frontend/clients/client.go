package clients

import (
	"context"
	"time"

	"github.com/laidingqing/dabanshan/common/config"
	grpclb "github.com/laidingqing/dabanshan/common/registry"
	"github.com/laidingqing/dabanshan/pb"
	"google.golang.org/grpc"
)

var accountCli pb.AccountServiceClient
var accountServiceName = "account_service"

//GetAccountClient get grpc client
func GetAccountClient() pb.AccountServiceClient {
	if accountCli == nil {
		r := grpclb.NewResolver(accountServiceName)
		b := grpc.RoundRobin(r)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		conn, err := grpc.DialContext(ctx, config.Service.RegistryLocation, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
		if err != nil {
			panic(err)
		}
		accountCli = pb.NewAccountServiceClient(conn)
	}
	return accountCli
}
