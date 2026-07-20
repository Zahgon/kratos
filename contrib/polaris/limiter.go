package polaris

import (
	"time"

	"github.com/go-kratos/kratos/v3/middleware/ratelimit"

	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
)

type (
	LimiterOption func(*limiterOptions)
)

type limiterOptions struct {
	namespace string

	service string

	timeout time.Duration

	retryCount int

	token uint32
}

func WithLimiterNamespace(namespace string) LimiterOption {
	_ = "STUB: not implemented"
	return *new(LimiterOption)
}

func WithLimiterService(service string) LimiterOption {
	_ = "STUB: not implemented"
	return *new(LimiterOption)
}

func WithLimiterTimeout(timeout time.Duration) LimiterOption {
	_ = "STUB: not implemented"
	return *new(LimiterOption)
}

func WithLimiterRetryCount(retryCount int) LimiterOption {
	_ = "STUB: not implemented"
	return *new(LimiterOption)
}

func WithLimiterToken(token uint32) LimiterOption {
	_ = "STUB: not implemented"
	return *new(LimiterOption)
}

type Limiter struct {
	limitAPI polaris.LimitAPI

	opts limiterOptions
}

func buildRequest(opts limiterOptions) polaris.QuotaRequest {
	_ = "STUB: not implemented"
	return *new(polaris.QuotaRequest)
}

func (l *Limiter) Allow(method string, argument ...model.Argument) (ratelimit.DoneFunc, error) {
	_ = "STUB: not implemented"
	return *new(ratelimit.DoneFunc), nil
}
