package middleware

import (
	"context"
)

type Handler func(ctx context.Context, req any) (any, error)

type Middleware func(Handler) Handler

func Chain(m ...Middleware) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }
