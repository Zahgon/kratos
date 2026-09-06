package logging

import (
	"log/slog"

	"github.com/go-kratos/kratos/v3/middleware"
)

type Redacter interface {
	Redact() string
}

func Server(logger *slog.Logger) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func Client(logger *slog.Logger) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func extractArgs(req any) string { _ = "STUB: not implemented"; return "" }

func extractError(err error) (slog.Level, string) {
	_ = "STUB: not implemented"
	return *new(slog.Level), ""
}
