package selector

type SelectOptions struct {
	NodeFilters []NodeFilter
}

type SelectOption func(*SelectOptions)

func WithNodeFilter(fn ...NodeFilter) SelectOption {
	_ = "STUB: not implemented"
	return *new(SelectOption)
}
