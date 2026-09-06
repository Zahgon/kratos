package errors

import (
	stderrors "errors"
)

var ErrUnsupported = stderrors.ErrUnsupported

func Is(err, target error) bool { _ = "STUB: not implemented"; return false }

func As(err error, target any) bool { _ = "STUB: not implemented"; return false }

func Unwrap(err error) error { _ = "STUB: not implemented"; return nil }

func Join(errs ...error) error { _ = "STUB: not implemented"; return nil }
