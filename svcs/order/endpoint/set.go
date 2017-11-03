package endpoint

import (
	"context"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	rl "github.com/juju/ratelimit"
	m_order "github.com/laidingqing/dabanshan/svcs/order/model"
	"github.com/laidingqing/dabanshan/svcs/order/service"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
)

// Set collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	CreateOrderEndpoint endpoint.Endpoint
	GetOrdersEndpoint   endpoint.Endpoint
	AddCartEndpoint     endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.Service, logger log.Logger, duration metrics.Histogram, trace stdopentracing.Tracer) Set {
	var (
		createOrderEndpoint endpoint.Endpoint
		getOrdersEndpoint   endpoint.Endpoint
		addCartEndpoint     endpoint.Endpoint
	)
	{
		createOrderEndpoint = MakeCreateOrderEndpoint(svc)
		createOrderEndpoint = ratelimit.NewTokenBucketLimiter(rl.NewBucketWithRate(1, 1))(createOrderEndpoint)
		createOrderEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(createOrderEndpoint)
		createOrderEndpoint = opentracing.TraceServer(trace, "CreateOrder")(createOrderEndpoint)
		createOrderEndpoint = LoggingMiddleware(log.With(logger, "method", "CreateOrder"))(createOrderEndpoint)
		createOrderEndpoint = InstrumentingMiddleware(duration.With("method", "CreateOrder"))(createOrderEndpoint)
	}
	{
		getOrdersEndpoint = MakeGetOrderEndpoint(svc)
		getOrdersEndpoint = ratelimit.NewTokenBucketLimiter(rl.NewBucketWithRate(1, 1))(getOrdersEndpoint)
		getOrdersEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getOrdersEndpoint)
		getOrdersEndpoint = opentracing.TraceServer(trace, "GetOrders")(getOrdersEndpoint)
		getOrdersEndpoint = LoggingMiddleware(log.With(logger, "method", "GetOrders"))(getOrdersEndpoint)
		getOrdersEndpoint = InstrumentingMiddleware(duration.With("method", "GetOrders"))(getOrdersEndpoint)

	}
	{
		addCartEndpoint = MakeAddCartEndpoint(svc)
		addCartEndpoint = ratelimit.NewTokenBucketLimiter(rl.NewBucketWithRate(1, 1))(addCartEndpoint)
		addCartEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(addCartEndpoint)
		addCartEndpoint = opentracing.TraceServer(trace, "AddCart")(addCartEndpoint)
		addCartEndpoint = LoggingMiddleware(log.With(logger, "method", "AddCart"))(addCartEndpoint)
		addCartEndpoint = InstrumentingMiddleware(duration.With("method", "AddCart"))(addCartEndpoint)

	}

	return Set{
		CreateOrderEndpoint: createOrderEndpoint,
		GetOrdersEndpoint:   getOrdersEndpoint,
		AddCartEndpoint:     addCartEndpoint,
	}
}

// CreateOrder implements the service interface, so Set may be used as a service.
func (s Set) CreateOrder(ctx context.Context, a m_order.CreateOrderRequest) (m_order.CreatedOrderResponse, error) {
	resp, err := s.CreateOrderEndpoint(ctx, a)
	if err != nil {
		return m_order.CreatedOrderResponse{}, err
	}
	response := resp.(m_order.CreatedOrderResponse)
	return response, response.Err
}

// GetOrders implements the service interface, so Set may be used as a service.
func (s Set) GetOrders(ctx context.Context, a m_order.GetOrdersRequest) (m_order.GetOrdersResponse, error) {
	resp, err := s.GetOrdersEndpoint(ctx, a)
	if err != nil {
		return m_order.GetOrdersResponse{}, err
	}
	response := resp.(m_order.GetOrdersResponse)
	return response, response.Err
}

// AddCart implements the service interface, so Set may be used as a service.
func (s Set) AddCart(ctx context.Context, a m_order.CreateCartRequest) (m_order.CreatedCartResponse, error) {
	resp, err := s.AddCartEndpoint(ctx, a)
	if err != nil {
		return m_order.CreatedCartResponse{}, err
	}
	response := resp.(m_order.CreatedCartResponse)
	return response, response.Err
}

// MakeCreateOrderEndpoint constructs a CreateOrder endpoint wrapping the service.
func MakeCreateOrderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(m_order.CreateOrderRequest)
		v, err := s.CreateOrder(ctx, req)
		return v, err
	}
}

// MakeGetOrderEndpoint constructs a GetOrders endpoint wrapping the service.
func MakeGetOrderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(m_order.GetOrdersRequest)
		v, err := s.GetOrders(ctx, req)
		return v, err
	}
}

// MakeAddCartEndpoint constructs a GetOrders endpoint wrapping the service.
func MakeAddCartEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(m_order.CreateCartRequest)
		v, err := s.AddCart(ctx, req)
		return v, err
	}
}
