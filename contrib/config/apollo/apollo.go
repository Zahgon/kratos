package apollo

import (
	"github.com/apolloconfig/agollo/v4"

	"github.com/go-kratos/kratos/v3/config"
)

type apollo struct {
	client agollo.Client
	opt    *options
}

const (
	yaml       = "yaml"
	yml        = "yml"
	json       = "json"
	properties = "properties"
)

var formats map[string]struct{}

type Option func(*options)

type options struct {
	appid          string
	secret         string
	cluster        string
	endpoint       string
	namespace      string
	isBackupConfig bool
	backupPath     string
	originConfig   bool
}

func WithAppID(appID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCluster(cluster string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEnableBackup() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDisableBackup() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSecret(secret string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithNamespace(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBackupPath(backupPath string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithOriginalConfig() Option { _ = "STUB: not implemented"; return *new(Option) }

func NewSource(opts ...Option) config.Source { _ = "STUB: not implemented"; return *new(config.Source) }

func format(ns string) string { _ = "STUB: not implemented"; return "" }

func (e *apollo) load() []*config.KeyValue { _ = "STUB: not implemented"; return nil }

func (e *apollo) getConfig(ns string) (*config.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e apollo) getOriginConfig(ns string) (*config.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *apollo) Load() (kv []*config.KeyValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *apollo) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}

func resolve(key string, value any, target map[string]any) { _ = "STUB: not implemented"; return }

func genKey(ns, sub string) string { _ = "STUB: not implemented"; return "" }

func init() {
	formats = make(map[string]struct{})

	formats[yaml] = struct{}{}
	formats[yml] = struct{}{}
	formats[json] = struct{}{}
	formats[properties] = struct{}{}
}
