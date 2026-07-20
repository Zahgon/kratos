package mcp

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-kratos/kratos/v3/transport"

	"github.com/mark3labs/mcp-go/server"
)

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
	_ http.Handler         = (*Server)(nil)
)

type MiddlewareFunc func(http.Handler) http.Handler

type ServerOption func(*Server)

func Address(addr string) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func Endpoint(endpoint *url.URL) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

func Middleware(m MiddlewareFunc) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func SrvOptions(opts ...server.ServerOption) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func SSEOptions(opts ...server.SSEOption) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

type Server struct {
	*server.MCPServer
	srv        *http.Server
	sse        *server.SSEServer
	middleware MiddlewareFunc
	address    string
	endpoint   *url.URL
	srvOpts    []server.ServerOption
	sseOpts    []server.SSEOption
}

func NewServer(name, version string, opts ...ServerOption) *Server {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) Endpoint() (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Server) Start(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
