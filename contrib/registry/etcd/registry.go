package etcd

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/go-kratos/kratos/v3/registry"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

type Option func(o *options)

type options struct {
	ctx       context.Context
	namespace string
	ttl       time.Duration
	maxRetry  int
}

func Context(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

func Namespace(ns string) Option { _ = "STUB: not implemented"; return *new(Option) }

func RegisterTTL(ttl time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func MaxRetry(num int) Option { _ = "STUB: not implemented"; return *new(Option) }

type Registry struct {
	opts   *options
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease

	ctxMap map[string]*serviceCancel
}

type serviceCancel struct {
	service *registry.ServiceInstance
	cancel  context.CancelFunc
}

func New(client *clientv3.Client, opts ...Option) (r *Registry) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) registerKey(service *registry.ServiceInstance) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) serviceKey(name string) string { _ = "STUB: not implemented"; return "" }

func (r *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

func (r *Registry) registerWithKV(ctx context.Context, key string, value string) (clientv3.LeaseID, error) {
	_ = "STUB: not implemented"
	return *new(clientv3.LeaseID), nil
}

func (r *Registry) heartBeat(ctx context.Context, leaseID clientv3.LeaseID, key string, value string) {
	_ = "STUB: not implemented"
	return
}
