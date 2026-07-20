package apollo

import (
	"context"

	"github.com/apolloconfig/agollo/v4/storage"

	"github.com/go-kratos/kratos/v3/config"
)

type watcher struct {
	out <-chan []*config.KeyValue

	ctx      context.Context
	cancelFn func()
}

type customChangeListener struct {
	in     chan<- []*config.KeyValue
	apollo *apollo
}

func (c *customChangeListener) onChange(namespace string, changes map[string]*storage.ConfigChange) []*config.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func (c *customChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	_ = "STUB: not implemented"
	return
}

func (c *customChangeListener) OnNewestChange(_ *storage.FullChangeEvent) {
	_ = "STUB: not implemented"
	return
}

func newWatcher(a *apollo) (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}

func (w *watcher) Next() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
