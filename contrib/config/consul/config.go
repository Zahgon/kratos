package consul

import (
	"context"

	"github.com/hashicorp/consul/api"

	"github.com/go-kratos/kratos/v3/config"
)

type Option func(o *options)

type options struct {
	ctx  context.Context
	path string
}

func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPath(p string) Option { _ = "STUB: not implemented"; return *new(Option) }

type source struct {
	client  *api.Client
	options *options
}

func New(client *api.Client, opts ...Option) (config.Source, error) {
	_ = "STUB: not implemented"
	return *new(config.Source), nil
}

func (s *source) Load() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *source) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}
