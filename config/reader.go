package config

import (
	"sync"
)

type Reader interface {
	Merge(...*KeyValue) error
	Value(string) (Value, bool)
	Source() ([]byte, error)
	Resolve() error
}

type reader struct {
	opts   options
	values map[string]any
	lock   sync.Mutex
}

func newReader(opts options) Reader { _ = "STUB: not implemented"; return *new(Reader) }

func (r *reader) Merge(kvs ...*KeyValue) error { _ = "STUB: not implemented"; return nil }

func (r *reader) Value(path string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func (r *reader) Source() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *reader) Resolve() error { _ = "STUB: not implemented"; return nil }

func (r *reader) cloneMap() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

func cloneMap(src map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertMap(src any) any { _ = "STUB: not implemented"; return *new(any) }

func readValue(values map[string]any, path string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func marshalJSON(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func unmarshalJSON(data []byte, v any) error { _ = "STUB: not implemented"; return nil }
