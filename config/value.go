package config

import (
	"sync/atomic"
	"time"
)

var (
	_ Value = (*atomicValue)(nil)
	_ Value = (*errValue)(nil)
)

type Value interface {
	Bool() (bool, error)
	Int() (int64, error)
	Float() (float64, error)
	String() (string, error)
	Duration() (time.Duration, error)
	Slice() ([]Value, error)
	Map() (map[string]Value, error)
	Scan(any) error
	Load() any
	Store(any)
}

type atomicValue struct {
	atomic.Value
}

func (v *atomicValue) typeAssertError() error { _ = "STUB: not implemented"; return nil }

func (v *atomicValue) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (v *atomicValue) Int() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *atomicValue) Slice() ([]Value, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *atomicValue) Map() (map[string]Value, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *atomicValue) Float() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *atomicValue) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (v *atomicValue) Duration() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (v *atomicValue) Scan(obj any) error { _ = "STUB: not implemented"; return nil }

type errValue struct {
	err error
}

func (v errValue) Bool() (bool, error)     { _ = "STUB: not implemented"; return false, nil }
func (v errValue) Int() (int64, error)     { _ = "STUB: not implemented"; return 0, nil }
func (v errValue) Float() (float64, error) { _ = "STUB: not implemented"; return 0, nil }
func (v errValue) Duration() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
func (v errValue) String() (string, error)        { _ = "STUB: not implemented"; return "", nil }
func (v errValue) Scan(any) error                 { _ = "STUB: not implemented"; return nil }
func (v errValue) Load() any                      { _ = "STUB: not implemented"; return *new(any) }
func (v errValue) Store(any)                      { _ = "STUB: not implemented"; return }
func (v errValue) Slice() ([]Value, error)        { _ = "STUB: not implemented"; return nil, nil }
func (v errValue) Map() (map[string]Value, error) { _ = "STUB: not implemented"; return nil, nil }
