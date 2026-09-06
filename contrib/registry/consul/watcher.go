package consul

import (
	"context"

	"github.com/go-kratos/kratos/v3/registry"
)

type watcher struct {
	event chan struct{}
	set   *serviceSet

	ctx    context.Context
	cancel context.CancelFunc
}

func (w *watcher) Next() (services []*registry.ServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
