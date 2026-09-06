package discovery

import (
	"errors"
	"time"

	"google.golang.org/grpc/resolver"

	"github.com/go-kratos/kratos/v3/registry"
)

const name = "discovery"

var ErrWatcherCreateTimeout = errors.New("discovery create watcher overtime")

type Option func(o *builder)

func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithInsecure(insecure bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSubset(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

type builder struct {
	discoverer registry.Discovery
	timeout    time.Duration
	insecure   bool
	subsetSize int
}

func NewBuilder(d registry.Discovery, opts ...Option) resolver.Builder {
	_ = "STUB: not implemented"
	return *new(resolver.Builder)
}

func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver), nil
}

func (*builder) Scheme() string { _ = "STUB: not implemented"; return "" }
