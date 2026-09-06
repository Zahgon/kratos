package errors

func BadRequest(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsBadRequest(err error) bool { _ = "STUB: not implemented"; return false }

func Unauthorized(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsUnauthorized(err error) bool { _ = "STUB: not implemented"; return false }

func Forbidden(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsForbidden(err error) bool { _ = "STUB: not implemented"; return false }

func NotFound(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsNotFound(err error) bool { _ = "STUB: not implemented"; return false }

func Conflict(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsConflict(err error) bool { _ = "STUB: not implemented"; return false }

func TooManyRequests(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsTooManyRequests(err error) bool { _ = "STUB: not implemented"; return false }

func ClientClosed(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsClientClosed(err error) bool { _ = "STUB: not implemented"; return false }

func InternalServer(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsInternalServer(err error) bool { _ = "STUB: not implemented"; return false }

func ServiceUnavailable(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsServiceUnavailable(err error) bool { _ = "STUB: not implemented"; return false }

func GatewayTimeout(reason, message string) *Error { _ = "STUB: not implemented"; return nil }

func IsGatewayTimeout(err error) bool { _ = "STUB: not implemented"; return false }
