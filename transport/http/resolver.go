package http

import (
	"context"

	"github.com/go-kratos/kratos/v3/registry"
	"github.com/go-kratos/kratos/v3/selector"
)

type Target struct {
	Scheme    string
	Authority string
	Endpoint  string
}

func parseTarget(endpoint string, insecure bool) (*Target, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type resolver struct {
	rebalancer selector.Rebalancer

	target      *Target
	watcher     registry.Watcher
	selectorKey string
	subsetSize  int

	insecure bool
}

func newResolver(ctx context.Context, discovery registry.Discovery, target *Target,
	rebalancer selector.Rebalancer, block, insecure bool, subsetSize int,
) (*resolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *resolver) update(services []*registry.ServiceInstance) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *resolver) Close() error { _ = "STUB: not implemented"; return nil }
