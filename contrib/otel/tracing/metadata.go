package tracing

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
)

const serviceHeader = "x-md-service-name"

type Metadata struct{}

var _ propagation.TextMapPropagator = Metadata{}

func (b Metadata) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

func (b Metadata) Extract(parent context.Context, carrier propagation.TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (b Metadata) Fields() []string { _ = "STUB: not implemented"; return nil }
