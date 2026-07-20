package selector

import (
	"context"
)

type peerKey struct{}

type Peer struct {
	Node Node
}

func NewPeerContext(ctx context.Context, p *Peer) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromPeerContext(ctx context.Context) (p *Peer, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
