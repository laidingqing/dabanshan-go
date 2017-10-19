package transport

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	jujuratelimit "github.com/juju/ratelimit"
	"github.com/laidingqing/dabanshan/svcs/product/endpoint"
	"github.com/laidingqing/dabanshan/svcs/product/service"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	oldcontext "golang.org/x/net/context"

	"github.com/laidingqing/dabanshan/pb"
)

type grpcServer struct {
	getproducts grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Set, tracer stdopentracing.Tracer, logger log.Logger) pb.ProductRpcServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		getproducts: grpctransport.NewServer(
			endpoints.GetProductsEndpoint,
			encodeGRPCGetProductsRequest,
			decodeGRPCGetProductsResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "GetProducts", logger)))...,
		),
	}
}

func (s *grpcServer) GetProducts(ctx oldcontext.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	_, rep, err := s.getproducts.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetProductsResponse), nil
}

func NewGRPCClient(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) service.Service {
	limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(100, 100))
	var getProductsEndpoint endpoint.Endpoint
	{
		getProductsEndpoint = grpctransport.NewClient(
			conn,
			"pb.Add",
			"Sum",
			encodeGRPCSumRequest,
			decodeGRPCSumResponse,
			pb.GetProductsResponse{},
			grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
		).Endpoint()
		getProductsEndpoint = opentracing.TraceClient(tracer, "GetProducts")(getProductsEndpoint)
		getProductsEndpoint = limiter(getProductsEndpoint)
		getProductsEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Sum",
			Timeout: 30 * time.Second,
		}))(getProductsEndpoint)
	}
}

func encodeGRPCGetProductsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetProductsRequest)
	return &pb.GetProductsRequest{Creatorid: req.A, Size: req.B}, nil
}

func decodeGRPCGetProductsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetProductsResponse)
	return endpoint.GetProductsResponse{V: reply.GetProducts(), Err: nil}, nil
}
