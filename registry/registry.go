package registry

import (
	"context"
)

type Registrar interface {
	Register(ctx context.Context, service *ServiceInstance) error

	Deregister(ctx context.Context, service *ServiceInstance) error
}

type Discovery interface {
	GetService(ctx context.Context, serviceName string) ([]*ServiceInstance, error)

	Watch(ctx context.Context, serviceName string) (Watcher, error)
}

type Watcher interface {
	Next() ([]*ServiceInstance, error)

	Stop() error
}

type ServiceInstance struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Version string `json:"version"`

	Metadata map[string]string `json:"metadata"`

	Endpoints []string `json:"endpoints"`
}

func (i *ServiceInstance) String() string { _ = "STUB: not implemented"; return "" }

func (i *ServiceInstance) Equal(o any) bool { _ = "STUB: not implemented"; return false }
