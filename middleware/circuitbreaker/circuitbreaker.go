package circuitbreaker

import (
	"github.com/go-kratos/kratos/v3/errors"
	internalbreaker "github.com/go-kratos/kratos/v3/internal/circuitbreaker"
	"github.com/go-kratos/kratos/v3/internal/group"
	"github.com/go-kratos/kratos/v3/middleware"
)

var ErrNotAllowed = errors.New(503, "CIRCUITBREAKER", "request failed due to circuit breaker triggered")

type CircuitBreaker = internalbreaker.CircuitBreaker

type Option func(*options)

func WithBreakerFactory(factory func() CircuitBreaker) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type options struct {
	group *group.Group[CircuitBreaker]
}

func Client(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}
