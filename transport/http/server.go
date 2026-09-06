package http

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"

	"github.com/go-kratos/kratos/v3/internal/matcher"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/transport"
)

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
	_ http.Handler         = (*Server)(nil)
)

type ServerOption func(*Server)

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

func Filter(filters ...FilterFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func RequestVarsDecoder(dec DecodeRequestFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func RequestQueryDecoder(dec DecodeRequestFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func RequestDecoder(dec DecodeRequestFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ResponseEncoder(en EncodeResponseFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func ErrorEncoder(en EncodeErrorFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func TLSConfig(c *tls.Config) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func StrictSlash(strictSlash bool) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func Listener(lis net.Listener) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func PathPrefix(prefix string) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func NotFoundHandler(handler http.Handler) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func MethodNotAllowedHandler(handler http.Handler) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

type Server struct {
	*http.Server
	lis         net.Listener
	tlsConf     *tls.Config
	endpoint    *url.URL
	err         error
	network     string
	address     string
	timeout     time.Duration
	filters     []FilterFunc
	middleware  matcher.Matcher
	decVars     DecodeRequestFunc
	decQuery    DecodeRequestFunc
	decBody     DecodeRequestFunc
	enc         EncodeResponseFunc
	ene         EncodeErrorFunc
	strictSlash bool
	router      *mux.Router
}

func NewServer(opts ...ServerOption) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Use(selector string, m ...middleware.Middleware) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) WalkRoute(fn WalkRouteFunc) error { _ = "STUB: not implemented"; return nil }

func (s *Server) WalkHandle(handle func(method, path string, handler http.HandlerFunc)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) Route(prefix string, filters ...FilterFunc) *Router {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) Handle(path string, h http.Handler) { _ = "STUB: not implemented"; return }

func (s *Server) HandlePrefix(prefix string, h http.Handler) { _ = "STUB: not implemented"; return }

func (s *Server) HandleFunc(path string, h http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (s *Server) HandleHeader(key, val string, h http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) filter() mux.MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(mux.MiddlewareFunc)
}

func (s *Server) Endpoint() (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Server) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) listenAndEndpoint() error { _ = "STUB: not implemented"; return nil }
