package config

import (
	"github.com/polarismesh/polaris-go"

	"github.com/go-kratos/kratos/v3/config"
)

type Option func(o *options)

type options struct {
	namespace  string
	fileGroup  string
	fileName   string
	configFile polaris.ConfigFile
}

func WithNamespace(namespace string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFileGroup(fileGroup string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFileName(fileName string) Option { _ = "STUB: not implemented"; return *new(Option) }

type source struct {
	client  polaris.ConfigAPI
	options *options
}

func New(client polaris.ConfigAPI, opts ...Option) (config.Source, error) {
	_ = "STUB: not implemented"
	return *new(config.Source), nil
}

func (s *source) Load() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *source) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}
