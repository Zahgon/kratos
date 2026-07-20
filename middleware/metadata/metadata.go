package metadata

import (
	"github.com/go-kratos/kratos/v3/metadata"
	"github.com/go-kratos/kratos/v3/middleware"
)

type Option func(*options)

type options struct {
	prefix []string
	md     metadata.Metadata
}

func (o *options) hasPrefix(key string) bool { _ = "STUB: not implemented"; return false }

func WithConstants(md metadata.Metadata) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPropagatedPrefix(prefix ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Server(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func Client(opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}
