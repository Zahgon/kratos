package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/go-kratos/kratos/v3/config"
)

type Option func(o *options)

type options struct {
	ctx    context.Context
	path   string
	prefix bool
}

func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPath(p string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPrefix(prefix bool) Option { _ = "STUB: not implemented"; return *new(Option) }

type source struct {
	client  *clientv3.Client
	options *options
}

func New(client *clientv3.Client, opts ...Option) (config.Source, error) {
	_ = "STUB: not implemented"
	return *new(config.Source), nil
}

func (s *source) Load() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *source) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}
