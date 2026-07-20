package log

import (
	"context"
	"io"
	"log/slog"
)

type Format int

const (
	FormatText Format = iota

	FormatJSON
)

type Option func(*handlerConfig)

type Extractor func(context.Context) []slog.Attr

type handlerConfig struct {
	writer      io.Writer
	format      Format
	level       Leveler
	addSource   bool
	replaceAttr func(groups []string, a slog.Attr) slog.Attr
	extractors  []Extractor
	filter      []FilterOption
}

func WithExtractor(extractors ...Extractor) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWriter(w io.Writer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFormat(f Format) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLevel(l Leveler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAddSource(b bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReplaceAttr(fn func(groups []string, a slog.Attr) slog.Attr) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithFilter(opts ...FilterOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewHandler(opts ...Option) slog.Handler { _ = "STUB: not implemented"; return *new(slog.Handler) }

func NewLogger(handler slog.Handler, opts ...Option) *slog.Logger {
	_ = "STUB: not implemented"
	return nil
}

func newComposedHandler(h slog.Handler, cfg *handlerConfig) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func newBaseHandler(cfg *handlerConfig) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}
