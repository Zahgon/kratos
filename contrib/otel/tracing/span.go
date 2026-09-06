package tracing

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func setClientSpan(ctx context.Context, span trace.Span, m any) { _ = "STUB: not implemented"; return }

func setServerSpan(ctx context.Context, span trace.Span, m any) { _ = "STUB: not implemented"; return }

func parseFullMethod(fullMethod string) (string, []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint:mnd

func peerAttr(addr string) []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func parseTarget(endpoint string) (address string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
