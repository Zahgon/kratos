package grpc

import (
	"context"
	"crypto/tls"
	"net"
	"net/url"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"

	"github.com/go-kratos/kratos/v3/internal/matcher"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/transport"
)

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
)

type ServerOption func(o *Server)

func Network(network string) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func Address(addr string) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func Endpoint(endpoint *url.URL) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func Timeout(timeout time.Duration) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func Middleware(m ...middleware.Middleware) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func StreamMiddleware(m ...middleware.Middleware) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func CustomHealth() ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func TLSConfig(c *tls.Config) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func Listener(lis net.Listener) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func UnaryInterceptor(in ...grpc.UnaryServerInterceptor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func StreamInterceptor(in ...grpc.StreamServerInterceptor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func DisableReflection() ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func Options(opts ...grpc.ServerOption) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

type Server struct {
	*grpc.Server
	baseCtx           context.Context
	tlsConf           *tls.Config
	lis               net.Listener
	err               error
	network           string
	address           string
	endpoint          *url.URL
	timeout           time.Duration
	middleware        matcher.Matcher
	streamMiddleware  matcher.Matcher
	unaryInts         []grpc.UnaryServerInterceptor
	streamInts        []grpc.StreamServerInterceptor
	grpcOpts          []grpc.ServerOption
	health            *health.Server
	customHealth      bool
	adminClean        func()
	disableReflection bool
}

func NewServer(opts ...ServerOption) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Use(selector string, m ...middleware.Middleware) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) Endpoint() (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Server) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) listenAndEndpoint() error { _ = "STUB: not implemented"; return nil }
