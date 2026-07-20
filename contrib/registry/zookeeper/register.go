package zookeeper

import (
	"context"

	"github.com/go-zookeeper/zk"
	"golang.org/x/sync/singleflight"

	"github.com/go-kratos/kratos/v3/registry"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

type Option func(o *options)

type options struct {
	namespace string
	user      string
	password  string
}

func WithRootPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDigestACL(user string, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type Registry struct {
	opts *options
	conn *zk.Conn

	group singleflight.Group
}

func New(conn *zk.Conn, opts ...Option) *Registry { _ = "STUB: not implemented"; return nil }

func (r *Registry) Register(_ context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) GetService(_ context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

func (r *Registry) ensureName(path string, data []byte, flags int32) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) reRegister(path string, data []byte) { _ = "STUB: not implemented"; return }
