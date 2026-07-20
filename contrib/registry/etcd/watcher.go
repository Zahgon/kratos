package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/go-kratos/kratos/v3/registry"
)

var _ registry.Watcher = (*watcher)(nil)

type watcher struct {
	key         string
	ctx         context.Context
	cancel      context.CancelFunc
	client      *clientv3.Client
	watchChan   clientv3.WatchChan
	watcher     clientv3.Watcher
	kv          clientv3.KV
	first       bool
	serviceName string
}

func newWatcher(ctx context.Context, key, name string, client *clientv3.Client) (*watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }

func (w *watcher) getInstance() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) reWatch() error { _ = "STUB: not implemented"; return nil }
