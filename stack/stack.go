// (c) 2014 Cergoo
// under terms of ISC license

// Package stack it's a implementation lockfree LIFO stack under counter & limiter items
package stack

import (
	"github.com/Cergoo/gol/counter"
	"github.com/Cergoo/gol/lockfree/stack"
)

type (
	// Tstack it's main structure
	Tstack struct {
		counter *counter.TCounter
		stack   *stack.Tstack
	}
)

// New it's constructor of new stack
func New() *Tstack {
	return &Tstack{
		counter: new(counter.TCounter),
		stack:   new(stack.Tstack),
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
func (t *Tstack) Pop() (v interface{}, ok bool) {
	v, ok = t.stack.Pop()
	if ok {
		t.counter.Dec()
	}
	return
}

// Counter get counter stack
func (t *Tstack) Counter() counter.ICounter {
	return t.counter
}