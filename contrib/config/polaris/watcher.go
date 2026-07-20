package config

import (
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"

	"github.com/go-kratos/kratos/v3/config"
)

type Watcher struct {
	configFile polaris.ConfigFile
	fullPath   string
}

type eventChan struct {
	closed bool
	event  chan model.ConfigFileChangeEvent
}

var eventChanMap = make(map[string]eventChan)

func getFullPath(namespace string, fileGroup string, fileName string) string {
	_ = "STUB: not implemented"
	return ""
}

func receive(event model.ConfigFileChangeEvent) { _ = "STUB: not implemented"; return }

func newWatcher(configFile polaris.ConfigFile) *Watcher { _ = "STUB: not implemented"; return nil }

func (w *Watcher) Next() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *Watcher) Stop() error { _ = "STUB: not implemented"; return nil }
