package selector

import (
	"context"
	"regexp"

	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/transport"
)

type (
	transporter func(ctx context.Context) (transport.Transporter, bool)
	MatchFunc   func(ctx context.Context, operation string) bool
)

var (
	serverTransporter transporter = func(ctx context.Context) (transport.Transporter, bool) {
		return transport.FromServerContext(ctx)
	}

	clientTransporter transporter = func(ctx context.Context) (transport.Transporter, bool) {
		return transport.FromClientContext(ctx)
	}
)

type Builder struct {
	client bool

	prefix   []string
	regex    []string
	path     []string
	match    MatchFunc
	compiled []*regexp.Regexp

	ms []middleware.Middleware
}

func Server(ms ...middleware.Middleware) *Builder { _ = "STUB: not implemented"; return nil }

func Client(ms ...middleware.Middleware) *Builder { _ = "STUB: not implemented"; return nil }

func (b *Builder) Prefix(prefix ...string) *Builder { _ = "STUB: not implemented"; return nil }

func (b *Builder) Regex(regex ...string) *Builder { _ = "STUB: not implemented"; return nil }

func (b *Builder) Path(path ...string) *Builder { _ = "STUB: not implemented"; return nil }

func (b *Builder) Match(fn MatchFunc) *Builder { _ = "STUB: not implemented"; return nil }

func (b *Builder) Build() middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func (b *Builder) matches(ctx context.Context, transporter transporter) bool {
	_ = "STUB: not implemented"
	return false
}

func selector(transporter transporter, match func(context.Context, transporter) bool, ms ...middleware.Middleware) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func pathMatch(path string, operation string) bool { _ = "STUB: not implemented"; return false }

func prefixMatch(prefix string, operation string) bool { _ = "STUB: not implemented"; return false }
