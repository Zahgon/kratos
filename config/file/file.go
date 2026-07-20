package file

import (
	"github.com/go-kratos/kratos/v3/config"
)

var _ config.Source = (*file)(nil)

type file struct {
	path string
}

func NewSource(path string) config.Source { _ = "STUB: not implemented"; return *new(config.Source) }

func (f *file) loadFile(path string) (*config.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *file) loadDir(path string) (kvs []*config.KeyValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *file) Load() (kvs []*config.KeyValue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *file) Watch() (config.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(config.Watcher), nil
}
