package discovery

import (
	"context"

	"github.com/go-kratos/kratos/v3/registry"
)

func (d *Discovery) Register(ctx context.Context, service *registry.ServiceInstance) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discovery) register(ctx context.Context, ins *discoveryInstance) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discovery) Deregister(_ context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}
