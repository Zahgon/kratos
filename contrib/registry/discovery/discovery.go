package discovery

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/go-resty/resty/v2"
)

type Discovery struct {
	config     *Config
	once       sync.Once
	ctx        context.Context
	cancelFunc context.CancelFunc
	httpClient *resty.Client

	node    atomic.Value
	nodeIdx atomic.Uint64

	mutex       sync.RWMutex
	apps        map[string]*appInfo
	registry    map[string]struct{}
	lastHost    string
	cancelPolls context.CancelFunc
}

type appInfo struct {
	resolver map[*Resolve]struct{}
	zoneIns  atomic.Value
	lastTs   int64
}

func New(c *Config) *Discovery { _ = "STUB: not implemented"; return nil }

func (d *Discovery) Close() error { _ = "STUB: not implemented"; return nil }

func (d *Discovery) selfProc(resolver *Resolve, event <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (d *Discovery) newSelf(zones map[string][]*discoveryInstance) {
	_ = "STUB: not implemented"
	return
}

func (d *Discovery) resolveBuild(appID string) *Resolve { _ = "STUB: not implemented"; return nil }

func (d *Discovery) serverProc() { _ = "STUB: not implemented"; return }

func (d *Discovery) pickNode() string { _ = "STUB: not implemented"; return "" }

func (d *Discovery) switchNode() { _ = "STUB: not implemented"; return }

func (d *Discovery) renew(ctx context.Context, ins *discoveryInstance) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discovery) cancel(ins *discoveryInstance) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discovery) broadcast(apps map[string]*disInstancesInfo) { _ = "STUB: not implemented"; return }

func (d *Discovery) polls(ctx context.Context) (apps map[string]*disInstancesInfo, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Resolve struct {
	id    string
	event chan struct{}
	d     *Discovery
}

func (r *Resolve) Watch() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (r *Resolve) fetch(_ context.Context) (ins *disInstancesInfo, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *Resolve) Close() error { _ = "STUB: not implemented"; return nil }
