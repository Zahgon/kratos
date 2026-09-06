package validate

import (
	"github.com/go-kratos/kratos/v3/middleware"
)

type ValidatorFunc func(v any) error

type validator interface {
	Validate() error
}

func Validator(validators ...ValidatorFunc) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}
