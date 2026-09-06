package proto

import (
	"google.golang.org/protobuf/proto"

	"github.com/go-kratos/kratos/v3/encoding"
)

const Name = "proto"

func init() {
	encoding.RegisterCodec(codec{})
}

type codec struct{}

func (codec) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (codec) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (codec) Name() string { _ = "STUB: not implemented"; return "" }

func getProtoMessage(v any) (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}
