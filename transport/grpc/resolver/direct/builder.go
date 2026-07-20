package direct

import (
	"google.golang.org/grpc/resolver"
)

const name = "direct"

func init() {
	resolver.Register(NewBuilder())
}

type directBuilder struct{}

func NewBuilder() resolver.Builder { _ = "STUB: not implemented"; return *new(resolver.Builder) }

func (d *directBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver), nil
}

func (d *directBuilder) Scheme() string { _ = "STUB: not implemented"; return "" }
