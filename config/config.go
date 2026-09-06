package config

import (
	"errors"
	"sync"

	_ "github.com/go-kratos/kratos/v3/encoding/json"
	_ "github.com/go-kratos/kratos/v3/encoding/proto"
	_ "github.com/go-kratos/kratos/v3/encoding/xml"
	_ "github.com/go-kratos/kratos/v3/encoding/yaml"
)

var _ Config = (*config)(nil)

var ErrNotFound = errors.New("key not found")

type Observer func(string, Value)

type Config interface {
	Load() error
	Scan(v any) error
	Value(key string) Value
	Watch(key string, o Observer) error
	Close() error
}

type config struct {
	opts      options
	reader    Reader
	cached    sync.Map
	observers sync.Map
	watchers  []Watcher
}

func New(opts ...Option) Config { _ = "STUB: not implemented"; return *new(Config) }

func (c *config) watch(w Watcher) { _ = "STUB: not implemented"; return }

func (c *config) Load() error { _ = "STUB: not implemented"; return nil }

func (c *config) Value(key string) Value { _ = "STUB: not implemented"; return *new(Value) }

func (c *config) Scan(v any) error { _ = "STUB: not implemented"; return nil }

func (c *config) Watch(key string, o Observer) error { _ = "STUB: not implemented"; return nil }

func (c *config) Close() error { _ = "STUB: not implemented"; return nil }

func Get[T any](c Config, key string) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }
