package http

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kratos/kratos/v3/middleware"
)

var _ Context = (*wrapper)(nil)

type Context interface {
	context.Context
	Vars() url.Values
	Query() url.Values
	Form() url.Values
	Header() http.Header
	Request() *http.Request
	Response() http.ResponseWriter
	Middleware(middleware.Handler) middleware.Handler
	Bind(any) error
	BindVars(any) error
	BindQuery(any) error
	BindForm(any) error
	Returns(any, error) error
	Result(int, any) error
	JSON(int, any) error
	XML(int, any) error
	String(int, string) error
	Blob(int, string, []byte) error
	Stream(int, string, io.Reader) error
	Reset(http.ResponseWriter, *http.Request)
}

type responseWriter struct {
	code int
	w    http.ResponseWriter
}

func (w *responseWriter) reset(res http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (w *responseWriter) Header() http.Header            { _ = "STUB: not implemented"; return *new(http.Header) }
func (w *responseWriter) WriteHeader(statusCode int)     { _ = "STUB: not implemented"; return }
func (w *responseWriter) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *responseWriter) Unwrap() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

type wrapper struct {
	router *Router
	req    *http.Request
	res    http.ResponseWriter
	w      responseWriter
}

func (c *wrapper) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (c *wrapper) Vars() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func (c *wrapper) Form() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func (c *wrapper) Query() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func (c *wrapper) Request() *http.Request { _ = "STUB: not implemented"; return nil }
func (c *wrapper) Response() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}
func (c *wrapper) Middleware(h middleware.Handler) middleware.Handler {
	_ = "STUB: not implemented"
	return *new(middleware.Handler)
}

func (c *wrapper) Bind(v any) error               { _ = "STUB: not implemented"; return nil }
func (c *wrapper) BindVars(v any) error           { _ = "STUB: not implemented"; return nil }
func (c *wrapper) BindQuery(v any) error          { _ = "STUB: not implemented"; return nil }
func (c *wrapper) BindForm(v any) error           { _ = "STUB: not implemented"; return nil }
func (c *wrapper) Returns(v any, err error) error { _ = "STUB: not implemented"; return nil }

func (c *wrapper) Result(code int, v any) error { _ = "STUB: not implemented"; return nil }

func (c *wrapper) JSON(code int, v any) error { _ = "STUB: not implemented"; return nil }

func (c *wrapper) XML(code int, v any) error { _ = "STUB: not implemented"; return nil }

func (c *wrapper) String(code int, text string) error { _ = "STUB: not implemented"; return nil }

func (c *wrapper) Blob(code int, contentType string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *wrapper) Stream(code int, contentType string, rd io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *wrapper) Reset(res http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (c *wrapper) Deadline() (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (c *wrapper) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *wrapper) Err() error { _ = "STUB: not implemented"; return nil }

func (c *wrapper) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }
