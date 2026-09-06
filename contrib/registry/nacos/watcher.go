package nacos

import (
	"context"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"github.com/go-kratos/kratos/v3/registry"
)

var _ registry.Watcher = (*watcher)(nil)

type watcher struct {
	serviceName    string
	clusters       []string
	groupName      string
	ctx            context.Context
	cancel         context.CancelFunc
	watchChan      chan struct{}
	cli            naming_client.INamingClient
	kind           string
	subscribeParam *vo.SubscribeParam
}

func newWatcher(ctx context.Context, cli naming_client.INamingClient, serviceName, groupName, kind string, clusters []string) (*watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *watcher) Stop() error { _ = "STUB: not implemented"; return nil }
