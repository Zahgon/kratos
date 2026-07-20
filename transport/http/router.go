package http

type WalkRouteFunc func(RouteInfo) error

type RouteInfo struct {
	Path   string
	Method string
}

type HandlerFunc func(Context) error

type Router struct {
	prefix  string
	srv     *Server
	filters []FilterFunc
}

func newRouter(prefix string, srv *Server, filters ...FilterFunc) *Router {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) Group(prefix string, filters ...FilterFunc) *Router {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) Handle(method, relativePath string, h HandlerFunc, filters ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) GET(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) HEAD(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) POST(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) PUT(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) PATCH(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) DELETE(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) CONNECT(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) OPTIONS(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) TRACE(path string, h HandlerFunc, m ...FilterFunc) {
	_ = "STUB: not implemented"
	return
}
