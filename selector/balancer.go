package selector

import (
	"context"
	"time"
)

type Balancer interface {
	Pick(ctx context.Context, nodes []WeightedNode) (selected WeightedNode, done DoneFunc, err error)
}

type BalancerBuilder interface {
	Build() Balancer
}

type WeightedNode interface {
	Node

	Raw() Node

	Weight() float64

	Pick() DoneFunc

	PickElapsed() time.Duration
}

type WeightedNodeBuilder interface {
	Build(Node) WeightedNode
}
