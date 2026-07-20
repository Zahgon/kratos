package eureka

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v3/registry"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

type Option func(o *Registry)

func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeartbeat(interval time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRefresh(interval time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEurekaPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

type Registry struct {
	ctx               context.Context
	api               *API
	heartbeatInterval time.Duration
	refreshInterval   time.Duration
	eurekaPath        string
}

func New(eurekaUrls []string, opts ...Option) (*Registry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

func (r *Registry) Endpoints(service *registry.ServiceInstance) []Endpoint {
	_ = "STUB: not implemented"
	return nil
}
