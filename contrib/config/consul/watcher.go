package consul

import (
	"context"

	"github.com/hashicorp/consul/api/watch"

	"github.com/go-kratos/kratos/v3/config"
)

type watcher struct {
	source          *source
	ch              chan []*config.KeyValue
	wp              *watch.Plan
	fileModifyIndex map[string]uint64
	ctx             context.Context
	cancel          context.CancelFunc
}

func (w *watcher) handle(_ uint64, data any) { _ = "STUB: not implemented"; return }

func newWatcher(s *source) (*watcher, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *watcher) Next() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
