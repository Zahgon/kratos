package log

import (
	"context"
	"log/slog"
)

func SetDefault(logger *slog.Logger) { _ = "STUB: not implemented"; return }

func Default() *slog.Logger { _ = "STUB: not implemented"; return nil }

func With(args ...any) *slog.Logger { _ = "STUB: not implemented"; return nil }

func WithGroup(name string) *slog.Logger { _ = "STUB: not implemented"; return nil }

func Handler() slog.Handler { _ = "STUB: not implemented"; return *new(slog.Handler) }

func Enabled(ctx context.Context, level Level) bool { _ = "STUB: not implemented"; return false }

func Debug(msg string, args ...any) { _ = "STUB: not implemented"; return }

func DebugContext(ctx context.Context, msg string, args ...any) { _ = "STUB: not implemented"; return }

func Info(msg string, args ...any) { _ = "STUB: not implemented"; return }

func InfoContext(ctx context.Context, msg string, args ...any) { _ = "STUB: not implemented"; return }

func Warn(msg string, args ...any) { _ = "STUB: not implemented"; return }

func WarnContext(ctx context.Context, msg string, args ...any) { _ = "STUB: not implemented"; return }

func Error(msg string, args ...any) { _ = "STUB: not implemented"; return }

func ErrorContext(ctx context.Context, msg string, args ...any) { _ = "STUB: not implemented"; return }

func Log(ctx context.Context, level Level, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

//nolint:revive // LogAttrs intentionally mirrors slog.Logger.LogAttrs.
func LogAttrs(ctx context.Context, level Level, msg string, attrs ...slog.Attr) {
	_ = "STUB: not implemented"
	return
}

func log(ctx context.Context, level Level, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}
