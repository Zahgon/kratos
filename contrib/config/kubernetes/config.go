package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/go-kratos/kratos/v3/config"
)

type Option func(*options)

type options struct {
	Namespace string

	LabelSelector string

	FieldSelector string

	KubeConfig string

	Master string
}

func Namespace(ns string) Option { _ = "STUB: not implemented"; return *new(Option) }

func LabelSelector(label string) Option { _ = "STUB: not implemented"; return *new(Option) }

func FieldSelector(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

func KubeConfig(config string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Master(master string) Option { _ = "STUB: not implemented"; return *new(Option) }

type kube struct {
	opts   options
	client *kubernetes.Clientset
}

func NewSource(opts ...Option) config.Source { _ = "STUB: not implemented"; return *new(config.Source) }

func (k *kube) init() (err error) {
	var config *rest.Config
	if k.opts.KubeConfig != "" {
		if config, err = clientcmd.BuildConfigFromFlags(k.opts.Master, k.opts.KubeConfig); err != nil {
			return err
		}
	} else {
		if config, err = rest.InClusterConfig(); err != nil {
			return err
		}
	}
	if k.client, err = kubernetes.NewForConfig(config); err != nil {
		return err
	}
	return nil
}

func (k *kube) load() (kvs []*config.KeyValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kube) configMap(cm v1.ConfigMap) (kvs []*config.KeyValue) {
	_ = "STUB: not implemented"
	return nil
}

func (k *kube) Load() ([]*config.KeyValue, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *kube) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}
