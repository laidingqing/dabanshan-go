package service

import (
	"context"
	"fmt"
	"time"

	"github.com/laidingqing/dabanshan/pb"
)

// RpcServer is used to implement helloworld.GreeterServer.
type RpcServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *RpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
