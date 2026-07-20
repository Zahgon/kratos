package http

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v3/transport"
)

var _ Transporter = (*Transport)(nil)

var _ ResponseTransporter = (*Transport)(nil)

type Transporter interface {
	transport.Transporter
	Request() *http.Request
	PathTemplate() string
}

type ResponseTransporter interface {
	Transporter
	Response() http.ResponseWriter
}

type Transport struct {
	endpoint     string
	operation    string
	reqHeader    headerCarrier
	replyHeader  headerCarrier
	request      *http.Request
	response     http.ResponseWriter
	pathTemplate string
}

func (tr *Transport) Kind() transport.Kind { _ = "STUB: not implemented"; return *new(transport.Kind) }

func (tr *Transport) Endpoint() string { _ = "STUB: not implemented"; return "" }

func (tr *Transport) Operation() string { _ = "STUB: not implemented"; return "" }

func (tr *Transport) Request() *http.Request { _ = "STUB: not implemented"; return nil }

func (tr *Transport) RequestHeader() transport.Header {
	_ = "STUB: not implemented"
	return *new(transport.Header)
}

func (tr *Transport) Response() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

func (tr *Transport) ReplyHeader() transport.Header {
	_ = "STUB: not implemented"
	return *new(transport.Header)
}

func (tr *Transport) PathTemplate() string { _ = "STUB: not implemented"; return "" }

func SetOperation(ctx context.Context, op string) { _ = "STUB: not implemented"; return }

func SetCookie(ctx context.Context, cookie *http.Cookie) { _ = "STUB: not implemented"; return }

func RequestFromServerContext(ctx context.Context) (*http.Request, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type headerCarrier http.Header

func (hc headerCarrier) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (hc headerCarrier) Set(key string, value string) { _ = "STUB: not implemented"; return }

func (hc headerCarrier) Add(key string, value string) { _ = "STUB: not implemented"; return }

func (hc headerCarrier) Keys() []string { _ = "STUB: not implemented"; return nil }

func (hc headerCarrier) Values(key string) []string { _ = "STUB: not implemented"; return nil }

func ResponseWriterFromServerContext(ctx context.Context) (http.ResponseWriter, bool) {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter), false
}
