package jwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"

	"github.com/go-kratos/kratos/v3/errors"
	"github.com/go-kratos/kratos/v3/middleware"
)

type authKey struct{}

const (
	bearerWord string = "Bearer"

	bearerFormat string = "Bearer %s"

	authorizationKey string = "Authorization"

	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

type Option func(*options)

type options struct {
	signingMethod jwt.SigningMethod
	claims        func() jwt.Claims
	tokenHeader   map[string]any
	parserOptions []jwt.ParserOption
}

func WithSigningMethod(method jwt.SigningMethod) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithClaims(f func() jwt.Claims) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTokenHeader(header map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithParserOptions(opts ...jwt.ParserOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Server(keyFunc jwt.Keyfunc, opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func Client(keyProvider jwt.Keyfunc, opts ...Option) middleware.Middleware {
	_ = "STUB: not implemented"
	return *new(middleware.Middleware)
}

func NewContext(ctx context.Context, info jwt.Claims) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) (token jwt.Claims, ok bool) {
	_ = "STUB: not implemented"
	return *new(jwt.Claims), false
}
