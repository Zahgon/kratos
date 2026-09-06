package selector

import (
	"github.com/go-kratos/kratos/v3/registry"
)

var _ Node = (*DefaultNode)(nil)

type DefaultNode struct {
	scheme   string
	addr     string
	weight   *int64
	version  string
	name     string
	metadata map[string]string
}

func (n *DefaultNode) Scheme() string { _ = "STUB: not implemented"; return "" }

func (n *DefaultNode) Address() string { _ = "STUB: not implemented"; return "" }

func (n *DefaultNode) ServiceName() string { _ = "STUB: not implemented"; return "" }

func (n *DefaultNode) InitialWeight() *int64 { _ = "STUB: not implemented"; return nil }

func (n *DefaultNode) Version() string { _ = "STUB: not implemented"; return "" }

func (n *DefaultNode) Metadata() map[string]string { _ = "STUB: not implemented"; return nil }

func NewNode(scheme, addr string, ins *registry.ServiceInstance) Node {
	_ = "STUB: not implemented"
	return *new(Node)
}
