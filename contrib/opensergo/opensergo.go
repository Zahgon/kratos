package opensergo

import (
	"context"

	v1 "github.com/opensergo/opensergo-go/proto/service_contract/v1"

	"github.com/go-kratos/kratos/v3"
)

type Option func(*options)

func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

type options struct {
	Endpoint string `json:"endpoint"`
}

func (o *options) ParseJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type OpenSergo struct {
	mdClient v1.MetadataServiceClient
}

func New(opts ...Option) (*OpenSergo, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *OpenSergo) ReportMetadata(ctx context.Context, app kratos.AppInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func listDescriptors() (services []*v1.ServiceDescriptor, types []*v1.TypeDescriptor, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func HTTPPatternInfo(pattern any) (method string, path string) {
	_ = "STUB: not implemented"
	return "", ""
}
