package log

import (
	"context"
	"log/slog"
)

type ctxAttrsKey struct{}

func ContextWithAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func AttrsFromContext(ctx context.Context) []slog.Attr { _ = "STUB: not implemented"; return nil }

func newContextHandler(next slog.Handler, extractors ...Extractor) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func compactExtractors(extractors []Extractor) []Extractor { _ = "STUB: not implemented"; return nil }

type contextHandler struct {
	next       slog.Handler
	extractors []Extractor
}

func (h *contextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *contextHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *contextHandler) attrs(ctx context.Context) []slog.Attr {
	_ = "STUB: not implemented"
	return nil
}

func (h *contextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *contextHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}
