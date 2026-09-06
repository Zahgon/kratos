package grpc

import (
	"context"
	"crypto/tls"
	"time"

	"google.golang.org/grpc"

	"github.com/go-kratos/kratos/v3/internal/matcher"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/registry"
	"github.com/go-kratos/kratos/v3/selector"
	"github.com/go-kratos/kratos/v3/selector/wrr"

	_ "github.com/go-kratos/kratos/v3/transport/grpc/resolver/direct"
)

func init() {
	if selector.GlobalSelector() == nil {
		selector.SetGlobalSelector(wrr.NewBuilder())
	}
}

type ClientOption func(o *clientOptions)

func WithEndpoint(endpoint string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithSubset(size int) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

func WithTimeout(timeout time.Duration) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithMiddleware(m ...middleware.Middleware) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithStreamMiddleware(m ...middleware.Middleware) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithDiscovery(d registry.Discovery) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithTLSConfig(c *tls.Config) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithUnaryInterceptor(in ...grpc.UnaryClientInterceptor) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithStreamInterceptor(in ...grpc.StreamClientInterceptor) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithOptions(opts ...grpc.DialOption) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithNodeFilter(filters ...selector.NodeFilter) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithHealthCheck(healthCheck bool) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

type clientOptions struct {
	endpoint          string
	subsetSize        int
	tlsConf           *tls.Config
	timeout           time.Duration
	discovery         registry.Discovery
	middleware        []middleware.Middleware
	streamMiddleware  []middleware.Middleware
	ints              []grpc.UnaryClientInterceptor
	streamInts        []grpc.StreamClientInterceptor
	grpcOpts          []grpc.DialOption
	balancerName      string
	filters           []selector.NodeFilter
	healthCheckConfig string
}

func NewClient(ctx context.Context, opts ...ClientOption) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unaryClientInterceptor(ms []middleware.Middleware, timeout time.Duration, filters []selector.NodeFilter) grpc.UnaryClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryClientInterceptor)
}

type wrappedClientStream struct {
	grpc.ClientStream
	ctx        context.Context
	middleware matcher.Matcher
}

func (w *wrappedClientStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (w *wrappedClientStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (w *wrappedClientStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func streamClientInterceptor(ms []middleware.Middleware, filters []selector.NodeFilter) grpc.StreamClientInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamClientInterceptor)
}
