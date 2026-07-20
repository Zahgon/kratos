package sentry

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/go-kratos/kratos/v3/middleware"
)

type ctxKey struct{}

type Option func(*options)

type options struct {
	repanic         bool
	waitForDelivery bool
	timeout         time.Duration
	tags            map[string]string
	contextTags     func(context.Context) map[string]string
}

func WithRepanic(repanic bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWaitForDelivery(waitForDelivery bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTags(kvs map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithContextTags(fn func(context.Context) map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Server(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func recoverWithSentry(ctx context.Context, opts options, hub *sentry.Hub, req any) {
	_ = "STUB: not implemented"
	return
}

func isBrokenPipeError(err any) bool { _ = "STUB: not implemented"; return false }

func GetHubFromContext(ctx context.Context) *sentry.Hub { _ = "STUB: not implemented"; return nil }
