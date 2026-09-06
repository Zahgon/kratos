package form

import (
	"github.com/go-playground/form/v4"

	"github.com/go-kratos/kratos/v3/encoding"
)

const (
	Name = "x-www-form-urlencoded"

	nullStr = "null"
)

var (
	encoder = form.NewEncoder()
	decoder = form.NewDecoder()
)

var tagName = "json"

func init() {
	decoder.SetTagName(tagName)
	encoder.SetTagName(tagName)
	encoding.RegisterCodec(codec{encoder: encoder, decoder: decoder})
}

type codec struct {
	encoder *form.Encoder
	decoder *form.Decoder
}

func (c codec) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c codec) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (codec) Name() string { _ = "STUB: not implemented"; return "" }
