package main

import (
	_ "embed"
)

//go:embed httpTemplate.tpl
var httpTemplate string

type serviceDesc struct {
	ServiceType string
	ServiceName string
	Metadata    string
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	Name         string
	OriginalName string
	Num          int
	Request      string
	Reply        string
	Comment      string

	Path                 string
	PathTemplate         string
	Method               string
	HasVars              bool
	HasBody              bool
	Body                 string
	BodyField            string
	BodyQueryName        string
	BodyHTTPBody         bool
	BodyMessage          bool
	ResponseBody         string
	ResponseBodyHTTPBody bool
	ReplyHTTPBody        bool
	ClientStreaming      bool
	ServerStreaming      bool
}

func (s *serviceDesc) execute() string { _ = "STUB: not implemented"; return "" }
