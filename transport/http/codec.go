package http

import (
	"net/http"
	"reflect"

	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/proto"

	"github.com/go-kratos/kratos/v3/encoding"
)

const SupportPackageIsVersion3 = true

const defaultHTTPBodyContentType = "application/octet-stream"

var protoMessageType = reflect.TypeOf((*proto.Message)(nil)).Elem()

type Redirector interface {
	error
	Redirect() (string, int)
}

type Request = http.Request

type ResponseWriter = http.ResponseWriter

type Flusher = http.Flusher

type DecodeRequestFunc func(*http.Request, any) error

type EncodeResponseFunc func(http.ResponseWriter, *http.Request, any) error

type EncodeErrorFunc func(http.ResponseWriter, *http.Request, error)

func DefaultRequestVars(r *http.Request, v any) error { _ = "STUB: not implemented"; return nil }

func DefaultRequestQuery(r *http.Request, v any) error { _ = "STUB: not implemented"; return nil }

func DefaultRequestDecoder(r *http.Request, v any) error { _ = "STUB: not implemented"; return nil }

func DefaultResponseEncoder(w http.ResponseWriter, r *http.Request, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func DefaultErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	_ = "STUB: not implemented"
	return
}

func CodecForRequest(r *http.Request, name string) (encoding.Codec, bool) {
	_ = "STUB: not implemented"
	return *new(encoding.Codec), false
}

func httpBody(v any) (*httpbody.HttpBody, bool) { _ = "STUB: not implemented"; return nil, false }

func decodeWithCodec(codec encoding.Codec, data []byte, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func BodyContentType(v any) string { _ = "STUB: not implemented"; return "" }
