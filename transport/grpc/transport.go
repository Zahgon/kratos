package grpc

import (
	"google.golang.org/grpc/metadata"

	"github.com/go-kratos/kratos/v3/selector"
	"github.com/go-kratos/kratos/v3/transport"
)

var _ transport.Transporter = (*Transport)(nil)

type Transport struct {
	endpoint    string
	operation   string
	reqHeader   headerCarrier
	replyHeader headerCarrier
	nodeFilters []selector.NodeFilter
}

func (tr *Transport) Kind() transport.Kind { _ = "STUB: not implemented"; return *new(transport.Kind) }

func (tr *Transport) Endpoint() string { _ = "STUB: not implemented"; return "" }

func (tr *Transport) Operation() string { _ = "STUB: not implemented"; return "" }

func (tr *Transport) RequestHeader() transport.Header {
	_ = "STUB: not implemented"
	return *new(transport.Header)
}

func (tr *Transport) ReplyHeader() transport.Header {
	_ = "STUB: not implemented"
	return *new(transport.Header)
}

func (tr *Transport) NodeFilters() []selector.NodeFilter { _ = "STUB: not implemented"; return nil }

type headerCarrier metadata.MD

func (mc headerCarrier) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (mc headerCarrier) Set(key string, value string) { _ = "STUB: not implemented"; return }

func (mc headerCarrier) Add(key string, value string) { _ = "STUB: not implemented"; return }

func (mc headerCarrier) Keys() []string { _ = "STUB: not implemented"; return nil }

func (mc headerCarrier) Values(key string) []string { _ = "STUB: not implemented"; return nil }
