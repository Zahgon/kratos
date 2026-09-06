package kubernetes

import (
	"k8s.io/apimachinery/pkg/watch"

	"github.com/go-kratos/kratos/v3/config"
)

type watcher struct {
	k       *kube
	watcher watch.Interface
}

func newWatcher(k *kube) (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}

func (w *watcher) Next() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
