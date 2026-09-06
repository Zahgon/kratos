package http

import "net/http"

type FilterFunc func(http.Handler) http.Handler

func FilterChain(filters ...FilterFunc) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}
