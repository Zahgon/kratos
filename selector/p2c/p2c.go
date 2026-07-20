package p2c

import (
	"context"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v3/selector"
)

const (
	forcePick = time.Second * 3

	Name = "p2c"
)

var _ selector.Balancer = (*Balancer)(nil)

type Option func(o *options)

type options struct{}

func New(opts ...Option) selector.Selector {
	_ = "STUB: not implemented"
	return *new(selector.Selector)
}

type Balancer struct {
	mu     sync.Mutex
	r      *rand.Rand
	picked atomic.Bool
}

func (s *Balancer) prePick(nodes []selector.WeightedNode) (nodeA selector.WeightedNode, nodeB selector.WeightedNode) {
	_ = "STUB: not implemented"
	return *new(selector.WeightedNode), *new(selector.WeightedNode)
}

func (s *Balancer) Pick(_ context.Context, nodes []selector.WeightedNode) (selector.WeightedNode, selector.DoneFunc, error) {
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
