package log

import (
	"context"
	"log/slog"
)

type discardHandler struct{}

func (discardHandler) Enabled(context.Context, slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}
func (discardHandler) Handle(context.Context, slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}
func (h discardHandler) WithAttrs([]slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}
func (h discardHandler) WithGroup(string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}
