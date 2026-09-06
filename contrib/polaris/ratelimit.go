package polaris

import (
	"github.com/go-kratos/kratos/v3/errors"
	"github.com/go-kratos/kratos/v3/middleware"
)

var (
	ErrLimitExceed = errors.New(429, "RATELIMIT", "service unavailable due to rate limit exceeded")
)

func Ratelimit(l Limiter) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}
