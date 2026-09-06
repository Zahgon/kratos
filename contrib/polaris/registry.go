package polaris

import (
	"context"
	"time"

	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"

	"github.com/go-kratos/kratos/v3/registry"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

type registryOptions struct {
	Namespace string

	ServiceToken string

	Weight int

	Priority int

	Healthy bool

	Isolate bool

	TTL int

	Timeout time.Duration

	RetryCount int
}

type RegistryOption func(o *registryOptions)

type Registry struct {
	opt      registryOptions
	provider polaris.ProviderAPI
	consumer polaris.ConsumerAPI
}

func WithRegistryServiceToken(serviceToken string) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func WithRegistryWeight(weight int) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func WithRegistryHealthy(healthy bool) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func WithRegistryIsolate(isolate bool) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func WithRegistryTTL(TTL int) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func WithRegistryTimeout(timeout time.Duration) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func WithRegistryRetryCount(retryCount int) RegistryOption {
	_ = "STUB: not implemented"
	return *new(RegistryOption)
}

func (r *Registry) Register(_ context.Context, instance *registry.ServiceInstance) error {
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

func merge(instances []model.Instance) map[string][]model.Instance {
	_ = "STUB: not implemented"
	return nil
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
	Channel          <-chan model.SubScribeEvent
	service          *model.InstancesResponse
	ServiceInstances map[string][]model.Instance
	first            bool
}

func newWatcher(ctx context.Context, namespace string, serviceName string, consumer polaris.ConsumerAPI) (*Watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Watcher) Next() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Watcher) Stop() error { _ = "STUB: not implemented"; return nil }

func instancesToServiceInstances(instances map[string][]model.Instance) []*registry.ServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

func mapClone[M ~map[K]V, K comparable, V any](m M) M { _ = "STUB: not implemented"; return *new(M) }
