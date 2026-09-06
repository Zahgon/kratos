package tracing

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	tracer trace.Tracer
	kind   trace.SpanKind
	opt    *options
}

func NewTracer(kind trace.SpanKind, opts ...Option) *Tracer { _ = "STUB: not implemented"; return nil }

func (t *Tracer) Start(ctx context.Context, operation string, carrier propagation.TextMapCarrier) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func (t *Tracer) End(_ context.Context, span trace.Span, m any, err error) {
	_ = "STUB: not implemented"
	return
}
