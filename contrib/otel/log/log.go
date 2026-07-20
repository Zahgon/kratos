package log

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	otellog "go.opentelemetry.io/otel/log"
)

type Option func(*options)

type options struct {
	otel []otelslog.Option
}

func WithLoggerProvider(provider otellog.LoggerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSchemaURL(schemaURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSource(source bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewHandler(name string, opts ...Option) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func newOptions(opts []Option) options { _ = "STUB: not implemented"; return *new(options) }

func newHandler(name string, cfg options) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

type traceHandler struct {
	next slog.Handler
}

func (h *traceHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *traceHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *traceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *traceHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func traceAttrs(ctx context.Context) []slog.Attr { _ = "STUB: not implemented"; return nil }
