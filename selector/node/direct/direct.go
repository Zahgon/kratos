package direct

import (
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v3/selector"
)

const (
	defaultWeight = 100
)

var (
	_ selector.WeightedNode        = (*Node)(nil)
	_ selector.WeightedNodeBuilder = (*Builder)(nil)
)

type Node struct {
	selector.Node

	lastPick atomic.Int64
}

type Builder struct{}

func (*Builder) Build(n selector.Node) selector.WeightedNode {
	_ = "STUB: not implemented"
	return *new(selector.WeightedNode)
}

func (n *Node) Pick() selector.DoneFunc { _ = "STUB: not implemented"; return *new(selector.DoneFunc) }

func (n *Node) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (n *Node) PickElapsed() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (n *Node) Raw() selector.Node { _ = "STUB: not implemented"; return *new(selector.Node) }
