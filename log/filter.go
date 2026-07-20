package log

import (
	"context"
	"log/slog"
)

const redactedValue = "***"

type FilterOption func(*filterConfig)

type filterConfig struct {
	keys   map[string]struct{}
	filter func(ctx context.Context, record slog.Record) bool
}

func FilterKey(keys ...string) FilterOption { _ = "STUB: not implemented"; return *new(FilterOption) }

func FilterFunc(fn func(ctx context.Context, record slog.Record) bool) FilterOption {
	_ = "STUB: not implemented"
	return *new(FilterOption)
}

func newFilterHandler(next slog.Handler, opts ...FilterOption) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

type filterHandler struct {
	next   slog.Handler
	cfg    *filterConfig
	groups []string
}

func (h *filterHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *filterHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *filterHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *filterHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *filterHandler) needsRewrite() bool { _ = "STUB: not implemented"; return false }

func (h *filterHandler) rewrite(record slog.Record) slog.Record {
	_ = "STUB: not implemented"
	return *new(slog.Record)
}

func (h *filterHandler) redactAttrs(groups []string, attrs []slog.Attr) []slog.Attr {
	_ = "STUB: not implemented"
	return nil
}

func (h *filterHandler) redactAttr(groups []string, a slog.Attr) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}

func (h *filterHandler) matchesKey(groups []string, key string) bool {
	_ = "STUB: not implemented"
	return false
}

func appendPath(groups []string, key string) []string { _ = "STUB: not implemented"; return nil }
