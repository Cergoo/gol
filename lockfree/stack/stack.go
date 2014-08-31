// (c) 2014 Cergoo
// under terms of ISC license

// Package stack it's a implementation lockfree LIFO stack
package stack

import (
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

// Push set item into stack
func (t *Tstack) Push(v interface{}) {
	node := &tnode{val: v}
	node.next = t.top
	for !atomic.CompareAndSwapPointer(&t.top, node.next, unsafe.Pointer(node)) {
		node.next = t.top
	}
}

// Pop get item from stack, if stack empty then ok == false
func (t *Tstack) Pop() (v interface{}, ok bool) {
	top := t.top
	if top == nil {
		return
	}
	for !atomic.CompareAndSwapPointer(&t.top, top, (*tnode)(top).next) {
		top = t.top
		if top == nil {
			return
		}
	}
	return (*tnode)(top).val, true
}
