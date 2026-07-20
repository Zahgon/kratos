package discovery

import (
	"context"

	"github.com/go-kratos/kratos/v3/registry"
)

func filterInstancesByZone(ins *disInstancesInfo, zone string) []*registry.ServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discovery) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Discovery) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

type watcher struct {
	resolve *Resolve

	cancelCtx   context.Context
	serviceName string
}

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
