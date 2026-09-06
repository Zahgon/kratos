package servicecomb

import (
	"context"
	"os"

	"github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/sc-client"

	"github.com/go-kratos/kratos/v3/registry"
)

func init() {
	appID = os.Getenv(appIDVar)
	if appID == "" {
		appID = "default"
	}
	env = os.Getenv(envVar)
}

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

var (
	curServiceID string
	appID        string
	env          string
)

const (
	appIDKey         = "appId"
	envKey           = "environment"
	envVar           = "CAS_ENVIRONMENT_ID"
	appIDVar         = "CAS_APPLICATION_NAME"
	frameWorkName    = "kratos"
	frameWorkVersion = "v2"
)

type RegistryClient interface {
	GetMicroServiceID(appID, microServiceName, version, env string, opts ...sc.CallOption) (string, error)
	FindMicroServiceInstances(consumerID, appID, microServiceName, versionRule string, opts ...sc.CallOption) ([]*discovery.MicroServiceInstance, error)
	RegisterService(microService *discovery.MicroService) (string, error)
	RegisterMicroServiceInstance(microServiceInstance *discovery.MicroServiceInstance) (string, error)
	Heartbeat(microServiceID, microServiceInstanceID string) (bool, error)
	UnregisterMicroServiceInstance(microServiceID, microServiceInstanceID string) (bool, error)
	WatchMicroService(microServiceID string, callback func(*sc.MicroServiceInstanceChangedEvent)) error
}

type Registry struct {
	cli RegistryClient
}

func NewRegistry(client RegistryClient) *Registry { _ = "STUB: not implemented"; return nil }

func (r *Registry) GetService(_ context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

func (r *Registry) Register(_ context.Context, svcIns *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Deregister(_ context.Context, svcIns *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}
