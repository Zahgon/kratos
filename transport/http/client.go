package http

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v3/encoding"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/registry"
	"github.com/go-kratos/kratos/v3/selector"
	"github.com/go-kratos/kratos/v3/selector/wrr"
)

func init() {
	if selector.GlobalSelector() == nil {
		selector.SetGlobalSelector(wrr.NewBuilder())
	}
}

type DecodeErrorFunc func(ctx context.Context, res *http.Response) error

type EncodeRequestFunc func(ctx context.Context, contentType string, in any) (body []byte, err error)

type DecodeResponseFunc func(ctx context.Context, res *http.Response, out any) error

type ClientOption func(*clientOptions)

type clientOptions struct {
	ctx          context.Context
	tlsConf      *tls.Config
	timeout      time.Duration
	endpoint     string
	userAgent    string
	encoder      EncodeRequestFunc
	decoder      DecodeResponseFunc
	errorDecoder DecodeErrorFunc
	transport    http.RoundTripper
	nodeFilters  []selector.NodeFilter
	discovery    registry.Discovery
	middleware   []middleware.Middleware
	block        bool
	subsetSize   int
}

func WithSubset(size int) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

func WithTransport(trans http.RoundTripper) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithTimeout(d time.Duration) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithUserAgent(ua string) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

func WithMiddleware(m ...middleware.Middleware) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithEndpoint(endpoint string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithRequestEncoder(encoder EncodeRequestFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithResponseDecoder(decoder DecodeResponseFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithErrorDecoder(errorDecoder DecodeErrorFunc) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithDiscovery(d registry.Discovery) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithNodeFilter(filters ...selector.NodeFilter) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithBlock() ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

func WithTLSConfig(c *tls.Config) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

type Client struct {
	opts     clientOptions
	target   *Target
	r        *resolver
	cc       *http.Client
	insecure bool
	selector selector.Selector
}

func NewClient(ctx context.Context, opts ...ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client *Client) Invoke(ctx context.Context, method, path string, args any, reply any, opts ...CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (client *Client) invoke(ctx context.Context, req *http.Request, args any, reply any, c callInfo, opts ...CallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (client *Client) Do(req *http.Request, opts ...CallOption) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client *Client) do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client *Client) Close() error { _ = "STUB: not implemented"; return nil }

func DefaultRequestEncoder(_ context.Context, contentType string, in any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DefaultResponseDecoder(_ context.Context, res *http.Response, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func DefaultErrorDecoder(_ context.Context, res *http.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func CodecForResponse(r *http.Response) encoding.Codec {
	_ = "STUB: not implemented"
	return *new(encoding.Codec)
}
