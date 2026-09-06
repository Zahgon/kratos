package config

import (
	"regexp"
)

type Decoder func(*KeyValue, map[string]any) error

type Resolver func(map[string]any) error

type Merge func(dst, src any) error

type Option func(*options)

type options struct {
	sources  []Source
	decoder  Decoder
	resolver Resolver
	merge    Merge
}

const (
	boolTrueValue  = "true"
	boolFalseValue = "false"
)

func WithSource(s ...Source) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDecoder(d Decoder) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithResolveActualTypes(enableConvertToType bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithResolver(r Resolver) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMergeFunc(m Merge) Option { _ = "STUB: not implemented"; return *new(Option) }

func defaultDecoder(src *KeyValue, target map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func newActualTypesResolver(enableConvertToType bool) func(map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultResolver(input map[string]any) error { _ = "STUB: not implemented"; return nil }

func resolver(input map[string]any, mapper func(name string) string, toType bool) error {
	_ = "STUB: not implemented"
	return nil
}

func mapper(input map[string]any) func(name string) string { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func convertToType(input string) any { _ = "STUB: not implemented"; return *new(any) }

var placeholderRegexp = regexp.MustCompile(`\${(.*?)}`)

func expand(s string, mapping func(string) string, toType bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

//nolint:mnd
