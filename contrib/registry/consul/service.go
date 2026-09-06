package consul

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/go-kratos/kratos/v3/registry"
)

type serviceSet struct {
	registry    *Registry
	serviceName string
	watcher     map[*watcher]struct{}
	ref         atomic.Int32
	services    *atomic.Value
	lock        sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc
}

func (s *serviceSet) broadcast(ss []*registry.ServiceInstance) { _ = "STUB: not implemented"; return }

func (s *serviceSet) delete(w *watcher) { _ = "STUB: not implemented"; return }
