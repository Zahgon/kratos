package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/go-kratos/kratos/v3/config"
)

type watcher struct {
	source *source
	ch     clientv3.WatchChan

	ctx    context.Context
	cancel context.CancelFunc
}

func newWatcher(s *source) *watcher { _ = "STUB: not implemented"; return nil }

func (w *watcher) Next() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
