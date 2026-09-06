package consul

import (
	"context"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"

	"github.com/go-kratos/kratos/v3/registry"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

type Option func(*Registry)

func WithHealthCheck(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDatacenter(dc Datacenter) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeartbeat(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithServiceResolver(fn ServiceResolver) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHealthCheckInterval(interval int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDeregisterCriticalServiceAfter(interval int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithServiceCheck(checks ...*api.AgentServiceCheck) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTags(tags []string) Option { _ = "STUB: not implemented"; return *new(Option) }

type Config struct {
	*api.Config
}

type Registry struct {
	cli               *Client
	enableHealthCheck bool
	registry          map[string]*serviceSet
	lock              sync.RWMutex
	timeout           time.Duration
}

func New(apiClient *api.Client, opts ...Option) *Registry { _ = "STUB: not implemented"; return nil }

func (r *Registry) Register(ctx context.Context, svc *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Deregister(ctx context.Context, svc *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) ListServices() (allServices map[string][]*registry.ServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

func (r *Registry) resolve(ctx context.Context, ss *serviceSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) tryDelete(ss *serviceSet) bool { _ = "STUB: not implemented"; return false }
