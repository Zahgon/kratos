package direct

import "google.golang.org/grpc/resolver"

type directResolver struct{}

func newDirectResolver() resolver.Resolver {
	_ = "STUB: not implemented"
	return *new(resolver.Resolver)
}

func (r *directResolver) Close() { _ = "STUB: not implemented"; return }

func (r *directResolver) ResolveNow(_ resolver.ResolveNowOptions) {
	_ = "STUB: not implemented"
	return
}
