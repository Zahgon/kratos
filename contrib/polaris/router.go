package polaris

import (
	"github.com/polarismesh/polaris-go/pkg/model"
	"github.com/polarismesh/polaris-go/pkg/model/pb"

	"github.com/go-kratos/kratos/v3/selector"
)

type router struct {
	service string
}

type RouterOption func(o *router)

func WithRouterService(service string) RouterOption {
	_ = "STUB: not implemented"
	return *new(RouterOption)
}

func (p *Polaris) NodeFilter(opts ...RouterOption) selector.NodeFilter {
	_ = "STUB: not implemented"
	return *new(selector.NodeFilter)
}

func (p *Polaris) processRouters(sourceService *model.ServiceInfo, dstInstances *pb.ServiceInstancesInProto) ([]model.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Polaris) getRouteRule(service model.ServiceMetadata) (model.ServiceRule, error) {
	_ = "STUB: not implemented"
	return *new(model.ServiceRule), nil
}

type serviceRuleResponse struct {
	resp *model.ServiceRuleResponse
}

func (r serviceRuleResponse) GetType() model.EventType {
	_ = "STUB: not implemented"
	return *new(model.EventType)
}

func (r serviceRuleResponse) IsInitialized() bool { _ = "STUB: not implemented"; return false }

func (r serviceRuleResponse) GetRevision() string { _ = "STUB: not implemented"; return "" }

func (r serviceRuleResponse) GetHashValue() uint64 { _ = "STUB: not implemented"; return 0 }

func (r serviceRuleResponse) IsNotExists() bool { _ = "STUB: not implemented"; return false }

func (r serviceRuleResponse) GetNamespace() string { _ = "STUB: not implemented"; return "" }

func (r serviceRuleResponse) GetService() string { _ = "STUB: not implemented"; return "" }

func (r serviceRuleResponse) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func (r serviceRuleResponse) GetRuleCache() model.RuleCache {
	_ = "STUB: not implemented"
	return *new(model.RuleCache)
}

func (r serviceRuleResponse) GetValidateError() error { _ = "STUB: not implemented"; return nil }

func (r serviceRuleResponse) IsCacheLoaded() bool { _ = "STUB: not implemented"; return false }

func buildPolarisInstance(namespace string, nodes []selector.Node) *pb.ServiceInstancesInProto {
	_ = "STUB: not implemented"
	return nil
}
