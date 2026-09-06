package grpc

import (
	"google.golang.org/grpc/encoding"
)

const jsonName = "json"

func init() {
	encoding.RegisterCodec(codec{})
}

type codec struct{}

func (codec) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (codec) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (codec) Name() string { _ = "STUB: not implemented"; return "" }
