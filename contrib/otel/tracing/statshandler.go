package tracing

import (
	"context"

	"google.golang.org/grpc/stats"
)

type ClientHandler struct{}

func (c *ClientHandler) HandleConn(_ context.Context, _ stats.ConnStats) {
	_ = "STUB: not implemented"
	return
}

func (c *ClientHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *ClientHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

func (c *ClientHandler) TagRPC(ctx context.Context, _ *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
