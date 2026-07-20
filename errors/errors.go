package errors

import (
	"google.golang.org/grpc/status"
)

const (
	UnknownCode = 500

	UnknownReason = ""

	SupportPackageIsVersion1 = true
)

type Error struct {
	Status
	cause error
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *Error) Is(err error) bool { _ = "STUB: not implemented"; return false }

func (e *Error) WithCause(cause error) *Error { _ = "STUB: not implemented"; return nil }

func (e *Error) WithMetadata(md map[string]string) *Error { _ = "STUB: not implemented"; return nil }

func (e *Error) GRPCStatus() *status.Status { _ = "STUB: not implemented"; return nil }

func New(code int, reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func Newf(code int, reason, format string, a ...any) *Error { _ = "STUB: not implemented"; return nil }

func Errorf(code int, reason, format string, a ...any) error { _ = "STUB: not implemented"; return nil }

func Code(err error) int { _ = "STUB: not implemented"; return 0 }

//nolint:mnd

func Reason(err error) string { _ = "STUB: not implemented"; return "" }

func Clone(err *Error) *Error { _ = "STUB: not implemented"; return nil }

func FromError(err error) *Error { _ = "STUB: not implemented"; return nil }
