package env

import (
	"github.com/go-kratos/kratos/v3/config"
)

type env struct {
	prefixes []string
}

func NewSource(prefixes ...string) config.Source {
	_ = "STUB: not implemented"
	return *new(config.Source)
}

func (e *env) Load() (kvs []*config.KeyValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *env) load(envs []string) []*config.KeyValue { _ = "STUB: not implemented"; return nil }

func (e *env) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}

func matchPrefix(prefixes []string, s string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
