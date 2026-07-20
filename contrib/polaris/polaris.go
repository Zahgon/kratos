package polaris

import (
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/api"

	"github.com/go-kratos/kratos/v3/config"
)

type Polaris struct {
	router    polaris.RouterAPI
	config    polaris.ConfigAPI
	limit     polaris.LimitAPI
	registry  polaris.ProviderAPI
	discovery polaris.ConsumerAPI
	namespace string
	service   string
}

type Option func(o *Polaris)

func WithNamespace(ns string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithService(service string) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(sdk api.SDKContext, opts ...Option) Polaris {
	_ = "STUB: not implemented"
	return *new(Polaris)
}

func (p *Polaris) Config(opts ...ConfigOption) (config.Source, error) {
	_ = "STUB: not implemented"
	return *new(config.Source), nil
}

func (p *Polaris) Registry(opts ...RegistryOption) (r *Registry) {
	_ = "STUB: not implemented"
	return nil
}

func (p *Polaris) Limiter(opts ...LimiterOption) (r Limiter) {
	_ = "STUB: not implemented"
	return *new(Limiter)
}
