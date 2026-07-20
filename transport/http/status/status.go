package status

import (
	"google.golang.org/grpc/codes"
)

const (
	ClientClosed = 499
)

type Converter interface {
	ToGRPCCode(code int) codes.Code

	FromGRPCCode(code codes.Code) int
}

type statusConverter struct{}

var DefaultConverter Converter = statusConverter{}

func (c statusConverter) ToGRPCCode(code int) codes.Code {
	_ = "STUB: not implemented"
	return *new(codes.Code)
}

func (c statusConverter) FromGRPCCode(code codes.Code) int { _ = "STUB: not implemented"; return 0 }

func ToGRPCCode(code int) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

func FromGRPCCode(code codes.Code) int { _ = "STUB: not implemented"; return 0 }
