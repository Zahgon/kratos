package ewma

import (
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v3/selector"
)

const (
	tau = int64(time.Millisecond * 600)

	penalty = uint64(time.Microsecond * 100)
)

var (
	_ selector.WeightedNode        = (*Node)(nil)
	_ selector.WeightedNodeBuilder = (*Builder)(nil)
)

type Node struct {
	selector.Node

	lag       atomic.Int64
	success   atomic.Uint64
	inflight  atomic.Int64
	inflights [200]atomic.Int64

	stamp atomic.Int64

	reqs atomic.Int64

	lastPick atomic.Int64

	errHandler   func(err error) (isErr bool)
	cachedWeight *atomic.Value
}

type nodeWeight struct {
	value    float64
	updateAt int64
}

type Builder struct {
	ErrHandler func(err error) (isErr bool)
}

func (b *Builder) Build(n selector.Node) selector.WeightedNode {
	_ = "STUB: not implemented"
	return *new(selector.WeightedNode)
}

func (n *Node) health() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *Node) load() (load uint64) { _ = "STUB: not implemented"; return 0 }

func (n *Node) predict(avgLag int64, now int64) (predict int64) {
	_ = "STUB: not implemented"
	return 0
}

func (n *Node) Pick() selector.DoneFunc { _ = "STUB: not implemented"; return *new(selector.DoneFunc) }

func (n *Node) Weight() (weight float64) { _ = "STUB: not implemented"; return 0 }

func (n *Node) PickElapsed() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (n *Node) Raw() selector.Node { _ = "STUB: not implemented"; return *new(selector.Node) }
