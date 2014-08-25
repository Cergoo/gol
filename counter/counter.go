/*
  easy atomic counter type
  (c) 2013-2014 Cergoo
  under terms of ISC license
*/

package counter

import (
	"sync/atomic"
)

type (
	T_counter struct {
		limit uint64
		value uint64
	}
)

// Get current count value
func (t *T_counter) Get() uint64 {
	return atomic.LoadUint64(&t.value)
}

// Set current count value
func (t *T_counter) Set(v uint64) {
	atomic.StoreUint64(&t.value, v)
}

// Increment
func (t *T_counter) Inc() uint64 {
	return atomic.AddUint64(&t.value, 1)
}

// Decrement
func (t *T_counter) Dec() uint64 {
	return atomic.AddUint64(&t.value, ^uint64(0)) // ^uint64(c-1)
}

// Add value
func (t *T_counter) Add(v int64) uint64 {
	return atomic.AddUint64(&t.value, uint64(v))
}

// Get current limit value
func (t *T_counter) GetLimit() uint64 {
	return atomic.LoadUint64(&t.limit)
}

// Set new limit value
func (t *T_counter) SetLimit(v uint64) {
	atomic.StoreUint64(&t.limit, v)
}

// Check limit value
func (t *T_counter) Check() bool {
	limit := t.GetLimit()
	return limit == 0 || limit > t.Get()
}

// Check limit value
func (t *T_counter) Check1(v uint64) bool {
	limit := t.GetLimit()
	return limit == 0 || limit > v
}
