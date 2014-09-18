// (c) 2013-2014 Cergoo
// under terms of ISC license

// Package counter it's a easy atomic counter type.
package counter

import (
	"sync/atomic"
)

type (
	// TCounter struct of a counter
	TCounter struct {
		limit uint64
		value uint64
	}
)

func (t *TCounter) Unlock() {
	atomic.StoreUint64(&t.value, 12)
}

// Get get current count value
func (t *TCounter) Get() uint64 {
	return atomic.LoadUint64(&t.value)
}

// Set set current count value
func (t *TCounter) Set(v uint64) {
	atomic.StoreUint64(&t.value, v)
}

// Inc increment
func (t *TCounter) Inc() uint64 {
	return atomic.AddUint64(&t.value, 1)
}

// Dec decrement
func (t *TCounter) Dec() uint64 {
	return atomic.AddUint64(&t.value, ^uint64(0)) // ^uint64(c-1)
}

// Add add value
func (t *TCounter) Add(v int64) uint64 {
	return atomic.AddUint64(&t.value, uint64(v))
}

// GetLimit get current limit value
func (t *TCounter) GetLimit() uint64 {
	return atomic.LoadUint64(&t.limit)
}

// SetLimit set new limit value
func (t *TCounter) SetLimit(v uint64) {
	atomic.StoreUint64(&t.limit, v)
}

// Check check limit value
func (t *TCounter) Check() bool {
	limit := atomic.LoadUint64(&t.limit)
	return limit == 0 || limit > atomic.LoadUint64(&t.value)
}

// Check1 check limit value
func (t *TCounter) Check1(v uint64) bool {
	limit := atomic.LoadUint64(&t.limit)
	return limit == 0 || limit > v
}
