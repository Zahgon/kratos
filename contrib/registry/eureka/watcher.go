package eureka

import (
	"context"

	"github.com/go-kratos/kratos/v3/registry"
)

var _ registry.Watcher = (*watcher)(nil)

type watcher struct {
	ctx        context.Context
	cancel     context.CancelFunc
	cli        *API
	watchChan  chan struct{}
	serverName string
}

func newWatch(ctx context.Context, cli *API, serverName string) (*watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Next() (services []*registry.ServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
