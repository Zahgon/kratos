package kuberegistry

import (
	"context"
	"errors"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listerv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/go-kratos/kratos/v3/registry"
)

const (
	LabelsKeyServiceID = "kratos-service-id"

	LabelsKeyServiceName = "kratos-service-app"

	LabelsKeyServiceVersion = "kratos-service-version"

	AnnotationsKeyMetadata = "kratos-service-metadata"

	AnnotationsKeyProtocolMap = "kratos-service-protocols"
)

type Registry struct {
	clientSet       *kubernetes.Clientset
	informerFactory informers.SharedInformerFactory
	podInformer     cache.SharedIndexInformer
	podLister       listerv1.PodLister

	stopCh chan struct{}
}

func NewRegistry(clientSet *kubernetes.Clientset, namespace string) *Registry {
	_ = "STUB: not implemented"
	return nil
}

func (s *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Registry) Deregister(ctx context.Context, _ *registry.ServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Registry) GetService(_ context.Context, name string) ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Registry) sendLatestInstances(ctx context.Context, name string, announcement chan []*registry.ServiceInstance) {
	_ = "STUB: not implemented"
	return
}

func (s *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(registry.Watcher), nil
}

func (s *Registry) Start() { _ = "STUB: not implemented"; return }

func (s *Registry) Close() { _ = "STUB: not implemented"; return }

const ServiceAccountNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

var currentNamespace = LoadNamespace()

func LoadNamespace() string { _ = "STUB: not implemented"; return "" }

func GetNamespace() string { _ = "STUB: not implemented"; return "" }

func GetPodName() string { _ = "STUB: not implemented"; return "" }

type protocolMap map[string]string

func (m protocolMap) GetProtocol(port int32) string { _ = "STUB: not implemented"; return "" }

type Iterator struct {
	ch     chan []*registry.ServiceInstance
	stopCh chan struct{}
}

func NewIterator(channel chan []*registry.ServiceInstance, stopCh chan struct{}) *Iterator {
	_ = "STUB: not implemented"
	return nil
}

func (iter *Iterator) Next() ([]*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (iter *Iterator) Stop() error { _ = "STUB: not implemented"; return nil }

func marshal(in any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func unmarshal(data string, in any) error { _ = "STUB: not implemented"; return nil }

func isEmptyObjectString(s string) bool { _ = "STUB: not implemented"; return false }

func getProtocolMapByEndpoints(endpoints []string) (protocolMap, error) {
	_ = "STUB: not implemented"
	return *new(protocolMap), nil
}

func getProtocolMapFromPod(pod *corev1.Pod) (protocolMap, error) {
	_ = "STUB: not implemented"
	return *new(protocolMap), nil
}

func getMetadataFromPod(pod *corev1.Pod) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getServiceInstanceFromPod(pod *corev1.Pod) (*registry.ServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ErrIteratorClosed = errors.New("iterator closed")

type ErrorHandleResource struct {
	Namespace string
	Name      string
	Reason    error
}

func (err *ErrorHandleResource) Error() string { _ = "STUB: not implemented"; return "" }
