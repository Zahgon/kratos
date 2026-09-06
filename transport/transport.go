package transport

import (
	"context"
	"net/url"

	_ "github.com/go-kratos/kratos/v3/encoding/form"
	_ "github.com/go-kratos/kratos/v3/encoding/json"
	_ "github.com/go-kratos/kratos/v3/encoding/proto"
	_ "github.com/go-kratos/kratos/v3/encoding/protojson"
	_ "github.com/go-kratos/kratos/v3/encoding/xml"
	_ "github.com/go-kratos/kratos/v3/encoding/yaml"
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type Endpointer interface {
	Endpoint() (*url.URL, error)
}

type Header interface {
	Get(key string) string
	Set(key string, value string)
	Add(key string, value string)
	Keys() []string
	Values(key string) []string
}

type Transporter interface {
	Kind() Kind

	Endpoint() string

	Operation() string

	RequestHeader() Header

	ReplyHeader() Header
}

type Kind string

func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

const (
	KindGRPC Kind = "grpc"
	KindHTTP Kind = "http"
)

type (
	serverTransportKey struct{}
	clientTransportKey struct{}
)

func NewServerContext(ctx context.Context, tr Transporter) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromServerContext(ctx context.Context) (tr Transporter, ok bool) {
	_ = "STUB: not implemented"
	return *new(Transporter), false
}

func NewClientContext(ctx context.Context, tr Transporter) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromClientContext(ctx context.Context) (tr Transporter, ok bool) {
	_ = "STUB: not implemented"
	return *new(Transporter), false
}
