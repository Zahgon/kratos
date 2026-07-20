package http

import (
	"net/http"
)

const (
	contentTypeJSON = "application/json"
	schemeDiscovery = "discovery"
	schemeHTTP      = "http"
	schemeHTTPS     = "https"
)

type CallOption interface {
	before(*callInfo) error

	after(*callInfo, *csAttempt)
}

type callInfo struct {
	contentType    string
	contentTypeSet bool
	accept         string
	operation      string
	pathTemplate   string
	headerCarrier  *http.Header
}

type EmptyCallOption struct{}

func (EmptyCallOption) before(*callInfo) error      { _ = "STUB: not implemented"; return nil }
func (EmptyCallOption) after(*callInfo, *csAttempt) { _ = "STUB: not implemented"; return }

type csAttempt struct {
	res *http.Response
}

func ContentType(contentType string) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type ContentTypeCallOption struct {
	EmptyCallOption
	ContentType string
}

func (o ContentTypeCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func Accept(contentType string) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type AcceptCallOption struct {
	EmptyCallOption
	ContentType string
}

func (o AcceptCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func defaultCallInfo(path string) callInfo { _ = "STUB: not implemented"; return *new(callInfo) }

func Operation(operation string) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type OperationCallOption struct {
	EmptyCallOption
	Operation string
}

func (o OperationCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func PathTemplate(pattern string) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type PathTemplateCallOption struct {
	EmptyCallOption
	Pattern string
}

func (o PathTemplateCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func Header(header *http.Header) CallOption { _ = "STUB: not implemented"; return *new(CallOption) }

type HeaderCallOption struct {
	EmptyCallOption
	header *http.Header
}

func (o HeaderCallOption) before(c *callInfo) error { _ = "STUB: not implemented"; return nil }

func (o HeaderCallOption) after(_ *callInfo, cs *csAttempt) { _ = "STUB: not implemented"; return }
