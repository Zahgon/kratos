package ratelimit

import (
	"github.com/go-kratos/kratos/v3/errors"
	internalratelimit "github.com/go-kratos/kratos/v3/internal/ratelimit"
	"github.com/go-kratos/kratos/v3/middleware"
)

var ErrLimitExceed = errors.New(429, "RATELIMIT", "service unavailable due to rate limit exceeded")

type DoneFunc = internalratelimit.DoneFunc

type DoneInfo = internalratelimit.DoneInfo

type Limiter = internalratelimit.Limiter

type Option func(*options)

func WithLimiter(limiter Limiter) Option { _ = "STUB: not implemented"; return *new(Option) }

type options struct {
	limiter Limiter
}

func Server(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}
