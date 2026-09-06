package json

import (
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/go-kratos/kratos/v3/encoding"
)

const Name = "json"

var (
	MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

func init() {
	encoding.RegisterCodec(codec{})
}

type codec struct{}

func (codec) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (codec) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (codec) Name() string { _ = "STUB: not implemented"; return "" }
