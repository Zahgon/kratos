package etcd

import (
	"github.com/go-kratos/kratos/v3/registry"
)

func marshal(si *registry.ServiceInstance) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func unmarshal(data []byte) (si *registry.ServiceInstance, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
