package kratos

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v3/registry"
)

type AppInfo interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Endpoint() []string
}

type App struct {
	opts     options
	ctx      context.Context
	cancel   context.CancelFunc
	mu       sync.Mutex
	instance *registry.ServiceInstance
}

func New(opts ...Option) *App { _ = "STUB: not implemented"; return nil }

func (a *App) ID() string { _ = "STUB: not implemented"; return "" }

func (a *App) Name() string { _ = "STUB: not implemented"; return "" }

func (a *App) Version() string { _ = "STUB: not implemented"; return "" }

func (a *App) Metadata() map[string]string { _ = "STUB: not implemented"; return nil }

func (a *App) Endpoint() []string { _ = "STUB: not implemented"; return nil }

func (a *App) Run() error { _ = "STUB: not implemented"; return nil }

func (a *App) Stop() (err error) { _ = "STUB: not implemented"; return nil }

func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type appKey struct{}

func NewContext(ctx context.Context, s AppInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) (s AppInfo, ok bool) {
	_ = "STUB: not implemented"
	return *new(AppInfo), false
}
