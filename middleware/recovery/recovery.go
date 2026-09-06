package recovery

import (
	"context"
	"log/slog"

	"github.com/go-kratos/kratos/v3/errors"
	"github.com/go-kratos/kratos/v3/middleware"
)

type Latency struct{}

var ErrUnknownRequest = errors.InternalServer("UNKNOWN", "unknown request error")

type HandlerFunc func(ctx context.Context, req, err any) error

type Option func(*options)

type options struct {
	handler HandlerFunc
	logger  *slog.Logger
}

func WithHandler(h HandlerFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogger(logger *slog.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func Recovery(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

//nolint:mnd
