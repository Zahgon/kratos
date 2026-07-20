package eureka

import (
	"context"
	"sync"
	"time"
)

type subscriber struct {
	appID    string
	callBack func()
}

type API struct {
	cli             *Client
	allInstances    map[string][]Instance
	subscribers     map[string]*subscriber
	refreshInterval time.Duration
	lock            sync.Mutex
}

func NewAPI(ctx context.Context, client *Client, refreshInterval time.Duration) *API {
	_ = "STUB: not implemented"
	return nil
}

func (e *API) refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *API) broadcast() { _ = "STUB: not implemented"; return }

func (e *API) cacheAllInstances() map[string][]Instance { _ = "STUB: not implemented"; return nil }

func (e *API) Register(ctx context.Context, serviceName string, endpoints ...Endpoint) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *API) Deregister(ctx context.Context, endpoints []Endpoint) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *API) Subscribe(serverName string, fn func()) error { _ = "STUB: not implemented"; return nil }

func (e *API) GetService(ctx context.Context, serverName string) []Instance {
	_ = "STUB: not implemented"
	return nil
}

func (e *API) Unsubscribe(serverName string) { _ = "STUB: not implemented"; return }

func (e *API) ToAppID(serverName string) string { _ = "STUB: not implemented"; return "" }
