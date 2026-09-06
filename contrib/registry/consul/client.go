package consul

import (
	"context"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v3/registry"

	"github.com/hashicorp/consul/api"
)

type Datacenter string

const (
	SingleDatacenter Datacenter = "SINGLE"
	MultiDatacenter  Datacenter = "MULTI"
)

type Client struct {
	dc  Datacenter
	cli *api.Client

	resolver ServiceResolver

	healthcheckInterval int

	heartbeat bool

	deregisterCriticalServiceAfter int

	serviceChecks api.AgentServiceChecks

	tags []string

	lock      sync.RWMutex
	cancelers map[string]*canceler
}

type canceler struct {
	ctx    context.Context
	cancel context.CancelFunc
	done   chan struct{}
}

func defaultResolver(_ context.Context, entries []*api.ServiceEntry) []*registry.ServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

type ServiceResolver func(ctx context.Context, entries []*api.ServiceEntry) []*registry.ServiceInstance

func (c *Client) Service(ctx context.Context, service string, index uint64, passingOnly bool) ([]*registry.ServiceInstance, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (c *Client) multiDCService(ctx context.Context, service string, index uint64, passingOnly bool) ([]*registry.ServiceInstance, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (c *Client) singleDCEntries(service, tag string, passingOnly bool, opts *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) Register(ctx context.Context, svc *registry.ServiceInstance, enableHealthCheck bool) error {
	_ = "STUB: not implemented"
	return nil
}

func sleepCtx(ctx context.Context, d time.Duration) error { _ = "STUB: not implemented"; return nil }

func (c *Client) Deregister(ctx context.Context, serviceID string) error {
	_ = "STUB: not implemented"
	return nil
}
