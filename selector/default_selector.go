package selector

import (
	"context"
	"sync/atomic"
)

var (
	_ Rebalancer = (*Default)(nil)
	_ Builder    = (*DefaultBuilder)(nil)
)

type Default struct {
	NodeBuilder WeightedNodeBuilder
	Balancer    Balancer

	nodes atomic.Value
}

func (d *Default) Select(ctx context.Context, opts ...SelectOption) (selected Node, done DoneFunc, err error) {
	_ = "STUB: not implemented"
	return *new(Node), *new(DoneFunc), nil
}

func (d *Default) Apply(nodes []Node) { _ = "STUB: not implemented"; return }

type DefaultBuilder struct {
	Node     WeightedNodeBuilder
	Balancer BalancerBuilder
}

func (db *DefaultBuilder) Build() Selector { _ = "STUB: not implemented"; return *new(Selector) }
