package wrr

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v3/selector"
)

const (
	Name = "wrr"
)

var _ selector.Balancer = (*Balancer)(nil)

type Option func(o *options)

type options struct{}

type Balancer struct {
	mu            sync.Mutex
	currentWeight map[string]float64
	lastNodes     []selector.WeightedNode
}

func equalNodes(a, b []selector.WeightedNode) bool { _ = "STUB: not implemented"; return false }

func New(opts ...Option) selector.Selector {
	_ = "STUB: not implemented"
	return *new(selector.Selector)
}

func (p *Balancer) Pick(_ context.Context, nodes []selector.WeightedNode) (selector.WeightedNode, selector.DoneFunc, error) {
	_ = "STUB: not implemented"
	return *new(selector.WeightedNode), *new(selector.DoneFunc), nil
}

func NewBuilder(opts ...Option) selector.Builder {
	_ = "STUB: not implemented"
	return *new(selector.Builder)
}

type Builder struct{}

func (b *Builder) Build() selector.Balancer {
	_ = "STUB: not implemented"
	return *new(selector.Balancer)
}
