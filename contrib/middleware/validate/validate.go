package validate

import (
	"github.com/go-kratos/kratos/v3/middleware"
)

type validator interface {
	Validate() error
}

func ProtoValidate() middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}
