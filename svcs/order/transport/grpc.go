package transport

import (
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	jujuratelimit "github.com/juju/ratelimit"
	"github.com/laidingqing/dabanshan/pb"
	o_endpoint "github.com/laidingqing/dabanshan/svcs/order/endpoint"
	"github.com/laidingqing/dabanshan/svcs/order/service"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type grpcServer struct {
	createOrder grpctransport.Handler
	getOrders   grpctransport.Handler
	addCart     grpctransport.Handler
}

// NewGRPCServer ...
func NewGRPCServer(endpoints o_endpoint.Set, tracer stdopentracing.Tracer, logger log.Logger) pb.OrderRpcServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		createOrder: grpctransport.NewServer(
			endpoints.CreateOrderEndpoint,
			decodeGRPCCreateOrderRequest,
			encodeGRPCCreateOrderResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "CreateOrder", logger)))...,
		),
		getOrders: grpctransport.NewServer(
			endpoints.GetOrdersEndpoint,
			decodeGRPCGetOrdersRequest,
			encodeGRPCGetOrdersResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "GetOrders", logger)))...,
		),
		addCart: grpctransport.NewServer(
			endpoints.AddCartEndpoint,
			decodeGRPCAddCartRequest,
			encodeGRPCAddCartResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddCart", logger)))...,
		),
	}
}

// GetUser RPC
func (s *grpcServer) CreateOrder(ctx oldcontext.Context, req *pb.CreateOrderRequest) (*pb.CreatedOrderResponse, error) {
	_, rep, err := s.createOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rep.(*pb.CreatedOrderResponse)
	return res, nil
}

// GetOrders

func (s *grpcServer) GetOrders(ctx oldcontext.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	_, rep, err := s.getOrders.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rep.(*pb.GetOrdersResponse)
	return res, nil
}

// AddCart
func (s *grpcServer) AddCart(ctx oldcontext.Context, req *pb.CreateCartRequest) (*pb.CreatedCartResponse, error) {
	_, rep, err := s.addCart.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rep.(*pb.CreatedCartResponse)
	return res, nil
}

// NewGRPCClient ...
func NewGRPCClient(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) service.Service {
	limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(100, 100))
	var createOrderEndpoint endpoint.Endpoint
	var getOrdersEndpoint endpoint.Endpoint
	var addCartEndpoint endpoint.Endpoint
	{
		createOrderEndpoint = grpctransport.NewClient(
			conn,
			"pb.OrderRpcService",
			"CreateOrder",
			encodeGRPCCreateOrderRequest,
			decodeGRPCCreateOrderResponse,
			pb.CreatedOrderResponse{},
			grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
		).Endpoint()
		createOrderEndpoint = opentracing.TraceClient(tracer, "CreateOrder")(createOrderEndpoint)
		createOrderEndpoint = limiter(createOrderEndpoint)
		createOrderEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "CreateOrder",
			Timeout: 30 * time.Second,
		}))(createOrderEndpoint)

		getOrdersEndpoint = grpctransport.NewClient(
			conn,
			"pb.OrderRpcService",
			"GetOrders",
			encodeGRPCGetOrdersRequest,
			decodeGRPCGetOrdersResponse,
			pb.GetOrdersResponse{},
			grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
		).Endpoint()
		getOrdersEndpoint = opentracing.TraceClient(tracer, "GetOrders")(getOrdersEndpoint)
		getOrdersEndpoint = limiter(getOrdersEndpoint)
		getOrdersEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "GetOrders",
			Timeout: 30 * time.Second,
		}))(getOrdersEndpoint)

		addCartEndpoint = grpctransport.NewClient(
			conn,
			"pb.OrderRpcService",
			"AddCart",
			encodeGRPCAddCartRequest,
			decodeGRPCAddCartResponse,
			pb.CreatedCartResponse{},
			grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
		).Endpoint()
		addCartEndpoint = opentracing.TraceClient(tracer, "AddCart")(addCartEndpoint)
		addCartEndpoint = limiter(addCartEndpoint)
		addCartEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "AddCart",
			Timeout: 30 * time.Second,
		}))(addCartEndpoint)
	}
	return o_endpoint.Set{
		CreateOrderEndpoint: createOrderEndpoint,
		GetOrdersEndpoint:   getOrdersEndpoint,
	}
}
