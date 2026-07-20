package polaris

import (
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"

	"github.com/go-kratos/kratos/v3/config"
)

type ConfigOption func(o *configOptions)

type configOptions struct {
	namespace  string
	files      []File
	configFile []polaris.ConfigFile
}

func WithConfigFile(file ...File) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

type File struct {
	Name  string
	Group string
}

type source struct {
	client  polaris.ConfigAPI
	options *configOptions
}

func (s *source) Load() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *source) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}

type ConfigWatcher struct {
	event chan model.ConfigFileChangeEvent
	cfg   []*config.KeyValue
}

func receive(event chan model.ConfigFileChangeEvent) func(m model.ConfigFileChangeEvent) {
	_ = "STUB: not implemented"
	return nil
}

func newConfigWatcher(configFile []polaris.ConfigFile) *ConfigWatcher {
	_ = "STUB: not implemented"
	return nil
}

func (w *ConfigWatcher) Next() ([]*config.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *ConfigWatcher) Stop() error { _ = "STUB: not implemented"; return nil }
