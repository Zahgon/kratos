package kratos

import (
	"context"
	"log/slog"
	"net/url"
	"os"
	"time"

	"github.com/go-kratos/kratos/v3/registry"
	"github.com/go-kratos/kratos/v3/transport"
)

type Option func(o *options)

type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []*url.URL

	ctx  context.Context
	sigs []os.Signal

	logger           *slog.Logger
	registrar        registry.Registrar
	registrarTimeout time.Duration
	stopTimeout      time.Duration
	servers          []transport.Server

	beforeStart []func(context.Context) error
	beforeStop  []func(context.Context) error
	afterStart  []func(context.Context) error
	afterStop   []func(context.Context) error
}

func ID(id string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Name(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Version(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Metadata(md map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Endpoint(endpoints ...*url.URL) Option { _ = "STUB: not implemented"; return *new(Option) }

func Context(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

func Logger(logger *slog.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func Server(srv ...transport.Server) Option { _ = "STUB: not implemented"; return *new(Option) }

func Signal(sigs ...os.Signal) Option { _ = "STUB: not implemented"; return *new(Option) }

func Registrar(r registry.Registrar) Option { _ = "STUB: not implemented"; return *new(Option) }

func RegistrarTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func StopTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func BeforeStart(fn func(context.Context) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func BeforeStop(fn func(context.Context) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func AfterStart(fn func(context.Context) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func AfterStop(fn func(context.Context) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
