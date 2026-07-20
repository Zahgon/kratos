package config

import (
	"context"

	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/go-kratos/kratos/v3/config"
)

type Watcher struct {
	dataID             string
	group              string
	content            chan string
	cancelListenConfig cancelListenConfigFunc

	ctx    context.Context
	cancel context.CancelFunc
}

type cancelListenConfigFunc func(params vo.ConfigParam) (err error)

func newWatcher(ctx context.Context, dataID string, group string, cancelListenConfig cancelListenConfigFunc) *Watcher {
	_ = "STUB: not implemented"
	return nil
}

func (w *Watcher) Next() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Watcher) Close() error { _ = "STUB: not implemented"; return nil }

func (w *Watcher) Stop() error { _ = "STUB: not implemented"; return nil }
