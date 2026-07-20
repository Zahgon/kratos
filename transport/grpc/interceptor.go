package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-kratos/kratos/v3/internal/matcher"
)

func (s *Server) unaryServerInterceptor() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

type wrappedStream struct {
	grpc.ServerStream
	ctx        context.Context
	middleware matcher.Matcher
}

func NewWrappedStream(ctx context.Context, stream grpc.ServerStream, m matcher.Matcher) grpc.ServerStream {
	_ = "STUB: not implemented"
	return *new(grpc.ServerStream)
}

func (w *wrappedStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *Server) streamServerInterceptor() grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

type stream struct {
	grpc.ServerStream
	streamMiddleware matcher.Matcher
}

func GetStream(ctx context.Context) grpc.ServerStream {
	_ = "STUB: not implemented"
	return *new(grpc.ServerStream)
}

func (w *wrappedStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (w *wrappedStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }
