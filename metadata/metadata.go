package metadata

import (
	"context"
)

type Metadata map[string][]string

func New(mds ...map[string][]string) Metadata { _ = "STUB: not implemented"; return *new(Metadata) }

func (m Metadata) Add(key, value string) { _ = "STUB: not implemented"; return }

func (m Metadata) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (m Metadata) Set(key string, value string) { _ = "STUB: not implemented"; return }

func (m Metadata) Range(f func(k string, v []string) bool) { _ = "STUB: not implemented"; return }

func (m Metadata) Values(key string) []string { _ = "STUB: not implemented"; return nil }

func (m Metadata) Clone() Metadata { _ = "STUB: not implemented"; return *new(Metadata) }

type serverMetadataKey struct{}

func NewServerContext(ctx context.Context, md Metadata) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromServerContext(ctx context.Context) (Metadata, bool) {
	_ = "STUB: not implemented"
	return *new(Metadata), false
}

type clientMetadataKey struct{}

func NewClientContext(ctx context.Context, md Metadata) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromClientContext(ctx context.Context) (Metadata, bool) {
	_ = "STUB: not implemented"
	return *new(Metadata), false
}

func AppendToClientContext(ctx context.Context, kv ...string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func MergeToClientContext(ctx context.Context, cmd Metadata) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
