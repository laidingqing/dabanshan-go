package clients

import (
	"context"
	"time"

	"github.com/laidingqing/dabanshan/common/config"
	grpclb "github.com/laidingqing/dabanshan/common/registry"
	"github.com/laidingqing/dabanshan/pb"
	"google.golang.org/grpc"
)

//GetAccountClient get grpc client
func GetAccountClient() pb.AccountServiceClient {
	r := grpclb.NewResolver(config.ServiceAccountName)
	b := grpc.RoundRobin(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, config.Service.RegistryLocation, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	accountCli := pb.NewAccountServiceClient(conn)

	return accountCli
}

//GetEpisodeClient get grpc client
func GetEpisodeClient() pb.EpisodeServiceClient {
	r := grpclb.NewResolver(config.ServiceEpisodeName)
	b := grpc.RoundRobin(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, config.Service.RegistryLocation, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	episodeCli := pb.NewEpisodeServiceClient(conn)
	return episodeCli
}

//GetProductsClient get grpc client
func GetProductsClient() pb.ProductsServiceClient {
	r := grpclb.NewResolver(config.ServiceProductsName)
	b := grpc.RoundRobin(r)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, config.Service.RegistryLocation, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	productsCli := pb.NewProductsServiceClient(conn)
	return productsCli
}
