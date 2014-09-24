// (c) 2014 Cergoo
// under terms of ISC license

// Package stacklf it's a implementation thread safe lockfree LIFO stack
package stacklf

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// Tstack it's a singly linked list
type Tstack struct {
	top unsafe.Pointer
}

type tnode struct {
	val  interface{}
	next unsafe.Pointer
}

// Push set item into stack if v not nil
func (t *Tstack) Push(v interface{}) {
	if v != nil {
		node := &tnode{val: v, next: t.top}
		for !atomic.CompareAndSwapPointer(&t.top, node.next, unsafe.Pointer(node)) {
			node.next = t.top
		}
	}
}

// Pop get item from stack, if stack empty then return nil
func (t *Tstack) Pop() interface{} {
	top := t.top
	if top == nil {
		return nil
	}
	for !atomic.CompareAndSwapPointer(&t.top, top, (*tnode)(top).next) {
		top = t.top
		if top == nil {
			return nil
		}
	}
	return (*tnode)(top).val
}

// PopWait get item from stack, if stack empty then wait
func (t *Tstack) PopWait() (v interface{}) {
	for {
		v = t.Pop()
		if v != nil {
			return
		}
		runtime.Gosched()
	}
}
