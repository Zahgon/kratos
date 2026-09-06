package selector

import (
	"context"

	"github.com/go-kratos/kratos/v3/errors"
)

var ErrNoAvailable = errors.ServiceUnavailable("no_available_node", "")

type Selector interface {
	Rebalancer

	Select(ctx context.Context, opts ...SelectOption) (selected Node, done DoneFunc, err error)
}

type Rebalancer interface {
	Apply(nodes []Node)
}

type Builder interface {
	Build() Selector
}

type Node interface {
	Scheme() string

	Address() string

	ServiceName() string

	InitialWeight() *int64

	Version() string

	Metadata() map[string]string
}

type DoneInfo struct {
	Err error

	ReplyMD ReplyMD

	BytesSent bool

	BytesReceived bool
}

type ReplyMD interface {
	Get(key string) string
}

type DoneFunc func(ctx context.Context, di DoneInfo)
