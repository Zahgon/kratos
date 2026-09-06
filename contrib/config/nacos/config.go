package config

import (
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"

	"github.com/go-kratos/kratos/v3/config"
)

type Option func(*options)

type options struct {
	group  string
	dataID string
}

func WithGroup(group string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDataID(dataID string) Option { _ = "STUB: not implemented"; return *new(Option) }

type Config struct {
	opts   options
	client config_client.IConfigClient
}

func NewConfigSource(client config_client.IConfigClient, opts ...Option) config.Source {
	_ = "STUB: not implemented"
	return *new(config.Source)
}

func (c *Config) Load() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Config) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}
