package nacos

import (
	"context"
	"errors"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"

	"github.com/go-kratos/kratos/v3/registry"
)

var ErrServiceInstanceNameEmpty = errors.New("kratos/nacos: ServiceInstance.Name can not be empty")

const (
	defaultKind = "grpc"
	kindKey     = "kind"
	versionKey  = "version"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

type options struct {
	prefix  string
	weight  float64
	cluster string
	group   string
	kind    string
}

type Option func(o *options)

func WithPrefix(prefix string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWeight(weight float64) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCluster(cluster string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGroup(group string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDefaultKind(kind string) Option { _ = "STUB: not implemented"; return *new(Option) }

type Registry struct {
	opts options
	cli  naming_client.INamingClient
}

func New(cli naming_client.INamingClient, opts ...Option) (r *Registry) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Register(_ context.Context, si *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Deregister(_ context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

func (r *Registry) GetService(_ context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
