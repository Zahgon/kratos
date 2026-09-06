package http

import (
	"regexp"
)

var pathTemplateParamRE = regexp.MustCompile(`{([.\w]+)(=[^{}]*)?}`)

type BuildPathOption func(*buildPathOptions)

type buildPathOptions struct {
	queryParams bool
	omitFields  []string
}

func WithQueryParams() BuildPathOption { _ = "STUB: not implemented"; return *new(BuildPathOption) }

func WithOmitFields(fields ...string) BuildPathOption {
	_ = "STUB: not implemented"
	return *new(BuildPathOption)
}

func BuildPath(pathTemplate string, msg any, opts ...BuildPathOption) string {
	_ = "STUB: not implemented"
	return ""
}

func omitQueryParams(values map[string][]string, fields []string) {
	_ = "STUB: not implemented"
	return
}
