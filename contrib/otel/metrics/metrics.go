package metrics

import (
	"go.opentelemetry.io/otel/metric"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"

	"github.com/go-kratos/kratos/v3/middleware"
)

const (
	metricLabelKind      = "kind"
	metricLabelOperation = "operation"
	metricLabelCode      = "code"
	metricLabelReason    = "reason"
)

const (
	DefaultServerSecondsHistogramName = "server_requests_seconds"
	DefaultServerRequestsCounterName  = "server_requests_code_total"
	DefaultClientSecondsHistogramName = "client_requests_seconds"
	DefaultClientRequestsCounterName  = "client_requests_code_total"
)

type Option func(*options)

func WithRequests(c metric.Int64Counter) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSeconds(histogram metric.Float64Histogram) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func DefaultRequestsCounter(meter metric.Meter, histogramName string) (metric.Int64Counter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter), nil
}

func DefaultSecondsHistogram(meter metric.Meter, histogramName string) (metric.Float64Histogram, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram), nil
}

func DefaultSecondsHistogramView(histogramName string) metricsdk.View {
	_ = "STUB: not implemented"
	return *new(metricsdk.View)
}

type options struct {
	requests metric.Int64Counter

	seconds metric.Float64Histogram
}

func Server(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func Client(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}
