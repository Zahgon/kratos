package servicecomb

import (
	"context"

	"github.com/go-kratos/kratos/v3/registry"
)

var _ registry.Watcher = (*Watcher)(nil)

type Watcher struct {
	cli RegistryClient
	ch  chan *registry.ServiceInstance
}

func newWatcher(_ context.Context, cli RegistryClient, serviceName string) (*Watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Watcher) Put(svcIns *registry.ServiceInstance) { _ = "STUB: not implemented"; return }

func (w *Watcher) Next() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Watcher) Stop() error { _ = "STUB: not implemented"; return nil }
