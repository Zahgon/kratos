package main

import (
	_ "embed"
)

//go:embed errorsTemplate.tpl
var errorsTemplate string

type errorInfo struct {
	Name       string
	Value      string
	HTTPCode   int
	CamelValue string
	Comment    string
	HasComment bool
}

type errorWrapper struct {
	Errors []*errorInfo
}

func (e *errorWrapper) execute() string { _ = "STUB: not implemented"; return "" }
