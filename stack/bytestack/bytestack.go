// (c) 2014 Cergoo
// under terms of ISC license

// Package bytestack it's a simple fixed length slice stack implementation. No thread safe.
package bytestack

import (
	"fmt"
	"github.com/Cergoo/gol/err"
)

type (
	// TStack it's a main struct
	TStack struct {
		Stack      []byte
		LenElement int
	}
)

// New create new bytes stack
func New(lenElement int) *TStack {
	return &TStack{LenElement: lenElement}
}

// Push push item into stack
func (t *TStack) Push(val []byte) error {
	if len(val) != t.LenElement {
		return fmt.Errorf("Mismatch length stack element %d, and length value %d", t.LenElement, len(val))
	}
	t.Stack = append(t.Stack, val...)
	return nil
}

// PopPoint pop slice as pointer to value item from stack, nil if stack empty
func (t *TStack) PopPoint() (val []byte) {
	n := len(t.Stack) - t.LenElement
	if n >= 0 {
		val = t.Stack[n:]
		t.Stack = t.Stack[:n]
	}
	return
}

// PopVal pop slice as copy value item from stack, nil if stack empty
func (t *TStack) PopVal() (val []byte) {
	n := len(t.Stack) - t.LenElement
	if n >= 0 {
		val = make([]byte, t.LenElement)
		copy(val, t.Stack[n:])
		t.Stack = t.Stack[:n]
	}
	return
}

// Pop pop from stack into slice, not pop if stack empty
func (t *TStack) Pop(val []byte) []byte {
	n := len(t.Stack) - t.LenElement
	if n >= 0 {
		val = append(val, t.Stack[n:]...)
		t.Stack = t.Stack[:n]
	}
	return val
}

// Range range from last to first element of a stack and return point to element as a PopPoint
func (t *TStack) Range() chan []byte {
	defer err.Nopanic(nil)
	ch := make(chan []byte)
	n2 := len(t.Stack)
	n1 := n2 - t.LenElement
	go func() {
		for n1 >= 0 {
			ch <- t.Stack[n1:n2]
			n2 = n1
			n1 -= t.LenElement
		}
	}()
	return ch
}

// DelLast del elememt if it == val
func (t *TStack) DelLast(val []byte) {
	n2 := len(t.Stack)
	n1 := n2 - t.LenElement
	for n1 >= 0 {
		if string(val) == string(t.Stack[n1:n2]) {
			t.Stack = append(t.Stack[:n1], t.Stack[n2:]...)
			return
		}
		n2 = n1
		n1 -= t.LenElement
	}
}

// Clear clear stack
func (t *TStack) Clear() {
	t.Stack = t.Stack[:0]
}
