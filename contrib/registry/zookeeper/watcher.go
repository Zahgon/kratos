package zookeeper

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/go-zookeeper/zk"

	"github.com/go-kratos/kratos/v3/registry"
)

var _ registry.Watcher = (*watcher)(nil)

var ErrWatcherStopped = errors.New("watcher stopped")

type watcher struct {
	ctx    context.Context
	event  chan zk.Event
	conn   *zk.Conn
	cancel context.CancelFunc

	first atomic.Bool

	prefix string

	serviceName string
}

func newWatcher(ctx context.Context, prefix, serviceName string, conn *zk.Conn) (*watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) watch(ctx context.Context) { _ = "STUB: not implemented"; return }

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }

func (w *watcher) getServices() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
