package polaris

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v3/registry"

	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/config"
	"github.com/polarismesh/polaris-go/pkg/model"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

const _instanceIDSeparator = "-"

type options struct {
	Namespace string

	ServiceToken string

	Protocol *string

	Weight int

	Priority int

	Healthy bool

	Heartbeat bool

	Isolate bool

	TTL int

	Timeout time.Duration

	RetryCount int
}

type Option func(o *options)

type Registry struct {
	opt      options
	provider api.ProviderAPI
	consumer api.ConsumerAPI
}

func WithNamespace(namespace string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithServiceToken(serviceToken string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProtocol(protocol string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWeight(weight int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHealthy(healthy bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithIsolate(isolate bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTTL(TTL int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRetryCount(retryCount int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeartbeat(heartbeat bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewRegistry(provider api.ProviderAPI, consumer api.ConsumerAPI, opts ...Option) (r *Registry) {
	_ = "STUB: not implemented"
	return nil
}

func NewRegistryWithConfig(conf config.Configuration, opts ...Option) (r *Registry) {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Register(_ context.Context, serviceInstance *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Deregister(_ context.Context, serviceInstance *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) GetService(_ context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

type Watcher struct {
	ServiceName      string
	Namespace        string
	Ctx              context.Context
	Cancel           context.CancelFunc
	Channel          <-chan *model.InstancesResponse
	ServiceInstances []*registry.ServiceInstance
	watchResponse    *model.WatchAllInstancesResponse
	first            bool
}

type instancesListener struct {
	events chan *model.InstancesResponse
}

func (l instancesListener) OnInstancesUpdate(resp *model.InstancesResponse) {
	_ = "STUB: not implemented"
	return
}

func newWatcher(ctx context.Context, namespace string, serviceName string, consumer api.ConsumerAPI) (*Watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Watcher) Next() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Watcher) Stop() error { _ = "STUB: not implemented"; return nil }

func instancesToServiceInstances(instances []model.Instance) []*registry.ServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

func instanceToServiceInstance(instance model.Instance) *registry.ServiceInstance {
	_ = "STUB: not implemented"
	return nil
}
