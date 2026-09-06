package env

import (
	"context"

	"github.com/go-kratos/kratos/v3/config"
)

var _ config.Watcher = (*watcher)(nil)

type watcher struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewWatcher() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}

func (w *watcher) Next() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
