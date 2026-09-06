package tracing

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/go-kratos/kratos/v3/middleware"
)

type Option func(*options)

type options struct {
	tracerName     string
	tracerProvider trace.TracerProvider
	propagator     propagation.TextMapPropagator
}

func WithPropagator(propagator propagation.TextMapPropagator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTracerProvider(provider trace.TracerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTracerName(tracerName string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Server(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func Client(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func TraceID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func SpanID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func TraceAttrs(ctx context.Context) []slog.Attr { _ = "STUB: not implemented"; return nil }
