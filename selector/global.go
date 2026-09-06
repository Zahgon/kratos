package selector

var globalSelector = &wrapSelector{}

var _ Builder = (*wrapSelector)(nil)

type wrapSelector struct{ Builder }

func GlobalSelector() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func SetGlobalSelector(builder Builder) { _ = "STUB: not implemented"; return }
