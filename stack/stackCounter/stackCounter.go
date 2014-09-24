// (c) 2014 Cergoo
// under terms of ISC license

// Package stackCounter it's a implementation lock free LIFO stack under counter & limiter items
package stackCounter

import (
	"github.com/Cergoo/gol/counter"
	"github.com/Cergoo/gol/stack/stacklf"
)

type (
	// Tstack it's main structure
	Tstack struct {
		counter counter.TCounter
		stack   stacklf.Tstack
	}
)

// New it's constructor of new stack
func New() *Tstack {
	return &Tstack{
		counter: counter.TCounter{},
		stack:   stacklf.Tstack{},
	}
}

// Push check counter, push item & increment counter
func (t *Tstack) Push(v interface{}) bool {
	if t.counter.Check() {
		t.counter.Inc()
		t.stack.Push(v)
		return true
	}
	return false
}

// Pop get item & decrement counter
func (t *Tstack) Pop() (v interface{}) {
	v = t.stack.Pop()
	if v != nil {
		t.counter.Dec()
	}
	return
}

// PopWait get item from stack, if stack empty then wait
func (t *Tstack) PopWait() (v interface{}) {
	v = t.stack.PopWait()
	t.counter.Dec()
	return
}

// Counter get counter stack
func (t *Tstack) Counter() counter.ICounter {
	return &t.counter
}
