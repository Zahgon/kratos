package discovery

import (
	"context"

	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"

	"github.com/go-kratos/kratos/v3/registry"
)

type discoveryResolver struct {
	w  registry.Watcher
	cc resolver.ClientConn

	ctx    context.Context
	cancel context.CancelFunc

	insecure    bool
	selectorKey string
	subsetSize  int
}

func (r *discoveryResolver) watch() { _ = "STUB: not implemented"; return }

func (r *discoveryResolver) update(ins []*registry.ServiceInstance) {
	_ = "STUB: not implemented"
	return
}

func (r *discoveryResolver) Close() { _ = "STUB: not implemented"; return }

func (r *discoveryResolver) ResolveNow(_ resolver.ResolveNowOptions) {
	_ = "STUB: not implemented"
	return
}

func parseAttributes(md map[string]string) (a *attributes.Attributes) {
	_ = "STUB: not implemented"
	return nil
}
