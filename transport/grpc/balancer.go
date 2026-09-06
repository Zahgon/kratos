package grpc

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/metadata"

	"github.com/go-kratos/kratos/v3/selector"
)

const (
	balancerName = "selector"
)

var (
	_ base.PickerBuilder = (*balancerBuilder)(nil)
	_ balancer.Picker    = (*balancerPicker)(nil)
)

func init() {
	b := base.NewBalancerBuilder(
		balancerName,
		&balancerBuilder{
			builder: selector.GlobalSelector(),
		},
		base.Config{HealthCheck: true},
	)
	balancer.Register(b)
}

type balancerBuilder struct {
	builder selector.Builder
}

func (b *balancerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	_ = "STUB: not implemented"
	return *new(balancer.Picker)
}

type balancerPicker struct {
	selector selector.Selector
}

func (p *balancerPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	_ = "STUB: not implemented"
	return *new(balancer.PickResult), nil
}

type Trailer metadata.MD

func (t Trailer) Get(k string) string { _ = "STUB: not implemented"; return "" }

type grpcNode struct {
	selector.Node
	subConn balancer.SubConn
}
